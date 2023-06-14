// Dashboard
"use client";

import { useSignInContext } from '@app/SignInContextProvider';
import axios from 'axios';
import { useState, useEffect } from 'react';
import Button from '@components/Button';
import Orders from '@app/dashboard/components/Orders';

const Dashboard = async () => {
    const baseUrl = 'https://cheezaram.tech/api/v1/admin';

    const { jwt, setJwt, role, setRole } = useSignInContext();
    
    if (window.sessionStorage.getItem('jwt')) {
            setJwt(window.sessionStorage.getItem('jwt'));
            setRole(window.sessionStorage.getItem('role'));

        }
        
    // if (typeof window !== 'undefined') {
    //     if (window.sessionStorage.getItem('jwt')) {
    //         setJwt(window.sessionStorage.getItem('jwt'));
    //         setRole(window.sessionStorage.getItem('role'));

    //     }
    // };

    const [orders, setOrders] = useState([]);
    const headers = { "Authorization": "Bearer " + jwt};
    
    // fetch a list of orders
    
    useEffect(() => {
    async function getOrders() {
        const { data, status } = await axios.get(baseUrl + '/orders', {headers: headers}) 
        if (status == 200) {
            setOrders(data.data);
            console.log(data);
        }
    };
        getOrders();
    }, []);
    
    return (
        <main className='grid md:grid-rows'>
            <div className='bg-slate-600'><h2>welcome to your dashboard<br/> fucking awesome shit..</h2></div>
            <div>
                { orders && orders.map((order: any) => (
            <div key={ order.id }>
                <Orders { ...order } />
            </div>
            ))}
            </div>
            <div className='bg-pink-600 '>
                <h2>Add new product</h2>
            </div>
            <div className='bg-slate-600 '>
                <h2></h2>
            </div>
            <div className='bg-purple-600 '>
                
            </div>
            <div className='bg-green-600 '></div>
            <div className='bg-blue-600 '></div>
            <div className='bg-yellow-600 '></div>
            <div></div>

        </main>
    )
};

export default Dashboard;