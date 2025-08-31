import { randomUUID } from "node:crypto";
import { Router } from "express";
import multer, { diskStorage } from "multer";
import { appConfig } from "../../config/app.config";
import { UploadController } from "../../controllers/upload/controller";
import { RabbitMQ } from "../../services/rabbit.service";

const upload = multer({
    dest: appConfig.StoragePath, storage: diskStorage({
        destination(_req, _file, callback) {
            callback(null, appConfig.StoragePath);
        },
        filename(_req, _file, callback) {
            const filename = `${randomUUID()}.mp4`
            callback(null, filename);
        },
    })
});

export const uploadRouter = Router();

const rabbit = new RabbitMQ(appConfig);

uploadRouter.post("/", upload.single("file"), UploadController(rabbit));
