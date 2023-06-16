// orders (for admins)
"use client";
import { useState, useEffect } from 'react';
import { useSignInContext } from '@app/SignInContextProvider';
import axios from 'axios';
import Button from '@components/Button';
import { NextPage } from 'next';

const Orders: NextPage<any> = async(props: any) => {
    // retrieve all orders
    
    const url = 'https://cheezaram.tech/api/v1/admin/orders';
    const { jwt, role, setJwt, setRole } = useSignInContext();
    
    if (typeof window !== 'undefined') {
        if (window.sessionStorage.getItem('jwt')) {
            setJwt(window.sessionStorage.getItem('jwt'));
            setRole(window.sessionStorage.getItem('role'));

        }
    };
    
    
    return (
        <main>
        <div className='orders_wrap'>
            {/* <h2>orders</h2> */}
            <p>{props.id}</p>
        </div>
        </main>
    )
    
};

export default Orders;
