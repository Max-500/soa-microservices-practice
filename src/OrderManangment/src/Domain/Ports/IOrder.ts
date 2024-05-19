import { IRabbitMQService } from "../../Application/Services/IRabbitMQService";
import { Order } from "../Entities/Order";

export interface IOrder {
    createOrders(data:any):Promise<Order[]|any>
    getOrders():Promise<Order[]|any>
    updateStatus(data:any, serviceRabbit:IRabbitMQService):Promise<any>
    getOrder(uuid:string, serviceRabbit:IRabbitMQService):Promise<any>
}