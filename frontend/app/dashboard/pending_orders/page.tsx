// Dashboard pending orders page
"use client";

import axios from 'axios';
import { useState, useEffect } from 'react';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer } from 'react-toastify';
import PendingOrdersCard from '@app/dashboard/components/PendingOrdersCard';
import BackButton from '@components/BackButton';
import PendingSearchBox from '@app/dashboard/components/PendingSearchBox';

const AdminPendingOrders = () => {
    const baseUrl = 'https://backend.wigit.com.ng/api/v1/admin';
    const urlObj = {url: baseUrl + '/orders/', status: 'pending'};

    let jwt: string | null = 'not authorized';
        if (typeof window !== 'undefined') {
            if (sessionStorage.getItem('jwt')) {
                jwt = sessionStorage.getItem('jwt');
            }
    }

    const [orders, setOrders] = useState([]);
    const headers = { "Authorization": "Bearer " + jwt};
    
    useEffect(() => {
    async function getOrders() {
        const { data, status } = await axios.get(baseUrl + '/orders/status/pending', {headers: headers}) 
        if (status == 200) {
            setOrders(data.data);
            console.log(data);
        }
    };
        getOrders();
    }, []);
    
    return (
        <main className='grid md:grid-rows'>
            <BackButton />
            <PendingSearchBox { ...urlObj } />
            <div className='bg-dark_bg/90 my-4 p-4 font-bold capitalize tracking-[3px] text-light_bg/90'><h2>Pending orders </h2></div>
            <div className='min-w-[80vw] md:w-[70vw] mx-auto flexbox md:flex md:flex-row md:gap-6 md:flex-wrap gap-4'>
                    { orders && orders.map((order: any) => (
                        <div key={ order.id }>
                            <PendingOrdersCard { ...order} />
                        </div>
                    ))
                    }
            </div>
            <ToastContainer />
        </main>
    )
};

export default AdminPendingOrders;