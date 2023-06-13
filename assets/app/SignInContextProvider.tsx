// sign in context provider
"use client";

import { useState, createContext, useContext, Dispatch, SetStateAction } from 'react';

interface SignInContextProps {
    jwt: string;
    setJwt: Dispatch<SetStateAction<string>>;
    role: string;
    setRole: Dispatch<SetStateAction<string>>;
    isSignedIn: boolean;
    setIsSignedIn: Dispatch<SetStateAction<boolean>>;
    user: object;
    setUser: Dispatch<SetStateAction<string>>;
}

export const SignInContext = createContext<SignInContextProps>({
    jwt: 'not authorized',
    setJwt: (): string => '',
    role: 'customer',
    setRole: (): string => '',
    isSignedIn: false,
    setIsSignedIn: (): boolean => false,
    user: {},
    setUser: ():any => {}
});

export const SignInContextProvider = ( { children } : { children: React.ReactNode}) => {
    const [ jwt, setJwt ] = useState('not authorized');
    const [ role, setRole ] = useState('Guest');
    const [ isSignedIn, setIsSignedIn ] = useState(false);
    const [ user, setUser ] = useState({});

     return (
        <SignInContext.Provider value={{jwt, setJwt, role, setRole, isSignedIn, setIsSignedIn, user, setUser}}>
            { children }
        </SignInContext.Provider>
    );
};

export const useSignInContext = () => useContext(SignInContext);
