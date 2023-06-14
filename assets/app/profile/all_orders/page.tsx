//all orders page
"use client";

import axios from 'axios';
import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { useSignInContext } from '@app/SignInContextProvider';

const AllOrders = () => {
    const router = useRouter();
    const [ allOrders, setAllOrders] = useState();
    const { jwt } = useSignInContext();
    const headers = {'Authorization': 'Bearer ' + jwt};
    const url = 'https://cheezaram.tech/api/v1/orders/status/pending';

    const Handleback = () => {
        router.push('/');
    };
    
    const getAllOrders = async () => {
        try {
            const {data, status} = await axios.get(url, { headers:headers });
            setAllOrders(data.data);
        } catch(error) {
            console.log(error);
        }
    };


    return (
        <section>
            <button onClick={Handleback}>
                <span>Back</span>
            </button>
            <h2>All orders</h2>
        </section>
    );
};

export default AllOrders;
