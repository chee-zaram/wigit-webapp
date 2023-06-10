// Product card component 
"use client";

import Button from '@components/Button';
import { Product } from '../interfaces/product';
import { NextPage } from 'next';
import { useState } from 'react';
import Image from 'next/image';
import axios from 'axios';
import { useSignInContext } from '@app/SignInContextProvider';

const ProductCard: NextPage<Product> = (props) => {
    const { jwt } = useSignInContext();
    const [ signInAlert, setSignInAlert ] = useState(false);
    const [ somethingWrong, setSomethingWrong ] = useState(false);
    
    const addToCart = async (product_id: string, quantity: number) => {
    // send the product id and quantity
    const url = "https://cheezaram.tech/api/v1/cart";
    const payload = {
        product_id,
        quantity
    }
    const headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + jwt
    }
    if (jwt === 'not authorized') {
        setSignInAlert(true);
        return;
    }
    const { status } = await axios.post(url, payload, {headers: headers});
    if (status != 200) {
        setSomethingWrong(true);
    }
};

return (
    <section>
        { !signInAlert ?
        <div className='bg-white shrink-0 shadow-lg overflow-hidden rounded w-[200px] min-h-[250px] flexbox_row hover:scale-105 duration 400'>
        <div className=' overflow-hidden h-[120px] w-full'>
            <Image src={ props.image_url } alt={ props.name } width={200} height={100}
             className='object-cover'
            />
        </div>
        <div className=' p-4 w-full'>
            <h2 className='capitalize text-sm font-bold text-neutral-700 '>{ props.name }</h2>
            <p className='my-1 text-neutral-500 '>{ props.description }</p>
            <p className={props.category === 'straight' ? 'bg-sky-700 tag' : props.category === 'wavy' ? 'bg-pink-700 tag' : 'bg-teal-600 tag'}>{ props.category}</p>
            <p className=' text-accent font-bold'>GHS { props.price }</p>
            <Button text='add to cart' onClick={() => {addToCart(props.id, 1)}} />
        </div>
    </div> :
    <div className='w-screen h-screen bg-slate-700 text-white absolute top-50 left-0'>please sign in first.. add link to sign in page...<br/> to be sorted out later ;) handle multiple entries too </div>
    }
    </section>
    
)};

export default ProductCard;
