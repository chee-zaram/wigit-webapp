//pending orders page
"use client";

import axios from 'axios';
import { useState, useEffect } from 'react';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import ProfileOrderCard from '@components/ProfileOrderCard';
import BackButton from '@components/BackButton';
import ProfileSearchBox from '@components/ProfileSearchBox';

const PendingOrders = () => {
    const url = 'https://backend.wigit.com.ng/api/v1/orders/status/pending';
    const searchUrl = 'https://backend.wigit.com.ng/api/v1/orders/';
    const urlObj = {url: searchUrl, status: 'pending'};
    const [ pendingOrders, setPendingOrders ] = useState<string []>([]);
        
    let jwt: string | null = 'not authorized';
        if (typeof window !== 'undefined') {
            if (sessionStorage.getItem('jwt')) {
                jwt = sessionStorage.getItem('jwt');
            }
    }
    const headers = {'Authorization': 'Bearer ' + jwt};

    useEffect(() => {
    async function getPendingOrders() {
        try {
            const { data, status } = await axios.get(url, {headers: headers}) 
            if (status == 200) {
                setPendingOrders(data.data);
            }
        } catch(error) {
            console.log(error);
        }
    };
        getPendingOrders();
    }, []);

    return (
        <section>
            { pendingOrders && pendingOrders.length > 0 ?
            <section>
            <BackButton />
            <ProfileSearchBox { ...urlObj} />
            <h2 className='font-bold text-lg text-accent mb-4'>Pending orders</h2>
            <div className='min-w-[80vw] md:w-[70vw] mx-auto flexbox md:flex md:flex-row md:gap-6 md:flex-wrap gap-4'>
                { pendingOrders &&  pendingOrders.map((order: any) => (
                    <div key={order.id} className='max-w-max mx-auto'>
                        <ProfileOrderCard { ...order } />
                    </div>
                ))}
                </div>
                </section> :
            <div className='no_orders_bg'>
                <BackButton />
                <p className='p-4 rounded shadow bg-light_bg/40 max-w-max mx-auto text-md font-bold text-dark_bg'>You currently have no pending orders</p>
            </div>
                }
            <ToastContainer />
        </section>
    );
};

export default PendingOrders;
