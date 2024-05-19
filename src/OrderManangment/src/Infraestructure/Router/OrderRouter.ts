import { Router } from "express";
import { createOrderController, getOrdersController, updateOrderController as updateOrderControllerPromise, getOrderController as getOrderControllerPromise } from "../../Dependencies";

export const router:Router = Router();

router.post("/", createOrderController.run.bind(createOrderController));
router.get("/", getOrdersController.run.bind(getOrdersController));

async function initializeUpdateOrderController() {
  const updateOrderController = await updateOrderControllerPromise;
  router.put("/:UUID", updateOrderController.run.bind(updateOrderController));

  const getOrderController = await getOrderControllerPromise;
  router.get("/:orderUuid", getOrderController.run.bind(getOrderController));
}

initializeUpdateOrderController();
