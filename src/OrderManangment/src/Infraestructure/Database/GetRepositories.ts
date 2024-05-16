import { OrderMySQLRepository } from "./MySQL/Repositories/OrderMySQLRepository";

export function getOrderRepository(dbType:string) {
    if(dbType == "MySQL") return new OrderMySQLRepository();
    return new OrderMySQLRepository();
}