import { v4 as uuidv4 } from "uuid";

export class Order {
    uuid:string;
    total:number;
    date:Date;
    status:string;

    constructor(total:number) {
        this.uuid = this.generateUuid();
        this.total = total;
        this.date = new Date();
        this.status = "";
    }

    generateUuid():string {
        return uuidv4();
    }
}