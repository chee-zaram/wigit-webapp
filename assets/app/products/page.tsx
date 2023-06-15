// The products page for the wigit web app

import ProductCard from './components/ProductCard';
import { Product } from './interfaces/product';

const productsUrl = "https://cheezaram.tech/api/v1/products";
const trendingUrl = "https://cheezaram.tech/api/v1/products/categories/trending";
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';


export const metadata = { title: 'wigit products' };

export async function getProducts(url: string): Promise<any> {

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
  const product_obj = await getProducts(productsUrl);
  const trendingProdsObj = await getProducts(trendingUrl);

  
  return (
    <section className=' products_page min-h-screen'>
      <header className='products_header overlay bg-dark_bg/20 min-h-[30vh]'>
        <div className='h-full w-[100vw]'>
          <div className=' w-full py-4 flex flex-col text-center min-h-full'>
            <h1 className='font-bold uppercase text-lg mx-auto text-accent'>Our luxury weaves</h1>
            <p className='text-sm text-dark_bg'>Nothing but class....</p>
          </div>
        </div>
      </header>
      <section className=' min-h-screen w-[100vw] bg-light_bg'>
        {/* <div className='min-h-[15vh] flexbox bg-dark_bg/20d'>
          <h3>Trending</h3>
        </div>
        <div className='flex flex-col items-center justify-center p-8 md:px-12'>
          { trendingProdsObj? 
          <div className="flex flex-wrap gap-4">
            { trendingProdsObj && trendingProdsObj.map((item: Product) => (
              <div key={trendingProdsObj.id}>
                <ProductCard { ...item } />
              </div>
            ))}
          </div> :
          <p>Every product is trending!!!</p>
            }
        </div> */}
        <div className='flex flex-col items-center justify-center p-8 md:px-12'>
          { product_obj? 
          <div className="flexbox_row gap-4 md:gap-8">
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
      <div className='h-[15vh] flexbox'>
        <p className='text-accent  text-center font-bold'>More coming soon...</p>
      </div>
    <ToastContainer />
    </section>
  )
}
