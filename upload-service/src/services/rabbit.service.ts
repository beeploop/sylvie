import amqp from "amqplib";
import { AppConfig } from "../config/app.config";

export type Job = {
    video_id: string;
    path: string;
    resolutions: string[];
}

export type RabbitStatus = "closed" | "initialized";

export class RabbitMQ {
    url: string;
    queueName: string;
    connection: amqp.ChannelModel;
    channel: amqp.Channel;
    status: RabbitStatus;

    constructor(appConfig: AppConfig) {
        this.url = appConfig.RabbitURL;
        this.queueName = appConfig.RabbitPublishQueueName;
        this.status = "closed";
    }

    async initialize() {
        this.connection = await amqp.connect(this.url);
        this.channel = await this.connection.createChannel();

        await this.channel.assertQueue(this.queueName, { durable: true });
        this.status = "initialized";
    }

    async publish(job: Job): Promise<void> {
        this.channel.sendToQueue(this.queueName, Buffer.from(JSON.stringify(job)));
        console.log(`Published: ${job}`);
    }

    async close(): Promise<void> {
        await this.channel.close();
        await this.connection.close();
        this.status = "closed";
    }

    getStatus(): RabbitStatus {
        return this.status;
    }
}
