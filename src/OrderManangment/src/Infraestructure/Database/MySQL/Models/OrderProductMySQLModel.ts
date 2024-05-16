import { DataTypes, Model } from "sequelize";
import sequelize from "../Config/database";

export class OrderProductMySQModel extends Model {
    uuid!:string;
    price!:number;
    amount!:number;
    orderUuid!:string;
    productUuid!:string;
}

OrderProductMySQModel.init({
    uuid: { type: DataTypes.UUID, defaultValue:DataTypes.UUIDV4, primaryKey:true, allowNull:false },
    price: { type: DataTypes.INTEGER },
    amount: { type: DataTypes.INTEGER }
}, { sequelize, modelName:'orders_products' })