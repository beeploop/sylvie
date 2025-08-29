import { Router } from "express";
import { indexRouter } from "./index/route";

export const router = Router();

router.use("/", indexRouter);
