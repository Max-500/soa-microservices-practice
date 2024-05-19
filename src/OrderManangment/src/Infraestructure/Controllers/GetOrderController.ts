import { Request, Response } from "express";
import { GetOrderUseCase } from "../../Application/UseCases/GetOrderUseCase";

export class GetOrderController {
    constructor(readonly useCase:GetOrderUseCase){}

    async run(req:Request, res:Response){
        const response = await this.useCase.run(req.params.orderUuid);
        return res.status(response.status).json(response)
    }
}