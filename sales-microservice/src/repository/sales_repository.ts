import { Sale } from "../models/Sales";

export interface SalesRepository {
    createSale(sale: Sale): Promise<Sale>;
    getSaleByID(saleID: string): Promise<Sale | null> ;
    getSalesByUserID(userID: number): Promise<Sale[]>;
    getSalesByDate(date: string): Promise<Sale[]>;
    deleteSale(saleID: string): Promise<boolean>;
}