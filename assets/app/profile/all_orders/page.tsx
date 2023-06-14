//all orders page
"use client";

import axios from 'axios';
import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useSignInContext } from '@app/SignInContextProvider';

const getAllOrders = async () => {
        try {
            const {data, status} = await axios.get(url, { headers:headers });
            setAllOrders(data.data);
        } catch(error) {
            console.log(error);
        }
    };

const AllOrders = async() => {
    const [ allOrders, setAllOrders] = useState();
    const { jwt, setJwt } = useSignInContext();
    const router = useRouter();

    if (typeof window !== 'undefined') {
        if (sessionStorage.getItem('jwt')) {
            setJwt(sessionStorage.getItem('jwt'));
        }
    }
    const headers = {'Authorization': 'Bearer ' + jwt};
    const url = 'https://cheezaram.tech/api/v1/orders/status/pending';

    const Handleback = () => {
        router.push('/');
    };

    const all = await getAllOrders();
    console.log(all);

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
