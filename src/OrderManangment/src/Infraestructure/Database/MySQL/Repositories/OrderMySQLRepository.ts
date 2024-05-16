import { IOrder } from "../../../../Domain/Ports/IOrder";
import { generateUuid } from "../../../Helpers/Functions";
import { OrderModel } from "../Models/OrderMySQLModel";
import { OrderProductMySQModel } from "../Models/OrderProductMySQLModel";

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

    async updateStatus(data: any): Promise<any> {
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
        throw new Error("Method not implemented.");
    }

}