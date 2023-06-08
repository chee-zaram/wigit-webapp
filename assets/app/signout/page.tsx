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
        <main>
            <Button onClick={ handleSignOut } text='sign out' />
        </main>
    )
};

export default SignOut;
