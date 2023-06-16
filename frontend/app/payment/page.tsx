// contact
"use client";
import { useRouter } from 'next/navigation';

const Contact = () => {
    const router = useRouter();
    
    return (
        <section>
            <h2>Payment instructions</h2>
            <div>
                <h3>Payment details</h3>
                <div>
                    <h3>Please use your reference number as the transaction's payment reference</h3>
                    <h4 className="mt-4">GtBank</h4>
                    <p>Account number: 01234579</p>
                    <p>Account name: Wigit company ltd</p>
                </div>
                <div>
                    <p>Thank you for shopping with us.. <span onClick={() => {router.push('/')}} className="cursor-pointer text-accent font-bold underline hover:text-dark_bg">Go back home</span></p>
                </div>
            </div>
        </section>
    );
};

export default Contact;
