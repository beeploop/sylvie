import cors, { CorsOptions } from "cors";
import express from "express";
import { router } from "./routes/router";
import { NotFoundController } from "./controllers/notfound/controller";

export const app = express();

const corsOptions: CorsOptions = {
    origin: "*",
    allowedHeaders: ["GET", "POST", "PUT", "PATCH", "DELETE"],
};

app.use(express.urlencoded({ extended: true }));
app.use(express.json());
app.use(cors(corsOptions));

app.use("/api/v1", router);

app.use(NotFoundController());
