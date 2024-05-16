import { IOrder } from "../../Domain/Ports/IOrder";

export class CreateOrdersUseCase {
    constructor(readonly repository:IOrder){}

    async run(data:any) {
        return await this.repository.createOrders(data);
    }
    
}