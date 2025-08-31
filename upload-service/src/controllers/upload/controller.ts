import { Request, Response } from "express";
import { Job, RabbitMQ } from "../../services/rabbit.service";

export function UploadController(rabbit: RabbitMQ) {
    return async (req: Request, res: Response) => {
        const file = req.file;
        if (!file) {
            res.status(500).json({ message: "Uploaded file not found" });
            return;
        }

        if (rabbit.getStatus() === "closed") {
            await rabbit.initialize();
        }

        const job: Job = {
            video_id: file.filename.split(".")[0],
            path: file.path,
            resolutions: ["1080p", "720p", "480p", "360p", "240p", "144p"],
        };

        rabbit.publish(job).catch((err) => console.log(err));

        res.status(200).json({ message: "upload success" });
    };
}
