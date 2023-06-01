// shopping cart homepage
"use client";

import { useState, useEffect } from 'react';
import axios from 'axios';

let data: null | [{name: string;}] = null;
const url = 'https://13cecq-8000.csb.app/products';

const Cart = () => {
    const [cart, setCart] = useState<any> ([]);
    useEffect(() => {
    fetch(url)
    .then(res => res.json())
    .then(data => setCart(data))
    }, [])

    return (
        <main>
            <h2>shopping cart</h2>
            <p>Sha pay and checkout</p>
            { cart && cart.map((item: any) => (
                <p>{item.name}</p>
            )) }
        </main>
    )
};

export default Cart;
