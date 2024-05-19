import { IOrder } from "../../Domain/Ports/IOrder";
import { IRabbitMQService } from "../Services/IRabbitMQService";

export class GetOrderUseCase {
    constructor(readonly repository:IOrder, readonly service:IRabbitMQService){}

    async run (uuid:string) {
        return await this.repository.getOrder(uuid, this.service);
    }

}