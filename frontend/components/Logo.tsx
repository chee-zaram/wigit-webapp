// App logo

import Image from "next/image";
import logo from '/public/assets/images/wigit.png';

const Logo = () => (
    <div className='logo_wrap p-3'>
        <Image
          src={logo}
          alt='Wigit Company Logo'
          width={70}
          height={40}
        />
    </div>
);

export default Logo;
