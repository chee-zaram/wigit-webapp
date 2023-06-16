// nav links for the footer
import Link from 'next/link';

const UsefulLinks = () => (
    <section>
        <nav className='flex flex-col justify-between'>
          <Link href='/' className=' duration-500 hover:underline  hover:text-light_bg/60'>Home</Link>
          <Link href='/about' className='duration-500 hover:underline  hover:text-light_bg/60'>About</Link>
          <Link href='/products' className='duration-500 hover:underline  hover:text-light_bg/60'>Our wigs</Link>
          <Link href='/contact' className='duration-500 hover:underline  hover:text-light_bg/60'>Contact us</Link>
          {/* <Link href='/services' className=' hover:underline  hover:text-light_bg'>Services</Link> */}
        </nav>
    </section>
);

export default UsefulLinks;