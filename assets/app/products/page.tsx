// The products page for the wigit web app

import ProductCard from './components/ProductCard';
import { Product } from './interface/product';
import Button from '@components/Button';

const url: string = "https://ovyevbodi-crispy-journey-jpj7vqj7r6xfqqvp-8000.preview.app.github.dev/products";
//"https://jsonplaceholder.typicode.com/todos";

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
      <div className='flex flex-col items-center justify-start'>
        <h1>Our wigs</h1>
        <p>Nothing but class....</p>
        <div className="inner">
          { data && data.map((item: Product) => (
            <ProductCard { ...item } />
          ))}
        </div>
      </div>
    </main>
  )
}
