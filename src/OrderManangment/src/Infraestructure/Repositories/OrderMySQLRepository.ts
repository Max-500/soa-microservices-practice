import { IOrder } from "../../Domain/Ports/IOrder";
import { generateUuid } from "../Helpers/Functions";
import { OrderModel } from "../Database/MySQL/Models/OrderMySQLModel";
import { OrderProductMySQModel } from "../Database/MySQL/Models/OrderProductMySQLModel";
import { IRabbitMQService } from "../../Application/Services/IRabbitMQService";

export class OrderMySQLRepository implements IOrder {
    async createOrders(data: any): Promise<any> {
        const promisesOrders = [];
        const promisesOrderProduct = [];
        try {
            for (const order of data) {
                const orderUuid = generateUuid();
                const promiseOrder = OrderModel.create({
                    uuid: orderUuid,
                    total: order.total,
                    date: new Date(),
                    status: 'CREATED'
                });
                const promiseOrderProduct = OrderProductMySQModel.create({
                    uuid: generateUuid(),
                    price: order.total/order.amount,
                    amount: order.amount,
                    orderUuid: orderUuid,
                    productUuid: order.id
                })
                promisesOrders.push(promiseOrder)
                promisesOrderProduct.push(promiseOrderProduct)
            }

            const orders = await Promise.all(promisesOrders);

            const orderProducts = await Promise.all(promisesOrderProduct);
    
            return {
                status: 200,
                orders,
                orderProducts
            }    
            
        } catch (error) {
            return {
                status: 500,
                error
            }
        }
    }

    async getOrders(): Promise<any> {
        try {
            const data = await OrderModel.findAll();
            return {
                status: 200,
                data
            }
        } catch (error) {
            return {
                status: 500, 
                error
            }
        }
    }

    async updateStatus(data: any, serviceRabbit:IRabbitMQService): Promise<any> {
        try {
            const order = await OrderModel.findByPk(data)
            if(!order){
                return{
                    status: 404,
                    message: "La orden no existe"
                }
            }

            order.status = "SEND";
            await order.save();

            console.log(order.dataValues);
            const product = await OrderProductMySQModel.findOne({ where: { orderUuid: order.dataValues.uuid } })
            const message = { 
                uuid: product?.dataValues.productUuid,
                amount: product?.dataValues.amount
             }
            serviceRabbit.sendMessage("update_stock_queue", message)
            return {
                status: 200,
                order
            }
        } catch (error) {
            return {
                status: 500,
                error
            }
        }
    }

    async getOrder(uuid: string, serviceRabbit:IRabbitMQService): Promise<any> {
        try {
            const order = await OrderProductMySQModel.findOne({ where: { orderUuid: uuid } });
            console.log(order?.dataValues);
            if(!order){
                return {
                    status: 404,
                    message: "La orden no existe"
                }
            }
            const message = {
                "uuid": order.dataValues.productUuid
            }
            await serviceRabbit.sendMessage("send_get_products_queue", message);
            console.log("Ya envie a la cola");
            const result = await serviceRabbit.receiveMessage("receive_get_products_queue");
            console.log("Ya recibi la cola");
            return {
                status: 200,
                result
            }
        } catch (error) {
            return{
                status: 500,
                error
            }
        }
    }
}