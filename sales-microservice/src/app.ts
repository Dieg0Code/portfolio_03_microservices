
import express from "express";
import { newDynamoDB } from "./db/config";
import { SalesRepositoryImpl } from "./repository/sales_repository_impl";
import { logger } from "./utils/logger";
import { SaleServiceImpl } from "./services/sale_service_impl";
import { SaleControllerImpl } from "./controllers/sale_controller_impl";
import { AppRouter } from "./routes/router";

const awsRegion = "sa-east-1";
const tableName = "sales";

const app = express();
app.use(express.json());

const db = newDynamoDB(awsRegion);
const saleRepo = new SalesRepositoryImpl(db, tableName, logger);
const saleService = new SaleServiceImpl(saleRepo, logger);
const saleController = new SaleControllerImpl(saleService, logger);
const router = new AppRouter(saleController);

app.use(router.initRoutes());

export default app;