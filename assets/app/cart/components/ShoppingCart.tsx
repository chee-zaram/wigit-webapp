// shopping cart component
"use client";

import { useState, useEffect } from 'react';
import { useSignInContext } from '@app/SignInContextProvider';
import axios from 'axios';
import Item from '@app/cart/interfaces/ShoppingCartProps';
import { NextPage } from 'next';

const ShoppingCart: NextPage<Item> = async (props) => {
    const [ newQty, setNewQty ] = useState(props.quantity);
    const [ newAmount, setNewAmount ] = useState(Number(props.amount));
    const { jwt } = useSignInContext();
    const headers = {'Authorization': 'Bearer ' + jwt};

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

    return (
        <main className='md:container'>
            <div className='flex center gap-4 p-8 border border-color-slate-700'>
                <h1>product name</h1>
                <h2>amount = { newAmount }</h2>
                <h2>qty = { newQty }</h2>

                {/* <p>{ productObj. name }</p> */}
                <button onClick={ handleQtyMinus } className='border p-4 rounded-full bg-blue-500'>-</button>
                <button onClick={ handleQtyPlus } className='border p-4 rounded-full bg-blue-500'>+</button>
                <button className='border p-4 rounded-full bg-blue-500'>remove</button>
            </div>
        </main>
    )
};

export default ShoppingCart;
