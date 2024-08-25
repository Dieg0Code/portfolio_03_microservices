import { Request, Response } from 'express';

export interface SaleController {
    createSale(req: Request, res: Response): Promise<void>;
    getSaleByID(req: Request, res: Response): Promise<void>;
    getSalesByUserID(req: Request, res: Response): Promise<void>;
    getSalesByDate(req: Request, res: Response): Promise<void>;
    getAllSales(req: Request, res: Response): Promise<void>;
    deleteSale(req: Request, res: Response): Promise<void>;
}
