// Dashboard
"use client";

import { useState, useEffect } from 'react';
import { useSignInContext } from '@app/SignInContextProvider';
import axios from 'axios';
import Button from '@components/Button';

const Dashboard = async () => {

    const { jwt, setJwt, role, setRole } = useSignInContext();

    if (typeof window !== 'undefined') {
    if (window.sessionStorage.getItem('jwt')) {
        setJwt(window.sessionStorage.getItem('jwt'));
        setRole(window.sessionStorage.getItem('role'));

    }
    }
    
    return (
        <main className='grid md:grid-rows-3'>
            <div className='bg-slate-600'><h2>welcome to your dashboard<br/> fucking awesome shit..</h2></div>
            <div></div>
            <div></div>
            <div></div>
            <div></div>
            <div></div>
            <div></div>
            <div></div>
            <div></div>

        </main>
    )
};

export default Dashboard;