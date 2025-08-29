import { Request, Response, Router } from "express";
import { IndexController } from "../../controllers/index/controller";

export const indexRouter = Router();

indexRouter.get("/", IndexController());
