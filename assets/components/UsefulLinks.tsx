// nav links for the footer
import Link from 'next/link';

const UsefulLinks = () => (
    <section>
        <nav className='flex flex-col justify-between'>
          <Link href='/' className=' hover:underline  hover:text-[#AB841F]'>Home</Link>
          <Link href='/about' className=' hover:underline  hover:text-[#AB841F]'>About</Link>
          <Link href='/products' className=' hover:underline  hover:text-[#AB841F]'>Our wigs</Link>
        </nav>
    </section>
);

export default UsefulLinks;