import sequelize from './database';
import { DatabaseConfig } from './IDatabase'
import "./Relations";

export class MySQLConfig implements DatabaseConfig {
  async initialize(): Promise<void> {
    await sequelize.sync({ force: false });
    console.log('MySQL database synchronized.');
  }
}