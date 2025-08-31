import { Router } from "express";
import { indexRouter } from "./index/route";
import { uploadRouter } from "./upload/router";

export const router = Router();

router.use("/", indexRouter);
router.use("/uploads", uploadRouter);
