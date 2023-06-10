
import Image from 'next/image';
import headerImage from '@/public/assets/images/afro_girl.png';
import { getProducts } from '@app/products/page';
import { Product } from './products/interfaces/product';
import ProductCard from './products/components/ProductCard';


export default async function Home() {
  
  const trendingUrl = "https://cheezaram.tech/api/v1/products/categories/trending";
  const trendingProducts = await getProducts(trendingUrl); 
  
  return (
    <main className='grid max-w-[75vw] mx-auto grid-rows-[repeat(10,_minmax(0,_1fr))] gap-4 md:gap-8 grid-flow-col min-h-screen'>
      <section className='flexbox row-span-3 home_section bg-[#E0DEDD]'>
        <div className='flex flex-wrap gap-2 justify-center sm:justify-between items-center'>
          <div className='bg-[#E0DEDD]'>
            <Image
            src={ headerImage }
            alt='girl on afro, smiling'
            width={400}
            height={1101}
            />
          </div>
          <div className='p-4 md:p-8 md:mr-16 bg-yellow-500'>
            <h3 className='text-sky-900 text-3xl font-extrabold'>Amazing weave care deals</h3>
            <p className='text-sky-900 text-l'>Don't miss out on our discounts</p>
            <button className='bg-accent duration-300 shadow hover:bg-accent/40 px-4 py-1 capitalize rounded-full text-bg text-light_bg'>view deals</button>
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
        <div className='flex flexl-nowrap rounded-fulloverflow-hidden  shadowmax- w-[70px bg-green-200]'>
          <input type='text' placeholder='search' className='outline-accent px-4 bg-slate-100' />
          <button className='w-max'>search</button>
        </div>
      </section>
      <section className=' p-4 md:px-10 md:py-8 row-span-3 home_section  bg-accent/80' >
        <div>
          <h2 className='text-sky-900 uppercase text-2xl font-extrabold'>See what's trending</h2>
        </div>
        <div className=' flex gap-4 max-w-[70vw] p-4  overflow-x-scroll'>
          {
            trendingProducts && trendingProducts.map((item: Product) => (
              < ProductCard { ...item } />
            ))
          }
        </div>
      </section>
      <section className='row-span-3 home_section grid grid-rows-4 grid-cols-4' >

          <div className='row-span-1 col-span-- md:row-span-2 md:col-span-2 home_section bg-red-100'>hi</div>
          <div className='row-span-1 md:row-span-2 md:col-span-2 home_section bg-red-200'></div>
          <div className='row-span-1 md:row-span-2 md:col-span-2 home_section bg-red-300'></div>
          <div className='row-span-1 md:row-span-2 md:col-span-2 home_section bg-red-400'></div>
      </section>
    </main>

  )}
