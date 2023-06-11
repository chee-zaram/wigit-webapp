// form component
// remove console logs for handlers

"use client";
import { useState } from 'react';
import Button from '@components/Button';
import Input from '@components/Input';
import axios from 'axios';
import { useRouter } from 'next/navigation';
import Image from 'next/image';
import signupimg from '@/public/assets/images/signup.svg';


export const metadata = { title: 'sign up wigit' };

const SignUpForm = () => {
    
    const [ email, setEmail ] = useState('');
    const [ password, setPassword ] = useState('');
    const [ confirmPassword, setConfirmPassword ] = useState('');
    const [ firstName, setFirstName ] = useState('');
    const [ lastName, setLastName ] = useState('');
    const [ address, setAddress ] = useState('');
    const [ phoneNumber, setPhoneNumber ] = useState('');
    const router = useRouter();
    const url = "https://cheezaram.tech/api/v1/signup";
    const [ setJwt ] = useState();

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
    const pushToSignIn = (): void => {
      router.push('/signin');
    };
    
    async function handleSignUp (event: any){
        event.preventDefault();
        const newUser = { email, password, repeat_password: confirmPassword, first_name: firstName, last_name: lastName, phone: phoneNumber, address };
        try {
            const res = await axios.post(url, newUser);
            console.log(res);
        // setJwt(data.jwt);
        // on success, redirect to home page, on error, render error message
        if (res.status != 201) {
            alert(res.data.msg);
        }
            router.push('/signin');
        }
        catch(error) {
            alert(`${error}`)
        }
    }
    
    // async function prod() {
    //     const { data } = await axios.get(url + 'products', { 'headers': {'Authorization': 'Bearer ' + jwt} });
    //     console.log(data + 'new prods query');
    // }
    
   
    return (
    <section className=' md:min-w-5xl md:flex flex-wrap rounded-lg shadow-md overflow-hidden'>
        <div className='md:w-1/2 flexbox'>
            <Image 
                src={ signupimg }
                alt=''
                width={220}
                height={300}/>
        </div>
        <form onSubmit={ handleSignUp } className='md:w-1/2 flex flex-col gap-2 p-4 bg-accent center max-w-max sm:max-w-l'>
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
            <p className='text-sm'>Already shopping with us? <button className='underline pointer text-light_bg text-xs hover:text-dark_bg' onClick={pushToSignIn}>sign in</button></p>
        </form>
    </section>
    )
};

export default SignUpForm;
