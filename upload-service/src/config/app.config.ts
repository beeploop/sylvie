import { config } from "dotenv";

config();

export type AppConfig = {
    PORT: number;
    StoragePath: string;
    RabbitURL: string;
    RabbitPublishQueueName: string;
};

export const appConfig: AppConfig = {
    PORT: process.env.PORT ? Number(process.env.PORT) : 3000,
    StoragePath: process.env.STORAGE_PATH || "uploads",
    RabbitURL: process.env.RABBIT_URL || "amqp://localhost",
    RabbitPublishQueueName: process.env.PUBLISH_QUEUE || "transcoding_queue",
};
