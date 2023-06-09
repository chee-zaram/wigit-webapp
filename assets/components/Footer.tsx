// footer

import UsefulLinks from "@components/UsefulLinks";
import Feedback from "@components/Feedback";
import Socials from "@components/Socials";

const Footer = () => (
    <footer className='footer flex justify-around py-4 h-[40vh] mt-10 items-center bg-dark_bg text-accent text-sm'>
        <Feedback />
        <UsefulLinks />
        <Socials />
    </footer>
);

export default Footer;
