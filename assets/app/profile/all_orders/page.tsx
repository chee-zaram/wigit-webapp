//all orders page
"use client";

import axios from 'axios';
import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useSignInContext } from '@app/SignInContextProvider';

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
        router.push('/');
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

    return (
        <section>
            <button onClick={Handleback}>
                <span>Back</span>
            </button>
            <h2>All orders</h2>
            <div>
                { allOrders && allOrders.map((order: any) => (
                    <div key={ order.id }>{ order.user_id }</div>
                ))
                }
            </div>
        </section>
    );
};

export default AllOrders;
