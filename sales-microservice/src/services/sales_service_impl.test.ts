import { SaleServiceImpl } from './sale_service_impl';
import { SalesRepository } from '../repository/sales_repository';
import { Logger } from 'winston';
import { CreateSaleRequest } from '../json/request/create_sale_request';
import { Sale } from '../models/Sales';

jest.mock('uuid', () => ({ v4: () => 'mocked-uuid' }));

describe('SaleServiceImpl', () => {
    let saleService: SaleServiceImpl;
    let mockSalesRepository: jest.Mocked<SalesRepository>;
    let mockLogger: jest.Mocked<Logger>;
    let originalDate: DateConstructor;
  
    beforeEach(() => {
      mockSalesRepository = {
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
  
      saleService = new SaleServiceImpl(mockSalesRepository, mockLogger);
  
      // Guardar la implementación original de Date
      originalDate = global.Date;
    });
  
    afterEach(() => {
      // Restaurar la implementación original de Date después de cada prueba
      global.Date = originalDate;
    });
  
    describe('createSale', () => {
      it('should create a sale successfully', async () => {
        const saleRequest: CreateSaleRequest = {
          userID: 1,
          products: [{ productID: 1, name: 'Prod 1', price: 50, quantity: 2 }],
          totalAmount: 100,
        };
  
        const mockDate = new Date('2024-08-24T00:15:17.650Z');
        // Usar una función de clase mock para Date
        global.Date = class extends Date {
          constructor() {
            super();
            return mockDate;
          }
        } as DateConstructor;
  
        await saleService.createSale(saleRequest);
  
        expect(mockSalesRepository.createSale).toHaveBeenCalledWith(expect.objectContaining({
          saleID: 'mocked-uuid',
          userID: 1,
          products: [{ productID: 1, name: 'Prod 1', price: 50, quantity: 2 }],
          totalAmount: 100,
          createdAt: mockDate.toISOString()
        }));
        expect(mockLogger.info).toHaveBeenCalledWith('Sale created with ID: mocked-uuid');
      });
  
      it('should throw an error if creation fails', async () => {
        mockSalesRepository.createSale.mockRejectedValue(new Error('DB error'));
  
        await expect(saleService.createSale({} as CreateSaleRequest)).rejects.toThrow('Error creating sale');
        expect(mockLogger.error).toHaveBeenCalled();
      });
    });
  
    describe('getSaleByID', () => {
      it('should return a sale when found', async () => {
        const mockSale: Sale = {
          saleID: 'test-id',
          userID: 1,
          products: [{ productID: 1, name: 'Prod 1', price: 50, quantity: 2 }],
          totalAmount: 100,
          createdAt: '2024-08-24T00:15:17.650Z',
        };
        mockSalesRepository.getSaleByID.mockResolvedValue(mockSale);
  
        const result = await saleService.getSaleByID('test-id');
  
        expect(result).toEqual(mockSale);
      });
  
      it('should return null when sale is not found', async () => {
        mockSalesRepository.getSaleByID.mockResolvedValue(null);
  
        const result = await saleService.getSaleByID('non-existent-id');
  
        expect(result).toBeNull();
      });
    });
  
    describe('getSalesByUserID', () => {
      it('should return sales for a given user ID', async () => {
        const mockSales: Sale[] = [
          {
            saleID: 'sale-1',
            userID: 1,
            products: [{ productID: 1, name: 'Prod 1', price: 50, quantity: 2 }],
            totalAmount: 100,
            createdAt: '2024-08-24T00:15:17.650Z',
          },
          {
            saleID: 'sale-2',
            userID: 1,
            products: [{ productID: 2, name: 'Prod 2', price: 75, quantity: 1 }],
            totalAmount: 75,
            createdAt: '2024-08-25T00:15:17.650Z',
          },
        ];
        mockSalesRepository.getSalesByUserID.mockResolvedValue(mockSales);
  
        const result = await saleService.getSalesByUserID(1);
  
        expect(result).toEqual(mockSales);
      });
    });
  
    describe('getSalesByDate', () => {
      it('should return sales for a given date', async () => {
        const mockDate = new Date('2024-08-24');
        const mockSales: Sale[] = [
          {
            saleID: 'sale-1',
            userID: 1,
            products: [{ productID: 1, name: 'Prod 1', price: 50, quantity: 2 }],
            totalAmount: 100,
            createdAt: '2024-08-24T00:15:17.650Z',
          },
        ];
        mockSalesRepository.getSalesByDate.mockResolvedValue(mockSales);
  
        const result = await saleService.getSalesByDate(mockDate);
  
        expect(result).toEqual(mockSales);
        expect(mockSalesRepository.getSalesByDate).toHaveBeenCalledWith(mockDate.toISOString());
      });
    });
  
    describe('deleteSale', () => {
      it('should delete a sale successfully', async () => {
        mockSalesRepository.deleteSale.mockResolvedValue(true);
  
        const result = await saleService.deleteSale('sale-id');
  
        expect(result).toBe(true);
        expect(mockLogger.info).toHaveBeenCalledWith('Sale with ID sale-id deleted');
      });
  
      it('should return false if sale not found', async () => {
        mockSalesRepository.deleteSale.mockResolvedValue(false);
  
        const result = await saleService.deleteSale('non-existent-id');
  
        expect(result).toBe(false);
        expect(mockLogger.error).toHaveBeenCalledWith('Sale with ID non-existent-id not found');
      });
    });
  });