import { Router } from "express";
import { createOrderController } from "../../Dependencies";

export const router:Router = Router();

router.post("/", createOrderController.run.bind(createOrderController));