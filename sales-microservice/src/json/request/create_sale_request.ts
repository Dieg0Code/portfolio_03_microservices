import { IsArray, IsNumber, ValidateNested } from "class-validator";
import { Product } from "../../models/Product";
import { Type } from "class-transformer";

export class CreateSaleRequest {

    @IsNumber()
    userID: number;

    @IsArray()
    @ValidateNested({ each: true })
    @Type(() => Product)
    products: Product[];
    totalAmount: number;

    constructor(userID: number, products: Product[], totalAmount: number) {
        this.userID = userID;
        this.products = products;
        this.totalAmount = totalAmount;
    }
}