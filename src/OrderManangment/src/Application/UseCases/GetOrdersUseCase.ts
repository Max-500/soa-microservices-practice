import { IOrder } from "../../Domain/Ports/IOrder";

export class GetOrdersUseCase {
    constructor(readonly repository:IOrder){}

    async run() {
        return await this.repository.getOrders();
    }
}