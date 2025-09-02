import express from "express";
import { appConfig } from "./config/app.config";
import cors, { CorsOptions } from "cors";

export const app = express();

const corsOptions: CorsOptions = {
    origin: "*",
    methods: ["GET", "POST"],
};

app.use(cors(corsOptions));
app.use(
    "/hls",
    express.static(appConfig.videos_directory, {
        setHeaders: (res, path) => {
            if (path.endsWith(".m3u8")) {
                res.set("Content-Type", "application/vnd.apple.mpegurl");
            } else if (path.endsWith(".ts")) {
                res.set("Content-Type", "video/mp2t");
            }
        },
    }),
);
