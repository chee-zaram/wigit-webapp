// The products page for the wigit web app

import ProductCard from './components/ProductCard';
import { Product } from './interfaces/product';

const url: string = "https://jel1cg-8000.csb.app/products";

export const metadata = { title: 'wigit products' };

async function getProducts(): Promise<any> {

  const res = await fetch(url, {
    headers: {"Content-Type": "application/json",
    }
  });

  const data = await res.json();
  if (res.ok) {
    return data;
  }
  return null; // fix this
}

export default async function Products() {
  const data = await getProducts();
  return (
    <main>
      <div className='flex flex-col items-center justify-center'>
        <h1>Our wigs</h1>
        <p>Nothing but class....</p>
        <div className="lg:max-w-4xl flex flex-wrap justify-center bg-neutral-400">
          { data && data.map((item: Product) => (
            <ProductCard { ...item } />
          ))}
        </div>
      </div>
    </main>
  )
}
