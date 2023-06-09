// nav links for the footer
import Link from 'next/link';

const UsefulLinks = () => (
    <section>
        <nav className='flex flex-col justify-between'>
          <Link href='/' className=' hover:underline  hover:text-light_bg'>Home</Link>
          <Link href='/about' className=' hover:underline  hover:text-light_bg'>About</Link>
          <Link href='/products' className=' hover:underline  hover:text-light_bg'>Our wigs</Link>
          <Link href='/contact' className=' hover:underline  hover:text-light_bg'>Contact us</Link>
          <Link href='/services' className=' hover:underline  hover:text-light_bg'>Services</Link>
        </nav>
    </section>
);

export default UsefulLinks;