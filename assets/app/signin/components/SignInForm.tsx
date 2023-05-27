// form component
// remove console logs for handlers

"use client";
import { useState } from 'react';
import Button from '@components/Button';
import Input from '@components/Input';


const signInForm = () => {
    
    const [ email, setEmail ] = useState('');
    const [ password, setPassword ] = useState('');

    const handleSetEmail = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setEmail(event.target.value);
        console.log(email);
    };
    const handleSetPassword = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setPassword(event.target.value);
        console.log(password);
    };
    const handleSignIn = () => {
        console.log('signed in successfully!' + email, password)
    };
    
    return (
        <form className='flex flex-col gap-2 p-4 center max-w-max sm:max-w-l'>
            <label htmlFor='email'></label>
            <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetEmail(event)}
                type='text'
                name='email'
                placeholder='Enter email'
                id='email'
            />
            <label htmlFor='password'></label>
            <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetPassword(event)}
                type='text'
                name='password'
                placeholder='Enter password'
                id='password'
            />
            <Button onClick={handleSignIn} text='sign in' />
            <p>Add sign up / sign in toggle here! :)</p>
        </form>
    )
};

export default signInForm;
