import { IsArray, IsNumber, IsString, ValidateNested } from "class-validator";
import { Product } from "../../models/Product";
import { Type } from "class-transformer";

export class SaleResponse {

    @IsString()
    saleID: string;

    @IsNumber()
    userID: number;

    @IsArray()
    @ValidateNested({ each: true })
    @Type(() => Product)
    products: Product[];

    @IsNumber()
    totalAmount: number;

    @IsString()
    createdAt: string;

    constructor(saleID: string, userID: number, products: Product[], totalAmount: number, createdAt: string) {
        this.saleID = saleID;
        this.userID = userID;
        this.products = products;
        this.totalAmount = totalAmount;
        this.createdAt = createdAt;
    }
}