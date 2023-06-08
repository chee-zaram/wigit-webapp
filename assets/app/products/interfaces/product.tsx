// Defines a product

export interface Product {
    name: string;
    id: string;
    category: "accessory" | "wavy" | "straight";
    stock: number;
    description: string;
    price: number;
    image_url: string;
}
