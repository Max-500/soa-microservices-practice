import { IOrder } from "../../Domain/Ports/IOrder";
import { RabbitMQService } from "../../Infraestructure/Services/RabbitMQService";

export class UpdateOrderUseCase {
    constructor(readonly repository:IOrder, readonly serviceRabbit:RabbitMQService){}

    async run(uuid:string) {
        
        return await this.repository.updateStatus(uuid, this.serviceRabbit);
    }
}