export class Product {
    productID: number;
    name: string;
    price: number;
    quantity: number;

    constructor(productID: number, name: string, price: number, quantity: number) {
        this.productID = productID;
        this.name = name;
        this.price = price;
        this.quantity = quantity;
    }
}