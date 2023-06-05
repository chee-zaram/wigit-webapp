// shopping cart homepage
"use client";

import { useState, useEffect } from 'react';
import { useSignInContext } from '@app/SignInContextProvider';
import ShoppingCart from '@app/cart/components/ShoppingCart';
import axios from 'axios';

const url = 'https://cheezaram.tech/api/v1/cart';



const Cart = () => {
    const { jwt } = useSignInContext();
    const [ cart, setCart ] = useState<any> ([]);
    const getCart = () => {
    fetch(url, {headers: {'Authorization': 'Bearer ' + jwt}})
    .then(res => res.json())
    .then(data => setCart(data.data))
    };
    
    useEffect(getCart, []);

    return (
        <main>
            <h2>shopping cart</h2>
            <p>Sha pay and checkout</p>
            <p>{jwt}</p>
            { cart ?
            <div>
                { cart && cart.map((item) => (
                <ShoppingCart { ...item } />
                )) }
            </div> :
            <p>no items in cart</p>
            }
            
        </main>
    )
};

export default Cart;
