import Image from 'next/image';
import headerImage from '@/public/assets/images/afro_girl.png';

export default function Home() {
  return (
    <main className='grid grid-rows-[repeat(10,_minmax(0,_1fr))] gap-4 md:gap-8 grid-flow-col min-h-screen px-8 md:px-12'>
      <section className='row-span-3 home_section bg-[#E0DEDD]'>
        <div className='flex flex-wrap gap-2 justify-center sm:justify-between items-center'>
          <div className='bg-[#E0DEDD]'>
            <Image
            src={ headerImage }
            alt='girl on afro, smiling'
            width={400}
            height={1101}
            />
          </div>
          <div className='p-4 md:p-8 md:mr-16'>
            <h3 className='text-sky-900 text-xxl font-extrabold'>We offer amazing weave care deals</h3>
            <p className='text-sky-900 text-l'>Don't miss out on our discounts</p>
            <button className='bg-accent px-4 py-1 capitalize rounded-full text-bg text-light_bg'>view deals</button>
          </div>
        </div>
      </section>
      <section className='row-span-1 home_section'>
        <div className='flexbox_row gap-2 md:gap-4 p-4 md:p-8'>
          <button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>straight<i className='btn_icon mr-1'></i></button>
          <button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>straight<i className='btn_icon mr-1'></i></button>
          <button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>straight<i className='btn_icon mr-1'></i></button>
          <button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>straight<i className='btn_icon mr-1'></i></button>
          <button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>straight<i className='btn_icon mr-1'></i></button>
          <button className='font-bold text-xs text-slate-700 py-1 px-4 shadow rounded-full bg-slate-300/80'>straight<i className='btn_icon mr-1'></i></button>
        </div>
        <div className='flex flexl-nowrap rounded-fulloverflow-hidden  shadowmax- w-[70px bg-green-200]'>
          <input type='text' placeholder='search' className='outline-accent px-4 bg-slate-100' />
          <button className='w-max'>search</button>
        </div>
      </section>
      <section className='row-span-3 home_section  bg-blue-600' >
        <div>
          <h2>trending carousel</h2>
        </div>
      </section>
      <section className='row-span-3 home_section bg-blue-800' >
        <div>
          <h2>group section</h2>
        </div>
      </section>
    </main>

  )}
