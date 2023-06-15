// shopping cart homepage
"use client";

import { useState, useEffect } from 'react';
import { useSignInContext } from '@app/SignInContextProvider';
import ShoppingCart from '@app/cart/components/ShoppingCart';
import axios from 'axios';
import Button from '@components/Button';
import { useRouter } from 'next/navigation';
import Item from '@app/cart/interfaces/ShoppingCartProps';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';


const url = 'https://cheezaram.tech/api/v1/cart';
const orderUrl = 'https://cheezaram.tech/api/v1/orders';
// export const metadata = { title: 'wigit Cart' };

const Cart = () => {
    
    const [ deliveryMethod, setDeliveryMethod ] = useState('');
    const { jwt, setJwt, role } = useSignInContext();
    const [ cart, setCart ] = useState<any> ([]);
    const [total, setTotal ] = useState(0);

if (typeof window !== 'undefined') {
    if (sessionStorage.getItem('jwt')) {
        setJwt(sessionStorage.getItem('jwt'));
    }
}
    
    const headers = {'Authorization': 'Bearer ' + jwt};
    const router = useRouter();
    let sum: number = 0;

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
            alert('failed to fetch cart, please try again');
        }
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
        console.log('cart ooo', cart, cart.length );
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

        const cartData = { delivery_method: deliveryMethod };
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
                router.push('/');
            }
        }
        catch(error) {
            //catch it here
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
    
    useEffect(getCart, []);
    return (
        <main>
            { jwt !== 'not authorized' ?
        // <div className='md:flex justify-between items-center md:container max-w-[80vw]'>
        <div className='md:min-w-5xl md:flex flex-wrap rounded-lg shadow-md max-w-[80vw] lg:max-w-[70vw] mx-auto overflow-hidden'>
            <section className='items_list p-4 md:p-8 lg:p-12 mb-4 md:mb-0 md:w-2/3'>
                <h2 className='text-xxl font-extrabold mb-2'>My shopping cart</h2>
                <button className='hover:bg-red-200 ' onClick={ handleEmptyCart }>
                    <svg xmlns="http://www.w3.org/2000/svg" height="48" viewBox="0 -960 960 960" width="48"><path d="m361-299 119-121 120 121 47-48-119-121 119-121-47-48-120 121-119-121-48 48 120 121-120 121 48 48ZM261-120q-24 0-42-18t-18-42v-570h-41v-60h188v-30h264v30h188v60h-41v570q0 24-18 42t-42 18H261Zm438-630H261v570h438v-570Zm-438 0v570-570Z"/></svg>
                    clear cart
                </button>
                { cart && cart.map((item: Item) => (
                <div key={item.id}>
                    <ShoppingCart { ...item } />
                </div>
                )) }
            </section>
            <section className='bg-neutral-300 md:w-1/3 p-4 md:p-8 lg:p-12'>
                <div className='w-full'>
                    <div>
                        <h4 className='border-b border-slate-200'>Order summary</h4>
                    </div>
                    <form onSubmit={ handleSubmit }>
                        {/* <h2> you chose { deliveryMethod }</h2> */}
                        <div>
                            <input required onClick={ handlePickup } id='pickup' name='delivery_method' type='radio' value='pickup' />
                            <label htmlFor='pickup'>pickup</label>
                        </div>
                        <div>
                            <input required onClick={ handleDelivery } id='delivery' name='delivery_method' type='radio' value='delivery' />
                            <label htmlFor='delivery'>delivery</label>
                        </div>
                        <p>Total {total}</p>
                        <button type='submit' className='rounded bg-dark_bg hover:bg-dark_bg/70 duration-300 text-light_bg w-full px-4 py-2 '>
                            checkout >>
                            {/* <svg xmlns="http://www.w3.org/2000/svg" height="30" viewBox="0 -960 960 960" width="30"><path d="m242-200 210-280-210-280h74l210 280-210 280h-74Zm252 0 210-280-210-280h74l210 280-210 280h-74Z"/></svg> */}
                        </button>
                    </form>
                </div>
            </section>
            <ToastContainer />
        </div> :
        <div className='cart_signin mx-auto w-[80vw] h-[40vh]'>
            <p className='bg-light_bg/70 p-8 rounded' >Please <button className='text-accent underline hover:text-accent/60' onClick={ () => router.push('/signin')}>sign in</button> to shop with us</p>
        </div>
        }
    </main>
    )
};

export default Cart;
