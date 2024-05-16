import { IOrder } from "../../Domain/Ports/IOrder";

export class UpdateOrderUseCase {
    constructor(readonly repository:IOrder){}

    async run(uuid:string) {
        return await this.repository.updateStatus(uuid);
    }
}