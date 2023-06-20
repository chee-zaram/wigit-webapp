// dashboard paid orders page
"use client";

// import { useSignInContext } from '@app/SignInContextProvider';
import axios from 'axios';
import { useState, useEffect } from 'react';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer, toast } from 'react-toastify';
import SearchBox from '@components/SearchBox';
import BackButton from '@components/BackButton';
import OrderCard from '@components/OrderCard';

const AdminPaidOrders = async () => {

    const baseUrl = 'https://cheezaram.tech/api/v1/admin';
    const searchUrl = baseUrl + '/orders/';
    const urlObj = {url: searchUrl, status: 'paid'};

    let jwt: string | null = '';
        if (typeof window !== 'undefined') {
            if (sessionStorage.getItem('jwt')) {
                jwt = sessionStorage.getItem('jwt');
            }
    }
    const [orders, setOrders] = useState<any>([]);
    const headers = { "Authorization": "Bearer " + jwt};

    useEffect(() => {
    async function getOrders() {
        try {
            const { data, status } = await axios.get(baseUrl + '/orders', {headers: headers}) 
            if (status == 200) {
            setOrders(data.data);
            }
        } catch (error) {
            toast.error('Oops, something went wrong fetching your orders', {
                position: "top-center",
                autoClose: 5000,
                hideProgressBar: true,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                theme: "colored",
            });
        }
    };
        getOrders();
    }, []);
    
    return (
        <main className='grid md:grid-rows'>
            <BackButton />
            <SearchBox { ...urlObj} />
            <div className='bg-dark_bg/90 my-4 p-4 font-bold capitalize tracking-[3px] text-light_bg/90'><h2>Paid orders </h2></div>
            <div className='min-w-[80vw] md:w-[70vw] mx-auto flexbox md:flex md:flex-row md:gap-6 md:flex-wrap gap-4'>
                    { orders ? orders.map((order: any) => (
                        <div key={order.id}>
                            <OrderCard { ...order } />
                        </div>
                    )):
                        <p>There are no paid orders to review at this time.</p>
                    }
            </div>
        </main>
    )
};

export default AdminPaidOrders;