import { DataTypes, Model } from "sequelize";
import sequelize from "../Config/database";

export class OrderModel extends Model {
    uuid!:string;
    total!:number;
    date!:Date;
    status!:string;
}

OrderModel.init({
    uuid: { type: DataTypes.UUID, defaultValue:DataTypes.UUIDV4, primaryKey:true, allowNull:false },
    total: { type:DataTypes.INTEGER, defaultValue:null },
    date: { type:DataTypes.DATE },
    status: { type:DataTypes.STRING }
},{sequelize, modelName:'orders'});