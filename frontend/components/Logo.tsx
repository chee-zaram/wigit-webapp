// App logo

import Image from "next/image";
import logo from '/public/assets/images/wigit.png';
import Link from 'next/link';

const Logo = () => (
    <div className='logo_wrap max-w-max p-3'>
        <Link href={'/'}><Image
          src={logo}
          alt='Wigit Company Logo'
          width={70}
          height={40}
        />
        </Link>
    </div>
);

export default Logo;
