// footer

import UsefulLinks from "@components/UsefulLinks";
import Feedback from "@components/Feedback";
import Socials from "@components/Socials";

const Footer = () => (
    <footer className='footer min-h-[40vh] mt-8 max-w-[100vw] text-light_bg/50 text-xs'>
        <section className='flex flex-col md:flex-row justify-center md:justify-around flex-wrap gap-7 py-4 items-center '>
            <Feedback />
            <UsefulLinks />
            <Socials />
        </section>
        <section className='border-t py-2 border-light_bg/30'>
            <h3 className='font-bold'>Developers</h3>
            <span className='mr-4 pt-2'>
                <a>github</a>
                <a href={'mailto:evbodiovo@gmail.com'}>Ovy Evbodi</a>
            </span>
            <span>
                <a>github</a>
                <a href={'mailto:ecokeke21@gmail.com'}>Chee-zaram Okeke</a>
            </span>
        </section>
    </footer>
);

export default Footer;
