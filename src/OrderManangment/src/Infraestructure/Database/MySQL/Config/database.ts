import { Sequelize } from 'sequelize';

const sequelize = new Sequelize("orders", "root", "", {
    host: "localhost",
    dialect: "mysql",
});
  
export default sequelize;