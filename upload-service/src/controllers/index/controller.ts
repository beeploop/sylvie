import { Request, Response } from "express";

export function IndexController() {
    return async (_req: Request, res: Response) => {
        const time = new Date().toUTCString();
        res.status(200).json({ time: time });
    };
}
