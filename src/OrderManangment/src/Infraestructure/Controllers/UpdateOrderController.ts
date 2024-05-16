import { Request, Response } from "express";
import { UpdateOrderUseCase } from "../../Application/UseCases/UpdateOrderUseCase";

export class UpdateOrderController {
    constructor(readonly useCase:UpdateOrderUseCase){}

    async run(req:Request, res:Response){
        const response = await this.useCase.run(req.params.UUID);
        return res.status(response.status).json(response);
    }
}