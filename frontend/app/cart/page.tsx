// shopping cart homepage
"use client";

import { useState, useEffect } from 'react';
import ShoppingCart from '@app/cart/components/ShoppingCart';
import axios from 'axios';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import Item from '@app/cart/interfaces/ShoppingCartProps';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';


const url = 'https://backend.wigit.com.ng/api/v1/cart';
const orderUrl = 'https://backend.wigit.com.ng/api/v1/orders';
// export const metadata = { title: 'wigit Cart' };

const Cart = () => {
    
    const [ deliveryMethod, setDeliveryMethod ] = useState('');
    const [ cart, setCart ] = useState<any> ([]);
    const [total, setTotal ] = useState(0);
    const [address, setAddress ] = useState('');

let jwt: string | null = 'not authorized';
        if (typeof window !== 'undefined') {
            if (sessionStorage.getItem('jwt')) {
                jwt = sessionStorage.getItem('jwt');
            }
    }
    
    const headers = {'Authorization': 'Bearer ' + jwt};
    const router = useRouter();
    let sum: number = 0;

    
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
        if (cart.length === 0) {
            toast.error('empty cart!', {
                position: "top-left",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                theme: "light",
            });
            return;
        }

        const cartData = { delivery_method: deliveryMethod, shipping_address: address };
        try {
            const { status } = await axios.post(orderUrl, cartData, {headers: headers});
            
            if ( status == 201 ) {
                 toast.success('Order sent, thank you for shopping. Go to profile to track your orders', {
                    position: "top-center",
                    autoClose: 5000,
                    hideProgressBar: false,
                    closeOnClick: true,
                    pauseOnHover: true,
                    draggable: true,
                    progress: undefined,
                    theme: "light",
                }); 
                router.push('/payment');
            }
        }
        catch(error) {
            toast.error('something went horribly wrong, and we lost your order. Please shop again.', {
                position: "top-center",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                theme: "colored",
            });
        }
    };
    const handleEmptyCart = async() => {
        await axios.delete(url, {headers: headers});
        router.push('/');
    };
    const handleAddress = (event: any) => {
        event.preventDefault();
        setAddress(event.target.value);
    };
    
    useEffect(() => {
        const getCart = () => {
        
        if (jwt === 'not authorized') {
            return;
        }
        try {
            fetch(url, {headers: headers, next: {"revalidate": 0}})
            .then(res => res.json())
            .then(data => setCart(data.data))
            
            cart.forEach((item:Item) => {
                let num = Number(item.amount);
                sum += num;
            })
            setTotal(sum);
            console.log(total, sum);

        } catch(error) {
            console.log('failed to fetch cart, please try again');
        }
    };
    getCart();
    }, []);
    return (
        <main>
            { jwt !== 'not authorized' ?
            <div>
            {cart && cart.length > 0 ?
            <div className='md:min-w-5xl md:flex flex-wrap rounded-lg shadow-md max-w-[80vw] lg:max-w-[70vw] mx-auto overflow-hidden'>
            <section className='items_list p-4 md:p-8 lg:p-12 mb-4 md:mb-0 md:w-2/3'>
                <h2 className='text-2xl capitalize flex items-center justify-around font-extrabold p-3 mb-2'>My shopping cart
                <button className='hover:bg-red-200 p-1' onClick={ handleEmptyCart }>
                    <svg xmlns="http://www.w3.org/2000/svg" height="30" viewBox="0 -960 960 960" width="30"><path d="m361-299 119-121 120 121 47-48-119-121 119-121-47-48-120 121-119-121-48 48 120 121-120 121 48 48ZM261-120q-24 0-42-18t-18-42v-570h-41v-60h188v-30h264v30h188v60h-41v570q0 24-18 42t-42 18H261Zm438-630H261v570h438v-570Zm-438 0v570-570Z"/></svg>
                </button></h2>
                
                <p className='border-b border-dark_bg/80'></p>
                { cart && cart.map((item: Item) => (
                <div key={item.id}>
                    <ShoppingCart { ...item } />
                </div>
                )) }
            </section>
            <section className='bg-neutral-300 md:w-1/3 p-4 md:p-8 lg:p-12'>
                <div className='w-full'>
                    <div>
                        <h4 className='border-b border-light_bg/80 font-bold mb-2 text-lg capitalize'>Order summary</h4>
                    </div>
                    <form onSubmit={ handleSubmit } className='flex flex-col gap-1'>
                         <div >
                            <label htmlFor='address' className='text-xs' >Ship to a different address? (optional)</label><br/>
                            <input onChange={(event: any) => { handleAddress(event) }} id='address' name='address' type='text' placeholder='enter new address' className=' p-1 text-center text-sm w-full outline-none ' />
                        </div>
                    <div className='flex gap-4 justify-center'>
                        <div>
                            <input required onClick={ handlePickup } id='pickup' name='delivery_method' type='radio' value='pickup' />
                            <label htmlFor='pickup'>pickup</label> 
                        </div>
                        <div>
                            <input required onClick={ handleDelivery } id='delivery' name='delivery_method' type='radio' value='delivery' />
                            <label htmlFor='delivery'>delivery</label>
                        </div>
                    </div>
                        <p>Total {total}</p>
                        <button type='submit' className='rounded bg-dark_bg hover:bg-dark_bg/70 duration-300 text-light_bg w-full px-4 py-2 '>
                            checkout &gt;&gt;
                        </button>
                        <Link href={'/products'} className='underline font-bold text-xs p-1 duration-300 hover:text-dark_bg/50'>Continue shopping</Link>
                    </form>
                </div>
            </section>
            <ToastContainer />
            </div> :
            <div className='empty_cart min-h-[70vh] mx-auto min-w-[60vw]'>
                <p className='text-xl max-w-max mx-auto py-8 font-bold text-dark_bg/80'>You have an empty cart, <button className='text-accent underline hover:text-accent/60' onClick={ () => router.push('/products')}>head to shop</button></p>
            </div> 
        }
        </div> :
        <div className='cart_signin mx-auto w-[80vw] h-[40vh]'>
            <p className='bg-light_bg/70 p-8 rounded'>Please <button className='text-accent underline hover:text-accent/60' onClick={ () => router.push('/signin')}>sign in</button> to shop with us</p>
        </div>
        }
    </main>
    )
};

export default Cart;
