// form component
// remove console logs for handlers

"use client";
import { useState } from 'react';
import Button from '@components/Button';
import Input from '@components/Input';
import axios from 'axios';
import { useRouter } from 'next/navigation';
import { useSignInContext } from '../../SignInContextProvider';
import Image from 'next/image';
import signin from '@/public/assets/images/undraw_mobile_login_re_9ntv.svg';


export const metadata = { title: 'sign in wigit' };

const signInForm = () => {
    
    const [ email, setEmail ] = useState('');
    const [ password, setPassword ] = useState('');

    const { jwt, setJwt, setRole, isSignedIn, setIsSignedIn, user, setUser } = useSignInContext();
    const router = useRouter();
    const url = "https://cheezaram.tech/api/v1/signin";


    const handleSetEmail = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setEmail(event.target.value);
    };
    const handleSetPassword = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setPassword(event.target.value);
    };
    const handleSignIn = () => {
        console.log('signed in successfully!' + email, password)
    };
    async function handleAxios (event: any){
        event.preventDefault();
        const credentials = { email, password };
        
        const { data, status } = await axios.post(url, credentials);
        if (status == 200) {
            setJwt(data.jwt);
            setRole(data.user.role);
            setIsSignedIn(true);
            // setUser(data.user);
            window.sessionStorage.setItem('jwt', data.jwt);
            window.sessionStorage.setItem('role', data.user.role);
            window.sessionStorage.setItem('isSignedIn', true);
            window.sessionStorage.setItem('user', JSON.stringify(data.user))
            console.log(data);
            console.log(isSignedIn);
            setUser(JSON.parse(window.sessionStorage.getItem('user')));
            console.log(user);
            console.log(window.sessionStorage.getItem('user'));
            router.push('/');
        }
    };
    const handleResetPassword = async () => {
        //event.preventDefault();
        await axios.post("https://cheezaram.tech/api/v1/reset_password", { email });
        router.push('/');
        alert("A password reset link has been sent to your email");
    };
    const pushToSignUp = (): void => {
      router.push('/signup');
    };

    return (
        <section className=' md:min-w-3xl md:flex flex-wrap rounded-lg shadow-md overflow-hidden'>
            <div className='md:w-1/2'>
                <Image 
                    src={ signin }
                    alt='Wigit Company Logo'
                    width={220}
                    height={300}/>
            </div>
            <form onSubmit={ handleAxios } className=' md:w-1/2 flex flex-col gap-2 p-4 bg-accent center max-w-max sm:max-w-l'>
                {/* <h1>Sign In</h1> */}
                <label htmlFor='email'></label>
                <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetEmail(event)}
                    type='text'
                    name='email'
                    placeholder='Enter email'
                    id='email'
                    required={ true }
                />
                <label htmlFor='password'></label>
                <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetPassword(event)}
                    type='password'
                    name='password'
                    placeholder='Enter password'
                    id='password'
                    required={ true }
                />
                <Button type='submit' text='sign in' />
                <p>Forgot password? <button className='underline pointer text-light_bg text-xs' onClick={handleResetPassword}>reset it here</button></p>
                <p>First time? <button className='underline pointer text-light_bg text-xs' onClick={pushToSignUp}>sign up :)</button></p>
            </form>
        </section>
    )
};

export default signInForm;