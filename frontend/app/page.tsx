// Home page
import Image from 'next/image';
import headerImage from '@/public/assets/images/afro_girl.png';
import { getProducts } from '@app/products/page';
import { Product } from '@app/products/interfaces/product';
import ProductCard from '@app/products/components/ProductCard';
import Link from 'next/link';
import { ToastContainer } from 'react-toastify';
import ProductSearchBox from '@components/ProductSearchBox';


export default async function Home() {
  const url = "https://cheezaram.tech/api/v1/products";
  const trendingUrl = "https://cheezaram.tech/api/v1/products/categories/trending";
  const trendingProducts = await getProducts(trendingUrl); 
  const searchObj = {url, tag: 'name'};
  
  return (
    <main className='home_page grid max-w-[100vw] mx-auto grid-rows-[repeat(10,_minmax(0,_1fr))] gap-4 md:gap-8 grid-flow-col min-h-screen'>
      <section className='flexbox home_header header_wrap row-span-2 home_section'>
        <div className=' flex w-full flex-wrap gap-2 justify-center sm:justify-around items-center'>
          <div className=''>
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
          <Link href={'/straight'}><button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>straight<i className='btn_icon mr-1'></i></button></Link>
          <Link href={'/wavy'}><button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>wavy<i className='btn_icon mr-1'></i></button></Link>
          <Link href={'/accessories'}><button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>accessory<i className='btn_icon mr-1'></i></button></Link>
          <Link href={'/'}><button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>trending<i className='btn_icon mr-1'></i></button></Link>
          <button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>cheap<i className='btn_icon mr-1'></i></button>
          <button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>luxury<i className='btn_icon mr-1'></i></button>
        </div>
        <ProductSearchBox { ...searchObj} />
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
      <section className='group capitalize row-span-4 home_section grid grid-rows-4 gap-4 grid-cols-4' >
          <div className='wigs outer_group group-hover:blur  group-hover:scale-90 hover:!scale-100 duration-500 hover:!blur-none row-span-1 col-span-4 md:row-span-2 md:col-span-2 home_section '>
            <div className='section_group'>
              <div className='inner_group_text'>
                <h3 className='text-2xl'>Our wigs</h3>
                <p className='text-xs'>Experience luxury when you wear wigit</p>
                <Link href={'/products'}><button className='py-2 px-6 border mt-1 hover:scale-105 duration-300 text-dark_bg/70 hover:underline border-accent bg-light_bg'>shop now</button></Link>
              </div>
            </div>
          </div>
          <div className='services outer_group group-hover:blur group-hover:scale-90 hover:!scale-100 duration-500 hover:!blur-none row-span-1 col-span-4 md:row-span-2 md:col-span-2 home_section'>
            <div className='section_group'>
              <div className='inner_group_text'>
                <h3 className='text-2xl'>Our services</h3>
                <p className='text-xs'>Let us show you what real pampering is</p>
                <Link href={'/services'}><button className='py-2 px-6 border mt-1 hover:scale-105 duration-300 text-dark_bg/70 hover:underline border-accent bg-light_bg'>Book now</button></Link>
              </div>
            </div>
          </div>
          <div className='accessories outer_group group-hover:blur  group-hover:scale-90 hover:!scale-100 duration-500 hover:!blur-none row-span-1 col-span-4 md:row-span-2 md:col-span-2 home_section'>
            <div className='section_group'>
              <div className='inner_group_text'>
                <h3 className='text-2xl'>accessories</h3>
                <p className='text-xs'>Gentle care tools to keep your wigs alive</p>
                <Link href={'/accessories'}><button className='py-2 px-6 border mt-1 hover:scale-105 duration-300 text-dark_bg/70 hover:underline border-accent bg-light_bg'>View tools</button></Link>
              </div>
            </div>
          </div>
          <div className='about outer_group group-hover:blur group-hover:scale-90 hover:!scale-100 duration-500 hover:!blur-none row-span-1 col-span-4 md:row-span-2 md:col-span-2 home_section'>
            <div className='section_group'>
              <div className='inner_group_text'>
                <h3 className='text-2xl'>About us</h3>
                <p className='text-xs text-dark_bg'>We are so much more than this box can take</p>
                <Link href={'/about'}><button className='py-2 px-6 border mt-1 hover:scale-105 duration-300 text-dark_bg/70 hover:underline border-accent bg-light_bg'>know us</button></Link>
              </div>
            </div>
          </div>
      </section>
      <ToastContainer />
    </main>

  )}
