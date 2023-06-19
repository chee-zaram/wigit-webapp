// paid orders
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

const AdminOrders = () => {
    const baseUrl = 'https://cheezaram.tech/api/v1/admin';
    const router = useRouter();

    let jwt: string | null = '';
        if (typeof window !== 'undefined') {
            if (sessionStorage.getItem('jwt')) {
                jwt = sessionStorage.getItem('jwt');
            }
    }

    const [orders, setOrders] = useState([]);
    const headers = { "Authorization": "Bearer " + jwt};
    
    function copy(text:string){
      navigator.clipboard.writeText(text);
      toast.info('Reference number copied!', {
        position: "top-center",
        autoClose: 500,
        hideProgressBar: true,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "light",
        });
    }
    
    useEffect(() => {
    async function getOrders() {
        const { data, status } = await axios.get(baseUrl + '/orders/status/paid', {headers: headers}) 
        if (status == 200) {
            setOrders(data.data);
            console.log(data);
        }
    };
        getOrders();
    }, []);
    
    return (
        <main className='grid md:grid-rows'>
            <div onClick={() => {router.back()}} className='hover:bg-accent/80 text-right ml-[10vw] duration-300 rounded-full p-3 max-w-max'>
               <svg xmlns="http://www.w3.org/2000/svg" height="40" viewBox="0 -960 960 960" width="40"><path d="M480-160 160-480l320-320 42 42-248 248h526v60H274l248 248-42 42Z"/></svg> 
            </div>
            <div className='bg-slate-600 my-4 p-4'><h2>Paid orders </h2></div>
            <div className='w-[80vw] md:w-[70vw] xl:w-[60vw] mx-auto flexbox gap-4'>
                    { orders && orders.map((order: any) => (
                        <Link href={'/dashboard/' + order.id} key={ order.id } className='border border-accent w-full py-3 px-6'>
                            <h3>Reference: 
                            <span
                            className=' px-2 text-accent text-sm underline font-bold'
                            onClick={() => copy(order.id.split('-')[0])}>{ order.id.split('-')[0]}</span>
                            <span className={order.status === 'pending' ? 'bg-red-500 px-3 py-1 rounded text-light_bg' : 'bg-green-500 px-3 py-1 rounded text-light_bg'}>{ order.status }</span>
                            </h3>
                            <div>
                                <p>Items: <span className='font-bold text-sm'>{ order.items.length }</span></p>
                                <p>Total: <span className='font-bold text-sm'>GHS { order.total_amount }</span></p>
                                <p>Delivery method: <span className='font-bold text-sm'>{ order.delivery_method }</span></p>
                            </div>
                        </Link>
                    ))
                    }
            </div>
        </main>
    )
};

export default AdminOrders;