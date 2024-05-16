import { OrderModel } from "../Models/OrderMySQLModel";
import { OrderProductMySQModel } from "../Models/OrderProductMySQLModel";

OrderModel.hasMany(OrderProductMySQModel, { foreignKey: 'orderUuid' });
OrderProductMySQModel.belongsTo(OrderModel, { foreignKey: 'orderUuid' });