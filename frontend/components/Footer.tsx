// footer

import UsefulLinks from "@components/UsefulLinks";
import Feedback from "@components/Feedback";
import Socials from "@components/Socials";

const Footer = () => (
    <footer className='px-8 footer min-h-[40vh] mt-8 max-w-[100vw] text-light_bg/50 text-xs'>
        <section className='flex flex-col md:flex-row justify-center md:justify-between flex-wrap gap-6 py-4 items-center'>
            <Feedback />
            <UsefulLinks />
            <Socials />
        </section>
        <section className='border-t py-2 border-light_bg/30'>
            <h3 className='font-bold mb-1'>Developers</h3>
            <span className='mr-4'>
                <a href='https://github.com/OvyEvbodi' target='_blank'><i className='fab fa-github hover:text-light_bg duration-500'></i></a>
                <a href={'mailto:evbodiovo@gmail.com'} className='ml-2 hover:text-light_bg duration-500'>Ovy Evbodi</a>
            </span>
            <span>
                <a href='https://github.com/chee-zaram' target='_blank'><i className='fab fa-github hover:text-light_bg duration-500'></i></a>
                <a href={'mailto:ecokeke21@gmail.com'} className='ml-2 hover:text-light_bg duration-500'>Chee-zaram Okeke</a>
            </span>
            <p className='text-light_bg/10 pt-2'>&copy; Wigit 2023, all rights reserved.</p>
        </section>
    </footer>
);

export default Footer;
