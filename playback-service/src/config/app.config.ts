import dotenv from "dotenv";

dotenv.config();

export type AppConfig = {
    port: number;
    videos_directory: string;
};

export const appConfig: AppConfig = {
    port: process.env.PORT ? Number(process.env.PORT) : 3002,
    videos_directory: process.env.VIDEOS_DIRECTORY || "videos",
};
