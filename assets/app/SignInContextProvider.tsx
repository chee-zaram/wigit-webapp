// sign in context provider
"use client";

import { SignInContext } from '@contexts/SignInContext';
import { useState } from 'react';

const SignInContextProvider = ( { children } : { children: React.ReactNode}) => {
    const [ jwt, setJwt ] = useState('');
    const [ isSignedIn, setIsSignedIn ] = useState(false);
    
    return (
        <SignInContext.Provider value={{jwt, setJwt, isSignedIn, setIsSignedIn}}>
            { children }
        </SignInContext.Provider>
    );
};

export default SignInContextProvider;
