// social links
import Link from 'next/link';

const Socials = () => (
    <aside className='flex md:flex-col gap-1'>
        <a href='' target='_blank'><i className='fab fa-github hover:text-light_bg duration-500 text-sm'></i></a>
        <a href='' target='_blank'><i className="fab fa-instagram hover:text-light_bg duration-500 text-sm"></i></a>
        <a href='' target='_blank'><i className="fab fa-twitter hover:text-light_bg duration-500 text-sm"></i></a>
        <a href='' target='_blank'><i className="fas fa-phone hover:text-light_bg duration-500 text-sm"></i></a>
        <a href='' target='_blank'><i className="fas fa-at hover:text-light_bg duration-500 text-sm"></i></a>
    </aside>
);

export default Socials;
