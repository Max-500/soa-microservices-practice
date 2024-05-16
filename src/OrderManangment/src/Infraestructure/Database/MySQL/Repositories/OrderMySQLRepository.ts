import { IOrder } from "../../../../Domain/Ports/IOrder";

export class OrderMySQLRepository implements IOrder {
    async createOrders(data: any): Promise<any> {
        throw new Error("Method not implemented.");
    }
    async getOrders(): Promise<any> {
        throw new Error("Method not implemented.");
    }
    async updateStatus(data: any): Promise<any> {
        throw new Error("Method not implemented.");
    }

}