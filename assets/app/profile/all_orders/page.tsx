//all orders page
"use client";

import axios from 'axios';
import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useSignInContext } from '@app/SignInContextProvider';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const url = 'https://cheezaram.tech/api/v1/orders';


// const getAllOrders = async () => {
//         const { jwt, setJwt } = useSignInContext();
//         if (typeof window !== 'undefined') {
//             if (sessionStorage.getItem('jwt')) {
//                 setJwt(sessionStorage.getItem('jwt'));
//             }
//     }
//     const headers = {'Authorization': 'Bearer ' + jwt};
//     const url = 'https://cheezaram.tech/api/v1/orders/status/pending';


//         try {
//             const {data, status} = await axios.get(url, { headers:headers });
//             return(data.data);
//         } catch(error) {
//             console.log(error);
//         }
//     };
 
 // fix this
// }

const AllOrders = async() => {
    const router = useRouter();
    const [ allOrders, setAllOrders ] = useState<string []>([]);

    const Handleback = () => {
        router.push('/profile');
    };

// const [allOrders, setAllOrders] = useState<any>([]);
    
    // const { jwt, setJwt } = useSignInContext();
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
                setAllOrders(data.data);
                console.log(data.data);
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
            <button onClick={Handleback}>
                <span>Back</span>
            </button>
            <h2 className='font-bold text-lg text-accent mb-4'>All orders</h2>
            <div className='w-[80vw] md:w-[70vw] xl:w-[60vw] mx-auto flexbox gap-4'>
                { allOrders && allOrders.map((order: any) => (
                    <div key={ order.id } className='border border-accent w-full py-3 px-6'>
                        <h3>Reference: 
                        <span
                        className=' px-2 text-accent text-sm underline font-bold'
                        onClick={() => copy(order.id.split('-')[0])}>{ order.id.split('-')[0]}</span>
                        <span>{ order.status }</span>

                        </h3>
                        <div>
                            <p>Total: <span className='font-bold text-sm'>GHS { order.total_amount }</span></p>
                            <p>Delivery method: <span>{ order.delivery_method }</span></p>
                        </div>
                    </div>
                ))
                }
            </div>
            {/* <span onClick={() => copy('somenewText')}>
            {'somenewText'}
            </span> */}
            <ToastContainer />
        </section>
    );
};

export default AllOrders;
