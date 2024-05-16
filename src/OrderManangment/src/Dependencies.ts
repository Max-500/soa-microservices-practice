import { CreateOrdersUseCase } from "./Application/UseCases/CreateOrdersUseCase";
import { CreateOrdersController } from "./Infraestructure/Controllers/CreateOrdersController";
import { getOrderRepository } from "./Infraestructure/Database/GetRepositories";
import { MySQLConfig } from "./Infraestructure/Database/MySQL/Config/DatabaseConfig";
import { DatabaseConfig } from "./Infraestructure/Database/MySQL/Config/IDatabase";

export type DatabaseType = 'MySQL' | 'MongoDB';
const dbType: DatabaseType = 'MySQL';

function getDatabaseConfig(): DatabaseConfig {
    if (dbType === 'MySQL') {
      return new MySQLConfig();
    }
    throw new Error('Unsupported repository type');
}

const dbConfig = getDatabaseConfig();
dbConfig.initialize().then(() => {
  console.log('Database initialized.')
});

const orderRepository = getOrderRepository(dbType);

const createOrdersUseCase = new CreateOrdersUseCase(orderRepository);
export const createOrderController = new CreateOrdersController(createOrdersUseCase);