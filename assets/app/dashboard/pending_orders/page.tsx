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

const AdminPendingOrders = async () => {
    const baseUrl = 'https://cheezaram.tech/api/v1/admin';
    const router = useRouter();

    let jwt: string | null = '';
        if (typeof window !== 'undefined') {
            if (sessionStorage.getItem('jwt')) {
                jwt = sessionStorage.getItem('jwt');
            }
    }

    const [orders, setOrders] = useState([]);
    const [showMark, setShowMark] = useState(false);
    const [ paid, setPaid ] = useState(false);
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
    
    const handleShowMarkAsPaid = () => {
      setShowMark(currValue => !currValue);  
    };
    const handleMarkAsPaid = async(id: string) => {
      try {
          const { status } = await axios.put(baseUrl + '/orders/' + id + '/paid', {status: "paid"}, {headers:headers});
          setPaid(true);
          if (status === 200) {
          toast.success('Order marked as paid!', {
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
        router.push('/dashboard/paid_orders');
        setPaid(false);
      } catch (error) {
          alert('something went wrong');
      }
    };
    
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
            <div onClick={() => {router.back()}} className='hover:bg-accent/80 text-right ml-[10vw] duration-300 rounded-full p-3 max-w-max'>
               <svg xmlns="http://www.w3.org/2000/svg" height="40" viewBox="0 -960 960 960" width="40"><path d="M480-160 160-480l320-320 42 42-248 248h526v60H274l248 248-42 42Z"/></svg> 
            </div>
            <div className='bg-slate-600 p-4 my-4'><h2>Pending orders </h2></div>
            <div className='w-[80vw] md:w-[70vw] xl:w-[60vw] mx-auto flexbox gap-4'>
                    { orders && orders.map((order: any) => (
                        <div key={ order.id } className={!paid ? 'border border-accent w-full py-3 px-6' : 'hidden'}>
                            <Link href={'/dashboard/' + order.id} className=' px-3 py-1 rounded mb-4 text-light_bg underline bg-dark_bg/80'><span>view order</span></Link>
                            <h3>Reference: 
                            <span
                            className=' px-2 text-accent text-sm underline font-bold'
                            onClick={() => copy(order.id.split('-')[0])}>{ order.id.split('-')[0]}</span>
                            {!showMark ?
                                <span onClick={handleShowMarkAsPaid} className={order.status === 'pending' ? 'bg-red-500 cursor-pointer px-3 py-1 rounded text-light_bg' : 'bg-green-500 px-3 py-1 rounded text-light_bg'}>{ order.status }</span> :
                                
                                <span>
                                    <button onClick={() => {handleMarkAsPaid(order.id)}} className='bg-green-200 mt-4 duration-300 hover:scale-105 py-2 px-4 rounded shadow-md border font-bold text-green-900 border-green-700'>Mark as paid</button>
                                    <span onClick={handleShowMarkAsPaid}>
                                        <svg xmlns="http://www.w3.org/2000/svg" height="48" viewBox="0 -960 960 960" width="48"><path d="m448-326 112-112 112 112 43-43-113-111 111-111-43-43-110 112-112-112-43 43 113 111-113 111 43 43ZM120-480l169-239q13-18 31-29.5t40-11.5h420q25 0 42.5 17.5T840-700v440q0 25-17.5 42.5T780-200H360q-22 0-40-11.5T289-241L120-480Z"/></svg>
                                    </span>
                                </span>
                            }
                            </h3>
                            <div>
                                <p>Items: <span className='font-bold text-sm'>{ order.items.length }</span></p>
                                <p>Total: <span className='font-bold text-sm'>GHS { order.total_amount }</span></p>
                                <p>Delivery method: <span className='font-bold text-sm'>{ order.delivery_method }</span></p>
                            </div>
                        </div>
                    ))
                    }
            </div>
            <ToastContainer />
        </main>
    )
};

export default AdminPendingOrders;