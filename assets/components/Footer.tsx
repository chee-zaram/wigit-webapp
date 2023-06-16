// footer

import UsefulLinks from "@components/UsefulLinks";
import Feedback from "@components/Feedback";
import Socials from "@components/Socials";

const Footer = () => (
    <footer className='footer max-w-[100vw] flex flex-col md:flex-row justify-center md:justify-around flex-wrap gap-7 py-4 min-h-[40vh] mt-8 items-center bg-dark_bg text-accent/50 text-xs'>
        <Feedback />
        <UsefulLinks />
        <Socials />
    </footer>
);

export default Footer;
