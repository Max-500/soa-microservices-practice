import { RabbitMQService } from "../../Infraestructure/Services/RabbitMQService";
import { Order } from "../Entities/Order";

export interface IOrder {
    createOrders(data:any):Promise<Order[]|any>
    getOrders():Promise<Order[]|any>
    updateStatus(data:any, serviceRabbit:RabbitMQService):Promise<any>
}