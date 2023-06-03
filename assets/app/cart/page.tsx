// shopping cart homepage
"use client";

import { useState, useEffect } from 'react';
import { useSignInContext } from '../SignInContextProvider';
import axios from 'axios';

let data: null | [{name: string;}] = null;
const url = 'https://cheezaram.tech/api/v1/cart';

const Cart = () => {
    const [cart, setCart] = useState<any> ([]);
    const { jwt} = useSignInContext();

    useEffect(() => {
    fetch(url, {headers: {'Authorization': 'Bearer ' + jwt}})
    .then(res => res.json())
    .then(data => setCart(data.data))
    }, [])

    return (
        <main>
            <h2>shopping cart</h2>
            <p>Sha pay and checkout</p>
            <p>{jwt}</p>
            { cart && cart.map((item: any) => (
                <p>{item.amount}</p>
            )) }
        </main>
    )
};

export default Cart;
