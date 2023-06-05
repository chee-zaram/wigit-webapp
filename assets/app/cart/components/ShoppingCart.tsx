// shopping cart component
"use client";

import { useState, useEffect } from 'react';
import { useSignInContext } from '@app/SignInContextProvider';
import axios from 'axios';
import Item from '@app/cart/interfaces/ShoppingCartProps';
import { NextPage } from 'next';

const ShoppingCart: NextPage<Item> = async (props) => {

    return (
        <main className='md:container'>
            <div className='flex center gap-4 p-8 border border-color-slate-700'>
                <h1>product name</h1>
                <h2>amount = { props.amount }</h2>
                <h2>qty = { props.quantity }</h2>

                {/* <p>{ productObj. name }</p> */}
                <button className='border p-4 rounded-full bg-blue-500'>-</button>
                <button className='border p-4 rounded-full bg-blue-500'>+</button>
                <button className='border p-4 rounded-full bg-blue-500'>remove</button>
            </div>
        </main>
    )
};

export default ShoppingCart;
