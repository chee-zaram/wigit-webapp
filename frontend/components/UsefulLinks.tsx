// nav links for the footer
import Link from 'next/link';

const UsefulLinks = () => (
    <section>
        <nav className='flex flex-col items-start justify-between'>
          <Link href='/' className=' duration-500 hover:underline  hover:text-light_bg'>Home</Link>
          <Link href='/about' className='duration-500 hover:underline  hover:text-light_bg'>About</Link>
          <Link href='/products' className='duration-500 hover:underline  hover:text-light_bg'>Our wigs</Link>
          <Link href='/about' className='duration-500 hover:underline  hover:text-light_bg'>Contact us</Link>
          <Link href='/services' className='duration-500 hover:underline  hover:text-light_bg'>Our services</Link>
          <Link href='/payment' className='duration-500 hover:underline  hover:text-light_bg'>Payment instructions</Link>

          {/* <Link href='/services' className=' hover:underline  hover:text-light_bg'>Services</Link> */}
        </nav>
    </section>
);

export default UsefulLinks;