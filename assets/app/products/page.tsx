// The products page for the wigit web app

import ProductCard from './components/ProductCard';
import { Product } from './interfaces/product';

const url = "https://cheezaram.tech/api/v1/products";

export const metadata = { title: 'wigit products' };

async function getProducts(): Promise<any> {

  const res = await fetch(url, {
    headers: {"Content-Type": "application/json"},
    next: {"revalidate": 0}
  });

  const data = await res.json();
  if (res.ok) {
    return data.data;
  }
  return null; // fix this
}

export default async function Products() {
  const product_obj = await getProducts();
  return (
    <main>
      <div className='flex flex-col items-center justify-center'>
        <h1>Our wigs</h1>
        <p>Nothing but class....</p>
        { product_obj? 
        <div className="lg:max-w-4xl flex flex-wrap justify-center">
          { product_obj && product_obj.map((item: Product) => (
            <ProductCard { ...item } />
          ))}
        </div> :
        <p>no products</p>
          }
      </div>
    </main>
  )
}
