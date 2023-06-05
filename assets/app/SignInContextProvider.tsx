// sign in context provider
"use client";

import { useState, createContext, useContext, Dispatch, SetStateAction } from 'react';

interface SignInContextProps {
    jwt: string;
    setJwt: Dispatch<SetStateAction<string>>;
    role: string;
    setRole: Dispatch<SetStateAction<string>>;
    
    // isSignedIn: boolean;
    // setIsSignedIn: Dispatch<SetStateAction<boolean>>;setIsSignedIn
}

export const SignInContext = createContext<SignInContextProps>({
    jwt: 'not authorized',
    setJwt: (): string => '',
    role: 'customer',
    setRole: (): string => ''
    // isSignedIn: false
});

export const SignInContextProvider = ( { children } : { children: React.ReactNode}) => {
    const [ jwt, setJwt ] = useState('not authorized');
    const [ role, setRole ] = useState('Guest');
    

     return (
        <SignInContext.Provider value={{jwt, setJwt, role, setRole}}>
            { children }
        </SignInContext.Provider>
    );
};

export const useSignInContext = () => useContext(SignInContext);
