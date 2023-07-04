//all orders page
"use client";

import axios from 'axios';
import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useSignInContext } from '@app/SignInContextProvider';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import Link from 'next/link';

const url = 'https://backend.wigit.com.ng/api/v1/orders';

const AllOrders = async() => {
    const router = useRouter();
    const [ allOrders, setAllOrders ] = useState<string []>([]);
    
    let jwt: string | null = 'not authorized';
        if (typeof window !== 'undefined') {
            if (sessionStorage.getItem('jwt')) {
                jwt = sessionStorage.getItem('jwt');
            }
    }
    const headers = {'Authorization': 'Bearer ' + jwt};

    useEffect(() => {
    async function getAllOrders() {
        try {
            const { data, status } = await axios.get(url, {headers: headers}) 
            if (status == 200) {
                setAllOrders(data.data);
            }
        } catch(error) {
            console.log(error);
        }
    };
        getAllOrders();
    }, []);

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

    return (
        <section>
        {allOrders.length > 0 ?
            <section>
                <div onClick={() => {router.back()}} className='mb-6 hover:bg-accent/80 text-right ml-[10vw] duration-300 rounded-full p-3 max-w-max'>
                   <svg xmlns="http://www.w3.org/2000/svg" height="40" viewBox="0 -960 960 960" width="40"><path d="M480-160 160-480l320-320 42 42-248 248h526v60H274l248 248-42 42Z"/></svg> 
                </div>
                <h2 className='font-bold text-lg text-accent mb-4'>All orders</h2>
                <div className='w-[80vw] md:w-[70vw] xl:w-[60vw] mx-auto flexbox gap-4'>
                    { allOrders && allOrders.map((order: any) => (
                        <Link href={'/profile/' + order.id} key={ order.id } className='border border-accent w-full py-3 px-6'>
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
                <ToastContainer />
            </section> :
            <div className='no_orders_bg'>
            <div onClick={() => {router.back()}} className='mb-6 hover:bg-accent/80 text-right ml-[10vw] duration-300 rounded-full p-3 max-w-max'>
               <svg xmlns="http://www.w3.org/2000/svg" height="40" viewBox="0 -960 960 960" width="40"><path d="M480-160 160-480l320-320 42 42-248 248h526v60H274l248 248-42 42Z"/></svg> 
            </div>
            <p className='p-4 rounded shadow bg-light_bg/40 max-w-max mx-auto text-md font-bold text-dark_bg'><span onClick={() => {router.push('/signin')}} className='cursor-pointer underline text-accent'>Shop now</span> to see your orders here </p>
        </div>
        }
        </section>
    );
};

export default AllOrders;
