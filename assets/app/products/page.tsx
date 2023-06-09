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
      <header className='flex flex-wrap w-[100vw]'>
        <div >
          {/* <img src='https://images.pexels.com/photos/13221796/pexels-photo-13221796.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1' /> */}
        </div>
        <div>
          <h1 className='font-bold uppercase text-accent'>Our luxury weaves</h1>
        </div>
      </header>
      <section className='bg-green-100 min-h-screen px-8 md:px-12 '>
        <div className='flex flex-col items-center justify-center'>
          <h1>Our wigs</h1>
          <p>Nothing but class....</p>
          { product_obj? 
          <div className="flex flex-wrap bg-slate-200 gap-4">
            { product_obj && product_obj.map((item: Product) => (
              <div key={product_obj.id}>
                <ProductCard { ...item } />
              </div>
            ))}
          </div> :
          <p>no products</p>
            }
        </div>
      </section>
    </main>
  )
}
