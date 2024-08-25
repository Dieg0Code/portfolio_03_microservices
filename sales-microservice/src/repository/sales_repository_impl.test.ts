import { DynamoDB } from 'aws-sdk';
import { Sale } from '../models/Sales';
import { SalesRepositoryImpl } from './sales_repository_impl';
import { Logger } from 'winston';

jest.mock('aws-sdk');
jest.mock('winston');

describe('SalesRepositoryImpl', () => {
    let salesRepository: SalesRepositoryImpl;
    let db: jest.Mocked<DynamoDB.DocumentClient>;
    let logger: jest.Mocked<Logger>;

    beforeEach(() => {
        db = new DynamoDB.DocumentClient() as jest.Mocked<DynamoDB.DocumentClient>;
        logger = {
            error: jest.fn(),
            info: jest.fn(),
            // Add other methods if necessary
        } as unknown as jest.Mocked<Logger>;
        salesRepository = new SalesRepositoryImpl(db, 'SalesTable', logger);
    });

    it('createSale_Success', async () => {
        const sale: Sale = {
            saleID: '123',
            userID: 1,
            products: [{ productID: 1, name: 'Product 1', price: 50, quantity: 2 }],
            totalAmount: 100,
            createdAt: "2021-09-01T00:00:00.000Z"
        };

        db.put.mockReturnValue({
            promise: jest.fn().mockResolvedValue({}),
        } as any);

        const result = await salesRepository.createSale(sale);
        expect(result).toEqual(sale);
        expect(logger.info).toHaveBeenCalledWith(`Sale with ID ${sale.saleID} created successfully`);
    });

    it('createSale_Failure', async () => {
        const sale: Sale = {
            saleID: '123',
            userID: 1,
            products: [{ productID: 1, name: 'Product 1', price: 50, quantity: 2 }],
            totalAmount: 100,
            createdAt: "2021-09-01T00:00:00.000Z"
        };

        const error = new Error('Something went wrong');
        db.put.mockReturnValue({
            promise: jest.fn().mockRejectedValue(error),
        } as any);

        await expect(salesRepository.createSale(sale)).rejects.toThrow('Something went wrong');
        expect(logger.error).toHaveBeenCalledWith(`Error creating sale with ID ${sale.saleID}: ${error}`);
    });

    it('getSaleByID_Success', async () => {
        const sale: Sale = {
            saleID: '123',
            userID: 1,
            products: [{ productID: 1, name: 'Product 1', price: 50, quantity: 2 }],
            totalAmount: 100,
            createdAt: "2021-09-01T00:00:00.000Z"
        };

        db.get.mockReturnValue({
            promise: jest.fn().mockResolvedValue({ Item: sale }),
        } as any);

        const result = await salesRepository.getSaleByID(sale.saleID);
        expect(result).toEqual(sale);
        expect(logger.error).not.toHaveBeenCalled();
    });

    it('getSaleByID_NotFound', async () => {
        const saleID = '123';

        db.get.mockReturnValue({
            promise: jest.fn().mockResolvedValue({ Item: undefined }),
        } as any);

        const result = await salesRepository.getSaleByID(saleID);
        expect(result).toBeNull();
        expect(logger.error).toHaveBeenCalledWith(`Sale with ID ${saleID} not found`);
    });

    it('getSaleByID_Failure', async () => {
        const saleID = '123';

        const error = new Error('Something went wrong');
        db.get.mockReturnValue({
            promise: jest.fn().mockRejectedValue(error),
        } as any);

        await expect(salesRepository.getSaleByID(saleID)).rejects.toThrow('Something went wrong');
        expect(logger.error).toHaveBeenCalledWith(`Error getting sale with ID ${saleID}: ${error}`);
    });

    it('getSalesByUserID_Success', async () => {
        const sales: Sale[] = [
            {
                saleID: '123',
                userID: 1,
                products: [{ productID: 1, name: 'Product 1', price: 50, quantity: 2 }],
                totalAmount: 100,
                createdAt: "2021-09-01T00:00:00.000Z"
            },
            {
                saleID: '124',
                userID: 1,
                products: [{ productID: 2, name: 'Product 2', price: 100, quantity: 1 }],
                totalAmount: 100,
                createdAt: "2021-09-01T00:00:00.000Z"
            }
        ];

        db.scan.mockReturnValue({
            promise: jest.fn().mockResolvedValue({ Items: sales }),
        } as any);

        const result = await salesRepository.getSalesByUserID(1);
        expect(result).toEqual(sales);
        expect(logger.error).not.toHaveBeenCalled();
    });

    it('getSalesByUserID_Failure', async () => {
        const userID = 1;

        const error = new Error('Something went wrong');
        db.scan.mockReturnValue({
            promise: jest.fn().mockRejectedValue(error),
        } as any);

        await expect(salesRepository.getSalesByUserID(userID)).rejects.toThrow('Something went wrong');
        expect(logger.error).toHaveBeenCalledWith(`Error getting sales for user with ID ${userID}: ${error}`);
    });

    it('getSalesByDate_Success', async () => {
        const sales: Sale[] = [
            {
                saleID: '123',
                userID: 1,
                products: [{ productID: 1, name: 'Product 1', price: 50, quantity: 2 }],
                totalAmount: 100,
                createdAt: "2021-09-01T00:00:00.000Z"
            },
            {
                saleID: '124',
                userID: 1,
                products: [{ productID: 2, name: 'Product 2', price: 100, quantity: 1 }],
                totalAmount: 100,
                createdAt: "2021-09-01T00:00:00.000Z"
            }
        ];

        db.scan.mockReturnValue({
            promise: jest.fn().mockResolvedValue({ Items: sales }),
        } as any);

        const result = await salesRepository.getSalesByDate(new Date('2021-09-01').toISOString());
        expect(result).toEqual(sales);
        expect(logger.error).not.toHaveBeenCalled();
    });

    it('getSalesByDate_Failure', async () => {
        const date = new Date('2021-09-01').toISOString();

        const error = new Error('Something went wrong');
        db.scan.mockReturnValue({
            promise: jest.fn().mockRejectedValue(error),
        } as any);

        await expect(salesRepository.getSalesByDate(date)).rejects.toThrow('Something went wrong');
        expect(logger.error).toHaveBeenCalledWith(`Error getting sales for date ${date}: ${error}`);
    });

    it('deleteSale_Success', async () => {
        const saleID = '123';

        db.delete.mockReturnValue({
            promise: jest.fn().mockResolvedValue({}),
        } as any);

        const result = await salesRepository.deleteSale(saleID);
        expect(result).toBe(true);
        expect(logger.info).toHaveBeenCalledWith(`Sale with ID ${saleID} deleted successfully`);
    });

    it('deleteSale_Failure', async () => {
        const saleID = '123';

        const error = new Error('Something went wrong');
        db.delete.mockReturnValue({
            promise: jest.fn().mockRejectedValue(error),
        } as any);

        await expect(salesRepository.deleteSale(saleID)).rejects.toThrow('Something went wrong');
        expect(logger.error).toHaveBeenCalledWith(`Error deleting sale with ID ${saleID}: ${error}`); 
    });
});
