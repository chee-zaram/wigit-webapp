// Dashboard
"use client";

// import { useSignInContext } from '@app/SignInContextProvider';
import axios from 'axios';
import { useState, useEffect } from 'react';
import Button from '@components/Button';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import Orders from '@app/dashboard/components/Orders';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer, toast } from 'react-toastify';
import SearchBox from '@components/SearchBox';
import Input from '@components/Input';
import BackButton from '@components/BackButton';
import OrderCard from '@components/OrderCard';


const AdminOrders = async () => {

    const baseUrl = 'https://cheezaram.tech/api/v1/admin';
    const searchUrl = baseUrl + '/orders/';
    const urlObj = {url: searchUrl};
    // const router = useRouter();

    let jwt: string | null = '';
        if (typeof window !== 'undefined') {
            if (sessionStorage.getItem('jwt')) {
                jwt = sessionStorage.getItem('jwt');
            }
    }
    const [orders, setOrders] = useState<any>([]);

    const headers = { "Authorization": "Bearer " + jwt};
      
    const [ searchResult, setSearchResult ] = useState<any>(null);
    const [searchInput, setSearchInput ] = useState<string>('');
    const [hideList, setHideList ] = useState(false);
        
    
    useEffect(() => {
    async function getOrders() {
        const { data, status } = await axios.get(baseUrl + '/orders', {headers: headers}) 
        if (status == 200) {
            setOrders(data.data);
        }
    };
        getOrders();
    }, []);
    
    return (
        <main className='grid md:grid-rows'>
            <BackButton />
            <SearchBox { ...urlObj} />
            <section>
                
            </section>
            <div className='bg-slate-600 my-4 p-4'><h2>All orders </h2></div>
            <div className='w-[80vw] md:w-[70vw] mx-auto flexbox md:flex md:flex-row md:gap-6 md:flex-wrap gap-4'>
                    { orders ? orders.map((order: any) => (
                        <div key={order.id}>
                            <OrderCard { ...order } />
                        </div>
                    )):
                        <p>There are no orders to review at this time.</p>
                    }
            </div>
        </main>
    )
};

export default AdminOrders;