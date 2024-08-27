import { CreateSaleRequest } from "../json/request/create_sale_request";
import { SaleResponse } from "../json/response/sale_response";

export interface SaleService {
    createSale(saleRequest: CreateSaleRequest): Promise<string>;
    getSaleByID(saleID: string): Promise<SaleResponse | null>;
    getSalesByUserID(userID: number): Promise<SaleResponse[]>;
    getSalesByDate(date: string): Promise<SaleResponse[]>;
    getAllSales(): Promise<SaleResponse[]>;
    deleteSale(saleID: string): Promise<boolean>;
}