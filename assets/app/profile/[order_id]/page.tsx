// order details page
"use client";
import axios from 'axios';
import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useSignInContext } from '@app/SignInContextProvider';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const OrderDetails = ({ params }: { params: {order_id: string } }) => {
     const router = useRouter();
    const [ order, setOrder ] = useState<any>(null);
    const url = 'https://cheezaram.tech/api/v1/orders/' + params.order_id;

    const HandleBack = () => {
        router.back();
    };

        let jwt: string | null = '';
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
            <button onClick={HandleBack} className='mb-6 ml-[10vw] hover:bg-accent/60 hover:text-light_bg block py-2 px-12 border border-accent rounded shadow text-start font-bold text-accent'>Back</button>
            <h3>Copy ref. number <span
                className='inline cursor-pointer p-2 rounded text-accent text-sm underline font-bold hover:bg-dark_bg/60'
                onClick={() => copy(params.order_id.split('-')[0])}>{ params.order_id.split('-')[0]}
            </span></h3>
            <p className='font-bold text-lg my-2'>Items</p>
            <div>
            {order && order.items.map((item:any) => (
                <div key={ item.id }
                    className='p-4 mb-4 shadow-md hover:border-l-4 hover:border-l-accent mx-auto max-w-[80vw]'
                >
                    <h3>{ item.product.name }</h3>
                    <p>quantity: { item.quantity }</p>
                    <p>amount: { item.amount }</p>
                </div>
            ))}
            </div>
            <p>Order Total: <span className='font-bold text-sm'>{ order && order.total_amount}</span></p>
            <ToastContainer />
            </div>
    )
};

export default OrderDetails;
// export default function Page({ params }: { params: { order_id: string } }) {
//   return <div>My Post: {params.order_id}</div>
// }