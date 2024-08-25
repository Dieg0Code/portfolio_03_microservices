import { Router } from "express";
import { SaleController } from "../controllers/sale_controller";

export class AppRouter {
    private saleController: SaleController;

    constructor(saleController: SaleController) {
        this.saleController = saleController;
    }

    public initRoutes(): Router {
        const router = Router();

        router.get("/", (_req, res) => {
            res.json({ message: "Welcome to Sales Microservice" });
            
        });

        router.get("/health", (_req, res) => {
            res.json({ message: "Sales Microservice is healthy" });
        });

        const baseRoute = Router();

        const salesRoute = Router();

        salesRoute.get("/:saleID", this.saleController.getSaleByID.bind(this.saleController));
        salesRoute.get("/user/:userID", this.saleController.getSalesByUserID.bind(this.saleController));
        salesRoute.get("/date/:date", this.saleController.getSalesByDate.bind(this.saleController));
        salesRoute.get("/", this.saleController.getAllSales.bind(this.saleController));
        salesRoute.post("/", this.saleController.createSale.bind(this.saleController));
        salesRoute.delete("/:saleID", this.saleController.deleteSale.bind(this.saleController));

        baseRoute.use("/sales", salesRoute);

        router.use("/api/v1", baseRoute);

        return router;
        
    }
}