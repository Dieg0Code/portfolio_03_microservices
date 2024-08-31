import { Request, Response } from "express";
import { SaleController } from "./sale_controller";
import { SaleService } from "../services/sale_service";
import { CreateSaleRequest } from "../json/request/create_sale_request";
import { validateClass } from "../utils/validator";
import { Logger } from "winston";
import { SaleResponse } from "../json/response/sale_response";

export class SaleControllerImpl implements SaleController {

    private saleService: SaleService
    private logger: Logger

    constructor(saleService: SaleService, logger: Logger) {
        this.saleService = saleService;
        this.logger = logger;
    }
    async getAllSales(req: Request, res: Response): Promise<void> {
        try {
            const sales = await this.saleService.getAllSales();

            const response: BaseResponse<SaleResponse[]> = {
                code: 200,
                status: "OK",
                msg: "Sales retrieved successfully",
                data: sales
            }

            res.status(200).json(response);
        } catch (error) {
            this.logger.error(`Failed to get all sales: ${error}`);
            const response: BaseResponse<null> = {
                code: 500,
                status: "Internal Server Error",
                msg: "Failed to retrieve sales",
                data: null
            }

            res.status(500).json(response);
        }
    }

    // Implement the createSale method
    async createSale(req: Request, res: Response): Promise<void> {
        try {
            const createSaleRequest = new CreateSaleRequest(
                req.body.userID,
                req.body.products,
                req.body.totalAmount
            );

            await validateClass(createSaleRequest);

            const saleID = await this.saleService.createSale(createSaleRequest);

            const response: BaseResponse<string> = {
                code: 201,
                status: "Created",
                msg: "Sale created successfully",
                data: saleID
            }

            this.logger.info(`Sale created with ID: ${saleID}`);

            res.status(201).json(response);
        } catch (error) {
            this.logger.error(`Failed to create sale: ${error}`);
            const response: BaseResponse<null> = {
                code: 500,
                status: "Internal Server Error",
                msg: "Failed to create sale",
                data: null
            }

            res.status(500).json(response);
            
        }
    }

    // Implement the getSaleByID method
    async getSaleByID(req: Request, res: Response): Promise<void> {
        
        try {
            const saleID = req.params.saleID;

            const sale = await this.saleService.getSaleByID(saleID);

            if (!sale) {
                const response: BaseResponse<null> = {
                    code: 404,
                    status: "Not Found",
                    msg: "Sale not found",
                    data: null
                }

                res.status(404).json(response);
                return;
            }

            const response: BaseResponse<SaleResponse> = {
                code: 200,
                status: "OK",
                msg: "Sale retrieved successfully",
                data: sale
            }

            res.status(200).json(response);
        } catch (error) {
            this.logger.error(`Failed to get sale: ${error}`);
            const response: BaseResponse<null> = {
                code: 500,
                status: "Internal Server Error",
                msg: "Failed to retrieve sale",
                data: null
            }

            res.status(500).json(response);
        }
    }

    // Implement the getSalesByUserID method
    async getSalesByUserID(req: Request, res: Response): Promise<void> {
        
        try {
            const userID = parseInt(req.params.userID);

            const sales = await this.saleService.getSalesByUserID(userID);

            const response: BaseResponse<SaleResponse[]> = {
                code: 200,
                status: "OK",
                msg: "Sales retrieved successfully",
                data: sales
            }

            res.status(200).json(response);
        } catch (error) {
            this.logger.error(`Failed to get sales: ${error}`);
            const response: BaseResponse<null> = {
                code: 500,
                status: "Internal Server Error",
                msg: "Failed to retrieve sales",
                data: null
            }

            res.status(500).json(response);
        }

    }

    // Implement the getSalesByDate method
    async getSalesByDate(req: Request, res: Response): Promise<void> {
        
        try {
            const date = req.params.date;

            const sales = await this.saleService.getSalesByDate(date);

            const response: BaseResponse<SaleResponse[]> = {
                code: 200,
                status: "OK",
                msg: "Sales retrieved successfully",
                data: sales
            }

            res.status(200).json(response);
        } catch (error) {
            this.logger.error(`Failed to get sales: ${error}`);
            const response: BaseResponse<null> = {
                code: 500,
                status: "Internal Server Error",
                msg: "Failed to retrieve sales",
                data: null
            }

            res.status(500).json(response);
        }
    }

    async deleteSale(req: Request, res: Response): Promise<void> {
        
        try {
            const saleID = req.params.saleID;

            const result = await this.saleService.deleteSale(saleID);

            if (!result) {
                const response: BaseResponse<null> = {
                    code: 404,
                    status: "Not Found",
                    msg: "Sale not found",
                    data: null
                }

                res.status(404).json(response);
                return;
            }

            const response: BaseResponse<null> = {
                code: 200,
                status: "OK",
                msg: "Sale deleted successfully",
                data: null
            }

            res.status(200).json(response);
        } catch (error) {
            this.logger.error(`Failed to delete sale: ${error}`);
            const response: BaseResponse<null> = {
                code: 500,
                status: "Internal Server Error",
                msg: "Failed to delete sale",
                data: null
            }

            res.status(500).json(response);
        }
    }

}