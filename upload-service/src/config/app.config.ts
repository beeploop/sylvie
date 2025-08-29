import { config } from "dotenv";

config();

type AppConfig = {
    PORT: number;
};

export const appConfig: AppConfig = {
    PORT: process.env.PORT ? Number(process.env.PORT) : 3000,
};
