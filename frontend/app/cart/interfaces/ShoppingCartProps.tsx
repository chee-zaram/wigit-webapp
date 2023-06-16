// shopping cart interface

export default interface Item {
    id: string;
    amount: string;
    quantity: number;
    product: {name: string; price:string; image_url:string; }
}