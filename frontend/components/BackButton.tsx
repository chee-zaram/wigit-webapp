// Round back button
"use client";
import { useRouter } from 'next/navigation';

const BackButton = () => {
    const router = useRouter();
    
    return (
    <div onClick={() => {router.back()}} className='mb-6 hover:bg-accent/80 text-right ml-[10vw] duration-300 rounded-full p-3 max-w-max'>
       <svg xmlns="http://www.w3.org/2000/svg" height="40" viewBox="0 -960 960 960" width="40"><path d="M480-160 160-480l320-320 42 42-248 248h526v60H274l248 248-42 42Z"/></svg> 
    </div>
    );
};

export default BackButton;
