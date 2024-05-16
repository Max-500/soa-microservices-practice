import { Request, Response } from "express";
import { GetOrdersUseCase } from "../../Application/UseCases/GetOrdersUseCase";

export class GetOrdersController {
    constructor(readonly useCase:GetOrdersUseCase){}

    async run(req:Request, res:Response){
        const response = await this.useCase.run();
        return res.status(response.status).json(response)
    }
}