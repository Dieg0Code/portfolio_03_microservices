import { DynamoDB } from "aws-sdk";
import { Sale } from "../models/Sales";
import { SalesRepository } from "./sales_repository";
import { Logger } from "winston";


export class SalesRepositoryImpl implements SalesRepository {

    private db: DynamoDB.DocumentClient;
    private tableName: string;
    private logger: Logger;

    constructor(db: DynamoDB.DocumentClient, tableName: string, logger: Logger) {
        this.db = db;
        this.tableName = tableName;
        this.logger = logger;
    }


    async createSale(sale: Sale): Promise<Sale> {
        try {
            const params = {
                TableName: this.tableName,
                Item: {
                    saleID: sale.saleID,
                    userID: sale.userID,
                    products: sale.products,
                    totalAmount: sale.totalAmount,
                    date: sale.date
                }
            };

            await this.db.put(params).promise();
            this.logger.info(`Sale with ID ${sale.saleID} created successfully`);
            return sale;
        } catch (error) {
            this.logger.error(`Error creating sale with ID ${sale.saleID}: ${error}`);
            throw error;
        }
    }

    async getSaleByID(saleID: string): Promise<Sale | null> {

        const params = {
            TableName: this.tableName,
            Key: {
                saleID: saleID
            }
        };

        try {
            const result = await this.db.get(params).promise();
            if (!result.Item) {
                this.logger.error(`Sale with ID ${saleID} not found`);
                return null;
            }
            return result.Item as Sale;
        } catch (error) {
            this.logger.error(`Error getting sale with ID ${saleID}: ${error}`);
            throw error;
        }
    }

    async getSalesByUserID(userID: number): Promise<Sale[]> {
        const params = {
            TableName: this.tableName,
            FilterExpression: 'userID = :userID',
            ExpressionAttributeValues: {
                ':userID': userID
            }
        };

        try {
            const result = await this.db.scan(params).promise();
            return result.Items as Sale[];
        } catch (error) {
            this.logger.error(`Error getting sales for user with ID ${userID}: ${error}`);
            throw error;
        }
    }

    async getSalesByDate(date: Date): Promise<Sale[]> {
        const params = {
            TableName: this.tableName,
            FilterExpression: 'date = :date',
            ExpressionAttributeValues: {
                ':date': date
            }
        };

        try {
            const result = await this.db.scan(params).promise();
            return result.Items as Sale[];
        } catch (error) {
            this.logger.error(`Error getting sales for date ${date}: ${error}`);
            throw error;
        }
    }

    async deleteSale(saleID: string): Promise<boolean> {
        const params = {
            TableName: this.tableName,
            Key: {
                saleID: saleID
            }
        };

        try {
            await this.db.delete(params).promise();
            this.logger.info(`Sale with ID ${saleID} deleted successfully`);
            return true;
        } catch (error) {
            this.logger.error(`Error deleting sale with ID ${saleID}: ${error}`);
            throw error;
        }
    }

}