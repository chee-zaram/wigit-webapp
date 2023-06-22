// Home page
import Image from 'next/image';
import headerImage from '@/public/assets/images/afro_girl.png';
import { getProducts } from '@app/products/page';
import { Product } from './products/interfaces/product';
import ProductCard from './products/components/ProductCard';
import Link from 'next/link';
import { ToastContainer } from 'react-toastify';


export default async function Home() {
  
  const trendingUrl = "https://cheezaram.tech/api/v1/products/categories/trending";
  const trendingProducts = await getProducts(trendingUrl); 
  
  return (
    <main className='home_page grid max-w-[100vw] mx-auto grid-rows-[repeat(10,_minmax(0,_1fr))] gap-4 md:gap-8 grid-flow-col min-h-screen'>
      <section className='flexbox row-span-2 home_section bg-[#E0DEDD]'>
        <div className='home_header flex w-full flex-wrap gap-2 justify-center sm:justify-around items-center'>
          <div className='bg-[#E0DEDD]'>
            <Image
            src={ headerImage }
            alt='girl on afro, smiling'
            width={400}
            height={1101}
            />
          </div>
          <div className='p-4 md:p-8 md:mr-16 home_header bg-yellow-500'>
            <h3 className='text-sky-900 text-3xl font-extrabold'>Amazing weave care deals</h3>
            <p className='text-sky-900 text-l'>Don't miss out on our discounts</p>
            <Link href='/products'><button className='bg-accent duration-300 shadow hover:bg-accent/40 px-4 py-1 capitalize rounded-full text-bg text-light_bg'>view deals</button></Link>
          </div>
        </div>
      </section>
      <section className='row-span-1 home_section'>
        <div className='flexbox_row gap-2 md:gap-4 p-4 md:p-8'>
          <button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>straight<i className='btn_icon mr-1'></i></button>
          <button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>wavy<i className='btn_icon mr-1'></i></button>
          <button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>accessory<i className='btn_icon mr-1'></i></button>
          <button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>trending<i className='btn_icon mr-1'></i></button>
          <button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>cheap<i className='btn_icon mr-1'></i></button>
          <button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>luxury<i className='btn_icon mr-1'></i></button>
        </div>
        <div className='flex flex-nowrap rounded-full overflow-hidden shadow max-w-[70px] bg-green-200'>
          <input type='text' placeholder='search' className='outline-accent px-4 bg-slate-100' />
          <button className='w-max'>search</button>
        </div>
      </section>
      <section className='flexbox max-w-[100vw] home_trending p-4 md:p-10 row-span-3 home_section  bg-accent/80' >
        <div className='mb-6'>
          <h2 className='text-sky-900 uppercase text-2xl  font-extrabold'>See what's trending</h2>
        </div>
        <div className=' flex gap-4 max-w-full overflow-x-scroll'>
          {
            trendingProducts && trendingProducts.map((item: Product) => (
              <div key={item.id}>
                < ProductCard { ...item } />
              </div>
            ))
          }
        </div>
      </section>
      <section className=' group row-span-4 home_section grid grid-rows-4 gap-4 grid-cols-4' >

          <div className='group-hover:blur group-hover:scale-90 hover:!scale-100 duration-500 hover:!blur-none row-span-1 col-span-4 md:row-span-2 md:col-span-2 home_section bg-gray-400'>accessories</div>
          <div className='group-hover:blur group-hover:scale-90 hover:!scale-100 duration-500 hover:!blur-none row-span-1 col-span-4 md:row-span-2 md:col-span-2 home_section bg-gray-500'>our services</div>
          <div className='group-hover:blur group-hover:scale-90 hover:!scale-100 duration-500 hover:!blur-none row-span-1 col-span-4 md:row-span-2 md:col-span-2 home_section bg-gray-500'>view events</div>
          <div className='group-hover:blur group-hover:scale-90 hover:!scale-100 duration-500 hover:!blur-none row-span-1 col-span-4 md:row-span-2 md:col-span-2 home_section bg-gray-700'>newsletter?? omo!!</div>
      </section>
      <ToastContainer />
    </main>

  )}
