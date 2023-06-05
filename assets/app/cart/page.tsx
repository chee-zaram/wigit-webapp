// shopping cart homepage
"use client";

import { useState, useEffect } from 'react';
import { useSignInContext } from '@app/SignInContextProvider';
import ShoppingCart from '@app/cart/components/ShoppingCart';
import axios from 'axios';
import Button from '@components/Button';

const url = 'https://cheezaram.tech/api/v1/cart';

const Cart = () => {
    const [ deliveryMethod, setDeliveryMethod ] = useState('');
    const { jwt } = useSignInContext();
    const [ cart, setCart ] = useState<any> ([]);
    const headers = {'Authorization': 'Bearer ' + jwt};
    const getCart = () => {
    fetch(url, {headers: headers})
    .then(res => res.json())
    .then(data => setCart(data.data))
    };
    
    // const handleSetDelivery = (event: React.ChangeEvent<HTMLInputElement>) => {
    //     console.log(deliveryMethod);
    //     event.preventDefault();
    //     setDeliveryMethod(event.target.value);
    //     console.log(deliveryMethod);
    // };
    const handlePickup = () => {
        setDeliveryMethod('pickup');
        console.log(deliveryMethod);
    }
    const handleDelivery = () => {
        setDeliveryMethod('delivery');
        console.log(deliveryMethod);
    }
    
    const handleSubmit = async() => {
        const cartData = {cart: cart, delivery_method: deliveryMethod}
        const { data, status } = await axios.post(url, headers, cartData);
        // if (status === 'ok') {
        
        // }
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
            <form onSubmit={ handleSubmit }>
                <div>
                    <input required onClick={ handlePickup } id='pickup' name='delivery_method' type='radio' value='pickup' />
                    <label htmlFor='pickup'>pickup</label>
                </div>
                <div>
                    <input required onClick={ handleDelivery } id='delivery' name='delivery_method' type='radio' value='delivery' />
                    <label htmlFor='delivery'>delivery</label>
                </div>
                <Button text='checkout' type='submit' />
            </form>
        </main>
    )
};

export default Cart;
