// App logo

import Image from "next/image";
import Link from 'next/link';

const Logo = () => (
    <div className='flex gap-2 flex-center'>
      <Link href='/'>
        <Image
          src='/assests/images/logo.svg'
          alt='Wigit Company Logo'
          width={50}
          height={50}
        />
        <p className='logo_text'>Class with sass...</p>
      </Link>
    </div>
);

export default Logo;
