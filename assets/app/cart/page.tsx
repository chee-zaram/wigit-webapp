// shopping cart homepage
"use client";

import { useState, useEffect } from 'react';
import { useSignInContext } from '@app/SignInContextProvider';
import ShoppingCart from '@app/cart/components/ShoppingCart';
import axios from 'axios';
import Button from '@components/Button';
import { useRouter } from 'next/navigation';
import Item from '@app/cart/interfaces/ShoppingCartProps';


const url = 'https://cheezaram.tech/api/v1/cart';
const orderUrl = 'https://cheezaram.tech/api/v1/orders';

const Cart = async () => {
    
    const [ deliveryMethod, setDeliveryMethod ] = useState('');
    const { jwt, setJwt, role } = useSignInContext();
    const [ cart, setCart ] = useState<any> ([]);

if (typeof window !== 'undefined') {
    if (sessionStorage.getItem('jwt')) {
        setJwt(sessionStorage.getItem('jwt'));
    }
}
    
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
        try {
            const { status } = await axios.post(orderUrl, cartData, {headers: headers});
            
            if ( status == 201 ) {
                router.push('/');
                alert('order sent'); 
            }
        }
        catch(error) {
            //catch it here
            alert('something went horribly wrong, and we lost your order. Please shop again.');
        }
    };
    const handleEmptyCart = async() => {
        await axios.delete(url, {headers: headers});
        router.push('/');
        
    };
    
    useEffect(getCart, []);
    // await getCart();
    return (
        <main>
            { jwt !== 'not authorized' ?
        <div>
            <h2 className='text-xxl font-extrabold mb-2'>shopping cart</h2>
            <button onClick={handleEmptyCart}>empty cart</button>
            <p>Sha pay and checkout</p>
            <p>{role}</p>
            { cart && cart.map((item: Item) => (
            <div key={item.id}>
                <ShoppingCart { ...item } />
            </div>
            )) }
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
        </div> :
        <div className='cart_signin mx-auto w-[80vw] h-[40vh]'>
            <p className='bg-light_bg/70 p-8 rounded' >Please <button className='text-accent underline hover:text-accent/60' onClick={ () => router.push('/signin')}>sign in</button> to shop with us</p>
        </div>
        }
    </main>
    )
};

export default Cart;
