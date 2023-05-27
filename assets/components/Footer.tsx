// footer

import UsefulLinks from "@components/UsefulLinks";
import Feedback from "@components/Feedback";
import Socials from "@components/Socials";

const Footer = () => (
    <footer className='footer flex justify-around py-4 items-center bg-neutral-900 text-slate-50'>
        <Feedback />
        <UsefulLinks />
        <Socials />
    </footer>
);

export default Footer;
