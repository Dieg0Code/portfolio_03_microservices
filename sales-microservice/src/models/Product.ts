export class Product {
    productID: string;
    name: string;
    price: number;
    quantity: number;

    constructor(productID: string, name: string, price: number, quantity: number) {
        this.productID = productID;
        this.name = name;
        this.price = price;
        this.quantity = quantity;
    }
}