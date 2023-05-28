// sign in / sign up page
"use client";
import { useState } from 'react';
import SignInForm from '@app/signin/components/SignInForm';
import { useRouter } from 'next/navigation';
import axios from 'axios';

const signin = () => {
    // check if user is signed in
    const [ isSignedIn, setIsSignedIn ] = useState(false);
    const router = useRouter();
    
    // const { data, error } = useQuery({ queryKey: ['signInSubmit'], queryFn: handleAxios})
    // console.log(data);
    // async function handleAxios () {
    //     const { data } = await axios.post("https://jel1cg-8000.csb.app/signin", { headers: {"Authorization": "newBossVee", "Content-Type": "Application/json"}, token: 'something sent from the client side'});
    //     console.log(data ? data : error);
    //     setIsSignedIn(true);
    // }
    
    const toggleSignIn = (): void => {
      router.push('/signup');
    };

    return (
        <main className='signin_main flex flex-col justify-around items-center'>
            {/* take this to rootlayout to conditionally render sign in link  */}
            <div className='welcome_message'>
                { isSignedIn ? 
                <h3>Welcome back, Vee baby!</h3> : 
                <h3>Hey there, we're glad you found us, Please sign in</h3> }
            </div>
            <SignInForm />
            <p>First time? <span className='underline pointer text-accent font-extrabold' onClick={toggleSignIn}>sign up</span> :)</p>
        </main>
    )
    
};

export default signin;