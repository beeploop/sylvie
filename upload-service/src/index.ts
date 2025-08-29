import { createServer } from "node:http";
import { app } from "./app";
import { appConfig } from "./config/app.config";

const server = createServer(app);

server.listen(appConfig.PORT, () =>
    console.log(`Listening in PORT: ${appConfig.PORT}`),
);
