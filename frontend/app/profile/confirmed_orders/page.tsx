//confirmed orders page
"use client";

import axios from 'axios';
import { useState, useEffect } from 'react';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import ProfileOrderCard from '@components/ProfileOrderCard';
import BackButton from '@components/BackButton';
import ProfileSearchBox from '@components/ProfileSearchBox';

const ConfirmedOrders = () => {
    const [ confirmedOrders, setConfirmedOrders ] = useState<string []>([]);
    const url = 'https://cheezaram.tech/api/v1/orders/status/paid';
    const searchUrl = 'https://cheezaram.tech/api/v1/orders/';
    const urlObj = {url: searchUrl, status: 'paid'};
    
    let jwt: string | null = '';
        if (typeof window !== 'undefined') {
            if (sessionStorage.getItem('jwt')) {
                jwt = sessionStorage.getItem('jwt');
            }
    }
    const headers = {'Authorization': 'Bearer ' + jwt};

    useEffect(() => {
    async function getConfirmedOrders() {
        try {
            const { data, status } = await axios.get(url, {headers: headers}) 
            if (status == 200) {
                setConfirmedOrders(data.data);
            }
        } catch(error) {
            console.log(error);
        }
    };
        getConfirmedOrders();
    }, []);

    return (
        <section>
            <BackButton />
            <ProfileSearchBox { ...urlObj} />
            <h2 className='font-bold text-lg text-accent mb-4'>confirmed orders</h2>
            <div className='min-w-[80vw] md:w-[70vw] mx-auto flexbox md:flex md:flex-row md:gap-6 md:flex-wrap gap-4'>
                { confirmedOrders ? confirmedOrders.map((order: any) => (
                    <div key={order.id} className='max-w-max mx-auto'>
                        <ProfileOrderCard { ...order } />
                    </div>
                )) :
            <div className='no_orders_bg'>
                <BackButton />
                <p className='p-4 rounded shadow bg-light_bg/40 max-w-max mx-auto text-md font-bold text-dark_bg'>You currently have no confirmed orders</p>
            </div>
                }
            </div>
            <ToastContainer />
        </section>
    );
};

export default ConfirmedOrders;
