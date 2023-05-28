// form component
// remove console logs for handlers

"use client";
import { useState } from 'react';
import Button from '@components/Button';
import Input from '@components/Input';
import axios from 'axios';
import { useRouter } from 'next/navigation';

const SignUpForm = () => {
    
    const [ email, setEmail ] = useState('');
    const [ password, setPassword ] = useState('');
    const [ confirmPassword, setConfirmPassword ] = useState('');
    const [ firstName, setFirstName ] = useState('');
    const [ lastName, setLastName ] = useState('');
    const [ address, setAddress ] = useState('');
    const [ phoneNumber, setPhoneNumber ] = useState('');
    const router = useRouter();

    const handleSetEmail = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setEmail(event.target.value);
    };
    const handleSetPassword = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setPassword(event.target.value);
    };
    const handleSetConfirmPassword = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setConfirmPassword(event.target.value);
    };
    const handleSetFirstName = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setFirstName(event.target.value);
    };
    const handleSetLastName = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setLastName(event.target.value);
    };
    const handleSetAddress = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setAddress(event.target.value);
    };
    const handleSetPhoneNumber = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setPhoneNumber(event.target.value);
    };
    

    async function handleSignUp (event: any){
        event.preventDefault();
        const newUser = { email, password, confirmPassword, firstName, lastName, phoneNumber, address };
        const { data } = await axios.post("https://cheezaram.tech/api/v1/signup", newUser);
        console.log(data ? data : "error loading data...");
        // on success, redirect to home page, on error, render error message
        router.push('/');
        console.log(newUser);
    } 
   
    return (
        <form onSubmit={ handleSignUp } className='flex flex-col gap-2 p-4 center max-w-max sm:max-w-l'>
            <h1>Sign Up</h1>
            <label htmlFor='email'></label>
            <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetEmail(event)}
                type='text'
                name='email'
                placeholder='Enter email'
                id='email'
                autocomplete='on'
                required={ true }
            />
            <label htmlFor='password'></label>
            <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetPassword(event)}
                type='password'
                name='password'
                placeholder='Enter password'
                id='password'
                autocomplete='off'
                required={ true }
            />
            <label htmlFor='confirm_password'></label>
            <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetConfirmPassword(event)}
                type='password'
                name='confirm password'
                placeholder='confirm password'
                id='confirm_password'
                autocomplete='off'
                required={ true }
            />
            <label htmlFor='first_name'></label>
            <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetFirstName(event)}
                type='text'
                name='first name'
                placeholder='Enter first name'
                id='first_name'
                autocomplete='on'
                required={ true }
            />
            <label htmlFor='last_name'></label>
            <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetLastName(event)}
                type='text'
                name='last name'
                placeholder='Enter last name'
                id='last_name'
                autocomplete='on'
                required={ true }
            />
            <label htmlFor='phone_number'></label>
            <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetPhoneNumber(event)}
                type='tel'
                name='phone number'
                placeholder='Enter phone number'
                id='phone_number'
                autocomplete='on'
                required={ true }
            />
            <label htmlFor='address'></label>
            <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetAddress(event)}
                type='text'
                name='address'
                placeholder='Enter address'
                id='address'
                autocomplete='on'
                required={ true }
            />
            <Button type='submit' text='sign up' />
        </form>
    )
};

export default SignUpForm;
