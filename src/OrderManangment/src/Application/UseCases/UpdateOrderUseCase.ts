import { IOrder } from "../../Domain/Ports/IOrder";
import { IRabbitMQService } from "../Services/IRabbitMQService";

export class UpdateOrderUseCase {
    constructor(readonly repository:IOrder, readonly serviceRabbit:IRabbitMQService){}

    async run(uuid:string) {
        
        return await this.repository.updateStatus(uuid, this.serviceRabbit);
    }
}