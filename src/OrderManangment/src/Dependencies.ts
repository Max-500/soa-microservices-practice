import { CreateOrdersUseCase } from "./Application/UseCases/CreateOrdersUseCase";
import { GetOrderUseCase } from "./Application/UseCases/GetOrderUseCase";
import { GetOrdersUseCase } from "./Application/UseCases/GetOrdersUseCase";
import { UpdateOrderUseCase } from "./Application/UseCases/UpdateOrderUseCase";
import { CreateOrdersController } from "./Infraestructure/Controllers/CreateOrdersController";
import { GetOrderController } from "./Infraestructure/Controllers/GetOrderController";
import { GetOrdersController } from "./Infraestructure/Controllers/GetOrdersController";
import { UpdateOrderController } from "./Infraestructure/Controllers/UpdateOrderController";
import { getOrderRepository } from "./Infraestructure/Database/GetRepositories";
import { MySQLConfig } from "./Infraestructure/Database/MySQL/Config/DatabaseConfig";
import { DatabaseConfig } from "./Infraestructure/Database/MySQL/Config/IDatabase";
import { RabbitMQService } from "./Infraestructure/Services/RabbitMQService";

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


// Conexion a RabbitMQ
const rabbitMQServicePromise = RabbitMQService.getInstance();

const orderRepository = getOrderRepository(dbType);

const createOrdersUseCase = new CreateOrdersUseCase(orderRepository);
export const createOrderController = new CreateOrdersController(createOrdersUseCase);

const getOrdersUseCase = new GetOrdersUseCase(orderRepository);
export const getOrdersController = new GetOrdersController(getOrdersUseCase);

export const updateOrderController = rabbitMQServicePromise.then(rabbitMQService => {
  const updateOrderUseCase = new UpdateOrderUseCase(orderRepository, rabbitMQService);
  const updateOrderController = new UpdateOrderController(updateOrderUseCase);
  return updateOrderController
});

export const getOrderController = rabbitMQServicePromise.then(rabbitMQService => {
  const getOrderUseCase = new GetOrderUseCase(orderRepository, rabbitMQService)
  return new GetOrderController(getOrderUseCase);
});