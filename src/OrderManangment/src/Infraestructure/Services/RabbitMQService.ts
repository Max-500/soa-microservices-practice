import { IRabbitMQService } from "../../Application/Services/IRabbitMQService";
import * as amqplib from 'amqplib';

export class RabbitMQService implements IRabbitMQService {
    private static instance: RabbitMQService;
    readonly connection: amqplib.Connection;

    private constructor(connection: amqplib.Connection) {
        this.connection = connection;
    }

    static async getInstance(): Promise<RabbitMQService> {
        if (!RabbitMQService.instance) {
            const connection = await amqplib.connect('amqp://guest:guest@localhost:5672');
            RabbitMQService.instance = new RabbitMQService(connection);
        }
        return RabbitMQService.instance;
    }

    async sendMessage(queue: string, message: any): Promise<void> {
        const channel = await this.connection.createChannel();
        await channel.assertQueue(queue, { durable: true });
        channel.sendToQueue(queue, Buffer.from(JSON.stringify(message)));
        console.log(`Mensaje enviado a la cola ${queue}: ${JSON.stringify(message)}`);
    }

    async receiveMessage(queue: string): Promise<any> {
        throw new Error("Method not implemented.");
    }
}
