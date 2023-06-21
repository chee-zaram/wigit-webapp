// order details page
"use client";
import axios from 'axios';
import { useState, useEffect } from 'react';
import Image from 'next/image';
import { useRouter } from 'next/navigation';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const OrderDetails = ({ params }: { params: {order_id: string } }) => {
   
    const router = useRouter();
    const [ order, setOrder ] = useState<any>(null);
    const url = 'https://cheezaram.tech/api/v1/orders/' + params.order_id;

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
                setOrder(data.data);
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
        
        <div>
            <div onClick={() => {router.back()}} className='mb-6 hover:bg-accent/80 text-right ml-[10vw] duration-300 rounded-full p-3 max-w-max'>
               <svg xmlns="http://www.w3.org/2000/svg" height="40" viewBox="0 -960 960 960" width="40"><path d="M480-160 160-480l320-320 42 42-248 248h526v60H274l248 248-42 42Z"/></svg> 
            </div>
            <h3>Copy ref. number <span
                className='inline cursor-pointer p-2 rounded text-accent text-sm underline font-bold hover:bg-dark_bg/60'
                onClick={() => copy(params.order_id.split('-')[0])}>{ params.order_id.split('-')[0]}
            </span></h3>
            <p className='font-bold text-lg my-2'>Items</p>
            <div>
            {order && order.items.map((item:any) => (
                <div key={ item.id }
                    className='p-4 flex justify-center gap-12 items-center mb-4 shadow-md hover:border-l-4 hover:border-l-accent mx-auto max-w-[80vw]'
                >
                    <div className='bg-red-300'>
                        <Image src={item.product.image_url} alt={item.product.name} width={40} height={50} />
                    </div>
                    <div className=''>
                        <h3>{ item.product.name }</h3>
                        <p>quantity: { item.quantity }</p>
                        <p>amount: GHS { item.amount }</p>
                    </div>
                </div>
            ))}
            {order && order.delivery_method === 'delivery' ?
                <p>Delivery address: {order.shipping_address}</p> :
                <p>Delivery method: Pickup</p>
            }
            </div>
            <p>Order Total: GHS <span className='font-bold text-sm'>{ order && order.total_amount}</span></p>
            <p className='pt-2 mx-auto border-t border-accent max-w-[150px] font-bold text-lg'>Tracking</p>
            {order &&<div className='pt-1'>Order placed at {order.created_at.split('T')[1].split('Z')[0]} on {order.created_at.split('T')[0]}</div>}
            {order && order.delivered_updated_by !== '' ?
                <div className='mx-auto max-w-[80vw]'>
                    <div className='w-[200px] mx-auto bg-light_bg border shadow rounded-full overflow-hidden h-[20px]'>
                        <div className='w-[100%] bg-dark_bg h-full'></div>
                    </div>
                    <div className='pt-1'>Order delivered on: {order.updated_at.split('T')[0]}</div>
                </div> :
                order && order.shipped_updated_by !== '' ?
                <div className='mt-4 mx-auto max-w-[80vw]'>
                    <div className='mt-4 w-[200px] mx-auto bg-light_bg border shadow rounded-full overflow-hidden h-[20px]'>
                        <div className='w-[70%] bg-dark_bg h-full'></div>
                    </div>
                    <div className='pt-1'>Order shipped on: {order.updated_at.split('T')[0]}</div>
                </div> :
                order && order.paid_updated_by !== '' ?
                <div className='mt-4 mx-auto max-w-[80vw]'>
                    <div className='mt-4 w-[200px] mx-auto bg-light_bg border shadow rounded-full overflow-hidden h-[20px]'>
                        <div className='w-[50%] bg-dark_bg h-full'></div>
                    </div>
                    <div className='pt-1'>Payment confirmed on: {order.updated_at.split('T')[0]}</div>
                </div> :
                order ?
                <div className='mt-4 mx-auto max-w-[80vw]'>
                    <div className='mt-4 w-[200px] mx-auto bg-light_bg border shadow rounded-full overflow-hidden h-[20px]'>
                        <div className='w-[25%] bg-dark_bg h-full'></div>
                    </div>
                    <div className='pt-1'>Payment pending since: {order.created_at.split('T')[0]}</div>
                </div>:
                <div></div>
            }
            <ToastContainer />
            </div>
    )
};

export default OrderDetails;
