import { Product } from "./Product";

export interface Sale {
    saleID: string;
    userID: number;
    products: Product[];
    totalAmount: number;
    date: string;
}