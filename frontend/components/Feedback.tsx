// fedback form
"use client";
import Button from '@components/Button';
import Input from '@components/Input';
import { useState } from 'react';
// button form input

const Feedback = () =>  {
    const  [ email, setEmail ] = useState('');
    const handleSetEmail = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setEmail(event.target.value);
    };

    return(
    <form className='flexbox rounded bg-accent/30 m-2 p-2' >
        <h3 className='text-accent/70 p-2 font-bold uppercase'>We'd love to hear from you</h3>
        <div className='flexbox p-3 w-full'>
            <input className='pb-1 pl-3 outline-none duration-300 text-light_bg bg-transparent w-full border-b border-b-light_bg/70' onChange={handleSetEmail} type='text' name='email' id='email' placeholder='Enter email' />
            <textarea className='pt-3 pl-3 mb-4 outline-none duration-300 text-light_bg bg-transparent w-full border-b border-b-light_bg/70' placeholder='Your message here'></textarea>
            <Button text='Reach out' type='submit' />
        </div>
    </form>
);
}

export default Feedback;
