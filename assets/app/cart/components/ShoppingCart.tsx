// shopping cart component
"use client";

import { useState, useEffect } from 'react';
import { useSignInContext } from '@app/SignInContextProvider';
import axios from 'axios';
import Item from '@app/cart/interfaces/ShoppingCartProps';
// import { NextPage } from 'next';
import { useRouter } from 'next/navigation';
import Image from 'next/image';

const ShoppingCart: any = async (props: Item) => {
    const [ newQty, setNewQty ] = useState(props.quantity);
    const [ newAmount, setNewAmount ] = useState(Number(props.amount));
    const { jwt } = useSignInContext();
    const headers = {'Authorization': 'Bearer ' + jwt};
    const router = useRouter();

    const handleQtyMinus = async() => {
        if (newQty > 1) {
            setNewQty(newQty - 1);
            setNewAmount( newAmount - (Number(props.amount)/props.quantity));
            //get the price from data
        }
        const qtyUrl = 'https://cheezaram.tech/api/v1/cart/' + props.id + '/' + newQty;
        console.log(qtyUrl);
        const { data, status } = await axios.put(qtyUrl, newQty, {headers: headers});
        console.log(newQty);
    };
    const handleQtyPlus = async() => {
        // check stock
        setNewQty(newQty + 1);
        setNewAmount( newAmount + (Number(props.amount)/props.quantity));
        const qtyUrl = 'https://cheezaram.tech/api/v1/cart/' + props.id + '/' + newQty;
        console.log(qtyUrl);
        const { data, status } = await axios.put(qtyUrl, newQty, {headers: headers});
        console.log(newQty);
    };
    const handleRemoveItem = async() => {
        await axios.delete('https://cheezaram.tech/api/v1/cart/' + props.id, {headers: headers});
        router.push('/');
        // router.push('/cart');
        
    };

    return (
        <main className='md:container'>
            <section className='flex center gap-4 p-4 border border-color-slate-700'>
                <div>
                    <div className=' overflow-hidden h-[100px] w-full'>
                        <Image src={ props.product.image_url } alt={ props.product.name } width={70} height={50}
                        />
                        {/* <img className='max-w-[100%]' src={props.product.image_url} alt={ props.product.name } /> */}
                    </div>
                    <h1>{props.product.name}</h1>
                    {/* <h2>amount = { newAmount }</h2> */}
                    {/* <h2>qty = { newQty }</h2> */}
                </div>

                {/* <button onClick={ handleQtyMinus } className=' bg-slate-400'>-</button>
                <button onClick={ handleQtyPlus } className=' bg-blue-400'>+</button>
                <button onClick={ handleRemoveItem } className=' bg-red-500'>remove</button> */}
            </section>
        </main>
    )
};

export default ShoppingCart;
