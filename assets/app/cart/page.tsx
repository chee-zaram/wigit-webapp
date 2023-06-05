// shopping cart homepage
"use client";

import { useState, useEffect } from 'react';
import { useSignInContext } from '@app/SignInContextProvider';
import ShoppingCart from '@app/cart/components/ShoppingCart';
import axios from 'axios';
import Button from '@components/Button';
import { useRouter } from 'next/navigation';


const url = 'https://cheezaram.tech/api/v1/cart';
const orderUrl = 'https://cheezaram.tech/api/v1/orders';

const Cart = () => {
    const [ deliveryMethod, setDeliveryMethod ] = useState('');
    const { jwt, role } = useSignInContext();
    const [ cart, setCart ] = useState<any> ([]);
    const headers = {'Authorization': 'Bearer ' + jwt};
    const router = useRouter();

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
    
    const handleSubmit = async(event: any) => {
        event.preventDefault();

        const cartData = { delivery_method: deliveryMethod };
        const { data, status } = await axios.post(orderUrl, cartData, {headers: headers});
        router.push('/');
        alert('order sent');
        // if (status === 'ok') {
        
        // }
    };
    const handleEmptyCart = async() => {
        await axios.delete(url, {headers: headers});
        router.push('/');
        
    };
    
    useEffect(getCart, []);

    return (
        <main>
            <button onClick={handleEmptyCart}>empty cart</button>
            <h2>shopping cart</h2>
            <p>Sha pay and checkout</p>
            <p>{role}</p>
            { cart ?
            <div>
                { cart && cart.map((item) => (
                <div key={item.id}>
                    <ShoppingCart { ...item } />
                </div>
                )) }
            </div> :
            <p>no items in cart</p>
            }
            <form onSubmit={ handleSubmit }>
                <h2> you chose {deliveryMethod}</h2>
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
