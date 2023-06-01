// sign in context provider
"use client";

import { useState, createContext, useContext, Dispatch, SetStateAction } from 'react';

interface SignInContextProps {
    jwt: string;
    setJwt: Dispatch<SetStateAction<string>>;
    // isSignedIn: boolean;
    // setIsSignedIn: Dispatch<SetStateAction<boolean>>;setIsSignedIn
}

export const SignInContext = createContext<SignInContextProps>({
    jwt: 'none',
    setJwt: (): string => '',
    // isSignedIn: false
});

export const SignInContextProvider = ( { children } : { children: React.ReactNode}) => {
    const [ jwt, setJwt ] = useState('')
     return (
        <SignInContext.Provider value={{jwt, setJwt}}>
            { children }
        </SignInContext.Provider>
    );
};

export const useSignInContext = () => useContext(SignInContext);
