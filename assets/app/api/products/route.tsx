//exposes api route for product
import { Product } from "@/app/products/interface/product";

export async function GET(request: Request) {
    const productsList: Product[] = [
        {
            name: 'bone straight',
            product_id: 'mad01',
            category: 'wavy',
            stock: 10,
            description: 'cool ass stuff'
        },
        {
            name: 'blow dryer',
            product_id: 'mad02',
            category: 'accessory',
            stock: 12,
            description: '300 watts mini dryer'
        },
        {
            name: 'loose waves',
            product_id: 'mad03',
            category: 'wavy',
            stock: 3,
            description: 'remi human hair'
        },
        {
            name: 'coloured bob',
            product_id: 'mad04',
            category: 'straight',
            stock: 7,
            description: 'blue tinted bob'
        },
        {
            name: 'bone straight',
            product_id: 'mad05',
            category: 'wavy',
            stock: 4,
            description: 'Bohemian curly waves'
        },
        {
            name: 'indian straight',
            product_id: 'mad06',
            category: 'straight',
            stock: 54,
            description: 'straight hair'
        }
    ];

    return new Response(JSON.stringify(productsList))
}
