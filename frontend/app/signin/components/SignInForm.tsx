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
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

export const metadata = { title: 'sign in wigit' };

const signInForm = () => {
    
    const [ email, setEmail ] = useState('');
    const [ password, setPassword ] = useState('');

    const { setJwt, setRole, isSignedIn, setIsSignedIn, user, setUser } = useSignInContext();
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
    
    async function handleSignIn (event: any){
        event.preventDefault();
        const credentials = { email, password };
        
        try {
            const { data, status } = await axios.post(url, credentials);
        if (status == 200) {
            setJwt(data.jwt);
            setRole(data.user.role);
            setIsSignedIn(true);
            setUser(data.user);
            window.sessionStorage.setItem('jwt', data.jwt);
            window.sessionStorage.setItem('role', data.user.role);
            window.sessionStorage.setItem('isSignedIn', JSON.stringify(true));
            window.sessionStorage.setItem('user', JSON.stringify(data.user))
            toast.success('sign in successful!', {
                position: "top-center",
                autoClose: 3000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                theme: "light",
            });
            router.push('/');
        }
        }
        catch {
            toast.error('something went wrong, please check your credentials, and try again.', {
                position: "top-center",
                autoClose: 5000,
                hideProgressBar: true,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                theme: "colored",
            });
        }
    };
    const handleResetPassword = async () => {
        //event.preventDefault();
        try {
            await axios.post("https://cheezaram.tech/api/v1/reset_password", { email });
            toast.success('A password reset link has been sent to your email!', {
                position: "top-center",
                autoClose: 3000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                theme: "light",
            });
            router.push('/');
        } catch (error) {
            toast.error("We didn't find your email in our records, please sign up", {
                position: "top-center",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                theme: "colored",
            });
        }
    };
    const pushToSignUp = (): void => {
      router.push('/signup');
    };

    return (
        <section className=' md:min-w-5xl md:flex flex-wrap rounded-lg shadow-md overflow-hidden'>
            <div className='md:w-1/2 flexbox'>
                <Image 
                    src={ signin }
                    alt=''
                    width={220}
                    height={300}/>
            </div>
            <div className='bg-accent py-4 px-2 md:w-1/2'>
                <form onSubmit={ handleSignIn } className='flex flex-col gap-2 p-4  center max-w-max sm:max-w-l'>
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
                </form>
                <p className='text-sm px-2'>Forgot password? <button className='underline pointer text-light_bg text-xs hover:text-dark_bg' onClick={handleResetPassword}>Reset it here</button></p>
                <p className='text-sm px-2'>First time? <button className='underline pointer text-light_bg text-xs hover:text-dark_bg' onClick={pushToSignUp}>Sign Up :)</button></p>
            </div>
            <ToastContainer />
        </section>
    )
};

export default signInForm;