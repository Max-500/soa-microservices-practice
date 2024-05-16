import { Request, Response } from "express";
import { CreateOrdersUseCase } from "../../Application/UseCases/CreateOrdersUseCase";

export class CreateOrdersController {
    constructor(readonly createOrdersUseCase:CreateOrdersUseCase){}

    async run(req:Request, res:Response) {
        const response = await this.createOrdersUseCase.run(req.body.data);
        return res.status(response.status).json(response);
    }
}