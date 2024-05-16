import { Router } from "express";
import { createOrderController, getOrdersController, updateOrderController } from "../../Dependencies";

export const router:Router = Router();

router.post("/", createOrderController.run.bind(createOrderController));
router.get("/", getOrdersController.run.bind(getOrdersController));
router.put("/:UUID", updateOrderController.run.bind(updateOrderController));