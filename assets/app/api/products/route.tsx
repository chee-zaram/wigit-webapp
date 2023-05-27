//exposes api route for product
import { Product } from "@/app/products/interface/product";

export async function GET(): Promise<Response> {
    const productsList: Product[] = [
        {
            name: 'bone straight',
            product_id: 'mad01',
            category: 'wavy',
            stock: 10,
            price: 56,
            description: 'cool ass stuff',
            image_url: 'https://images.pexels.com/photos/13221796/pexels-photo-13221796.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1'
        },
        {
            name: 'blow dryer',
            product_id: 'mad02',
            category: 'accessory',
            stock: 12,
            description: '300 watts mini dryer',
            "price": 19,
            "image_url": "https://images.pexels.com/photos/973406/pexels-photo-973406.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
        },
        {
            name: 'loose waves',
            product_id: 'mad03',
            category: 'wavy',
            stock: 3,
            description: 'remi human hair',
            price: 99,
            image_url: "https://images.pexels.com/photos/1376042/pexels-photo-1376042.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
        },
        {
            name: 'coloured bob',
            product_id: 'mad04',
            category: 'straight',
            stock: 7,
            description: 'blue tinted bob',
            price: 69,
            image_url: "https://images.pexels.com/photos/15783139/pexels-photo-15783139/free-photo-of-model-in-wig-with-glamour-makeup.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
        },
        {
            name: 'bone straight',
            product_id: 'mad05',
            category: 'wavy',
            stock: 4,
            description: 'Bohemian curly waves',
            price: 100,
            image_url: "https://images.pexels.com/photos/7208632/pexels-photo-7208632.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
        },
        {
            name: 'indian straight',
            product_id: 'mad06',
            category: 'straight',
            stock: 54,
            description: 'straight hair',
            price: 100,
            image_url: "https://images.pexels.com/photos/7208632/pexels-photo-7208632.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
        }
    ];

    return new Response(JSON.stringify(productsList))
}
