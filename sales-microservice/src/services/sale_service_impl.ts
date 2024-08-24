import { Logger } from "winston";
import { CreateSaleRequest } from "../json/request/create_sale_request";
import { SaleResponse } from "../json/response/sale_response";
import { SalesRepository } from "../repository/sales_repository";
import { SaleService } from "./sale_service";
import { v4 as uuidv4 } from 'uuid';
import { Sale } from "../models/Sales";

export class SaleServiceImpl implements SaleService {

    private saleRepo: SalesRepository;
    private logger: Logger;

    constructor(saleRepo: SalesRepository, logger: Logger) {
        this.saleRepo = saleRepo;
        this.logger = logger;
    }
    async createSale(saleRequest: CreateSaleRequest): Promise<string> {
        try {
            const saleID = uuidv4();  // Generar un UUID para la venta
            const sale: Sale = {
                saleID,
                userID: saleRequest.userID,
                products: saleRequest.products,
                totalAmount: saleRequest.totalAmount,
                date: new Date().toISOString() // Fecha actual
            };

            await this.saleRepo.createSale(sale);
            this.logger.info(`Sale created with ID: ${saleID}`);
            return saleID;
        } catch (error) {
            this.logger.error(`Error creating sale: ${error}`);
            throw new Error("Error creating sale");
        }
    }
    async getSaleByID(saleID: string): Promise<SaleResponse | null> {
        try {
            const sale = await this.saleRepo.getSaleByID(saleID);
            if (!sale) {
                return null;
            }

            return {
                saleID: sale.saleID,
                userID: sale.userID,
                products: sale.products,
                totalAmount: sale.totalAmount,
                date: sale.date
            };

        } catch (error) {
            this.logger.error(`Failed to get sale with ID ${saleID}: ${error}`);
            throw new Error("Failed to retrieve sale");
        }
    }

    async getSalesByUserID(userID: number): Promise<SaleResponse[]> {
        try {
            const sales = await this.saleRepo.getSalesByUserID(userID);
            return sales.map(sale => {
                return {
                    saleID: sale.saleID,
                    userID: sale.userID,
                    products: sale.products,
                    totalAmount: sale.totalAmount,
                    date: sale.date
                };
            });
        } catch (error) {
            this.logger.error(`Failed to get sales for user with ID ${userID}: ${error}`);
            throw new Error("Failed to retrieve sales");
        }
    }
    async getSalesByDate(date: Date): Promise<SaleResponse[]> {
        const dateStr = date.toISOString();
        try {
            const sales = await this.saleRepo.getSalesByDate(dateStr);
            return sales.map(sale => {
                return {
                    saleID: sale.saleID,
                    userID: sale.userID,
                    products: sale.products,
                    totalAmount: sale.totalAmount,
                    date: sale.date
                };
            });
        } catch (error) {
            this.logger.error(`Failed to get sales for date ${date}: ${error}`);
            throw new Error("Failed to retrieve sales");
        }
    }
    async deleteSale(saleID: string): Promise<boolean> {
        try {
            const success = await this.saleRepo.deleteSale(saleID);
            if (success) {
                this.logger.info(`Sale with ID ${saleID} deleted`);
                return true;
            } else {
                this.logger.error(`Sale with ID ${saleID} not found`);
                return false;
            }
        } catch (error) {
            this.logger.error(`Failed to delete sale with ID ${saleID}: ${error}`);
            throw new Error("Failed to delete sale");
        }
    }
    

}