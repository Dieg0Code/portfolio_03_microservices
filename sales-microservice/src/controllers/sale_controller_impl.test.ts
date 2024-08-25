import 'reflect-metadata';
import { SaleControllerImpl } from './sale_controller_impl';
import { SaleService } from '../services/sale_service';
import { Logger } from 'winston';
import { Request, Response } from 'express';
import { SaleResponse } from '../json/response/sale_response';
import { validateClass } from '../utils/validator';
import { CreateSaleRequest } from '../json/request/create_sale_request';

// Mock the external dependencies
jest.mock('../utils/validator');
jest.mock('../json/request/create_sale_request');

describe('SaleControllerImpl', () => {
  let saleController: SaleControllerImpl;
  let mockSaleService: jest.Mocked<SaleService>;
  let mockLogger: jest.Mocked<Logger>;
  let mockRequest: Partial<Request>;
  let mockResponse: Partial<Response>;

  beforeEach(() => {
    mockSaleService = {
      createSale: jest.fn(),
      getSaleByID: jest.fn(),
      getSalesByUserID: jest.fn(),
      getSalesByDate: jest.fn(),
      deleteSale: jest.fn(),
    } as any;

    mockLogger = {
      info: jest.fn(),
      error: jest.fn(),
    } as any;

    saleController = new SaleControllerImpl(mockSaleService, mockLogger);

    mockResponse = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn(),
    };
  });

  describe('createSale', () => {
    it('should create a sale successfully', async () => {
      mockRequest = {
        body: {
          userID: 1,
          products: [{ productID: 1, name: 'Prod', price: 1000, quantity: 2 }],
          totalAmount: 2000,
        },
      };

      (CreateSaleRequest as jest.MockedClass<typeof CreateSaleRequest>).mockImplementation(() => ({
        userID: 1,
        products: [{ productID: 1, name: 'Prod', price: 1000, quantity: 2 }],
        totalAmount: 2000,
      }));

      (validateClass as jest.Mock).mockResolvedValue(undefined);
      mockSaleService.createSale.mockResolvedValue('mocked-sale-id');

      await saleController.createSale(mockRequest as Request, mockResponse as Response);

      expect(mockResponse.status).toHaveBeenCalledWith(201);
      expect(mockResponse.json).toHaveBeenCalledWith(expect.objectContaining({
        code: 201,
        status: 'Created',
        msg: 'Sale created successfully',
        data: 'mocked-sale-id',
      }));
    });

    it('should handle errors when creating a sale', async () => {
      mockRequest = {
        body: {
          userID: 1,
          products: [{ productID: 1, quantity: 2 }],
          totalAmount: 100,
        },
      };

      (validateClass as jest.Mock).mockRejectedValue(new Error('Validation error'));

      await saleController.createSale(mockRequest as Request, mockResponse as Response);

      expect(mockResponse.status).toHaveBeenCalledWith(500);
      expect(mockResponse.json).toHaveBeenCalledWith(expect.objectContaining({
        code: 500,
        status: 'Internal Server Error',
        msg: 'Failed to create sale',
        data: null,
      }));
    });
  });

  describe('getSaleByID', () => {
    it('should get a sale by ID successfully', async () => {
      const mockSale: SaleResponse = {
        saleID: 'mocked-sale-id',
        userID: 1,
        products: [{ productID: 1, name: 'Prod', price: 1000, quantity: 2 }],
        totalAmount: 100,
        createdAt: '2024-08-24T00:00:00.000Z',
      };

      mockRequest = {
        params: { saleID: 'mocked-sale-id' },
      };

      mockSaleService.getSaleByID.mockResolvedValue(mockSale);

      await saleController.getSaleByID(mockRequest as Request, mockResponse as Response);

      expect(mockResponse.status).toHaveBeenCalledWith(200);
      expect(mockResponse.json).toHaveBeenCalledWith(expect.objectContaining({
        code: 200,
        status: 'OK',
        msg: 'Sale retrieved successfully',
        data: mockSale,
      }));
    });

    it('should handle not found when getting a sale by ID', async () => {
      mockRequest = {
        params: { saleID: 'non-existent-id' },
      };

      mockSaleService.getSaleByID.mockResolvedValue(null);

      await saleController.getSaleByID(mockRequest as Request, mockResponse as Response);

      expect(mockResponse.status).toHaveBeenCalledWith(404);
      expect(mockResponse.json).toHaveBeenCalledWith(expect.objectContaining({
        code: 404,
        status: 'Not Found',
        msg: 'Sale not found',
        data: null,
      }));
    });
  });

  describe('getSalesByUserID', () => {
    it('should get sales by user ID successfully', async () => {
      const mockSales: SaleResponse[] = [
        {
          saleID: 'mocked-sale-id-1',
          userID: 1,
          products: [{ productID: 1, name: 'Prod', price: 1000, quantity: 2 }],
          totalAmount: 2000,
          createdAt: '2024-08-24T00:00:00.000Z',
        },
        {
          saleID: 'mocked-sale-id-2',
          userID: 1,
          products: [{ productID: 2, name: 'Prod2', price: 1000, quantity: 1 }],
          totalAmount: 1000,
          createdAt: '2024-08-25T00:00:00.000Z',
        },
      ];

      mockRequest = {
        params: { userID: '1' },
      };

      mockSaleService.getSalesByUserID.mockResolvedValue(mockSales);

      await saleController.getSalesByUserID(mockRequest as Request, mockResponse as Response);

      expect(mockResponse.status).toHaveBeenCalledWith(200);
      expect(mockResponse.json).toHaveBeenCalledWith(expect.objectContaining({
        code: 200,
        status: 'OK',
        msg: 'Sales retrieved successfully',
        data: mockSales,
      }));
    });
  });

  describe('getSalesByDate', () => {
    it('should get sales by date successfully', async () => {
      const mockSales: SaleResponse[] = [
        {
          saleID: 'mocked-sale-id-1',
          userID: 1,
          products: [{ productID: 1, name: 'Prod', price: 1000, quantity: 2 }],
          totalAmount: 100,
          createdAt: '2024-08-24T00:00:00.000Z',
        },
      ];

      mockRequest = {
        params: { date: '2024-08-24' },
      };

      mockSaleService.getSalesByDate.mockResolvedValue(mockSales);

      await saleController.getSalesByDate(mockRequest as Request, mockResponse as Response);

      expect(mockResponse.status).toHaveBeenCalledWith(200);
      expect(mockResponse.json).toHaveBeenCalledWith(expect.objectContaining({
        code: 200,
        status: 'OK',
        msg: 'Sales retrieved successfully',
        data: mockSales,
      }));
    });
  });

  describe('deleteSale', () => {
    it('should delete a sale successfully', async () => {
      mockRequest = {
        params: { saleID: 'mocked-sale-id' },
      };

      mockSaleService.deleteSale.mockResolvedValue(true);

      await saleController.deleteSale(mockRequest as Request, mockResponse as Response);

      expect(mockResponse.status).toHaveBeenCalledWith(200);
      expect(mockResponse.json).toHaveBeenCalledWith(expect.objectContaining({
        code: 200,
        status: 'OK',
        msg: 'Sale deleted successfully',
        data: null,
      }));
    });

    it('should handle not found when deleting a sale', async () => {
      mockRequest = {
        params: { saleID: 'non-existent-id' },
      };

      mockSaleService.deleteSale.mockResolvedValue(false);

      await saleController.deleteSale(mockRequest as Request, mockResponse as Response);

      expect(mockResponse.status).toHaveBeenCalledWith(404);
      expect(mockResponse.json).toHaveBeenCalledWith(expect.objectContaining({
        code: 404,
        status: 'Not Found',
        msg: 'Sale not found',
        data: null,
      }));
    });
  });
});