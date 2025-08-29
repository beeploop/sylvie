import { Request, Response } from "express";

export function NotFoundController() {
    return (_req: Request, res: Response) => {
        res.status(404).send("Route not found");
    }
}
