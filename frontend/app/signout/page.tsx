// sign out page 
"use client";
import Button from '@components/Button';
import { useSignInContext } from '@app/SignInContextProvider';
import { useRouter } from 'next/navigation';

const SignOut = () => {
    const router = useRouter();
    const { setJwt, setIsSignedIn, setRole } = useSignInContext();

    const handleSignOut = () => {
        window.sessionStorage.clear();
        setIsSignedIn(false);
        setRole('guest');
        setJwt('not authorized');
        router.push('/');
    };
    
    return (
        <div className='logout min-h-[60vh] min-w-[60vw] flex items-center justify-center'>
            <div className=''>
                <Button onClick={ handleSignOut } text='sign out'/>
            </div>
        </div>
    )
};

export default SignOut;
