// Product card component 
"use client";

import Button from '@components/Button';
import { Product } from '../interfaces/product';
import { NextPage } from 'next';
import Image from 'next/image';
import axios from 'axios';
import { useSignInContext } from '@app/SignInContextProvider';

const ProductCard: NextPage<Product> = (props) => {
    const { jwt } = useSignInContext();
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
    const { data } = await axios.post(url, payload, {headers: headers});
    console.log(data);
};

return (
    <div className='bg-white m-3 pb-3 hover:scale-105 duration 300 rounded-md shadow-md shadow-zinc-400 border w-70'>
        <div className=' w-70 object-cover bg-blue-200 '>
            <Image src={ props.image_url } alt={ props.name } width={70} height={80}/>
        </div>
        <div className='bg-pink-200 w-[100px]'>
            <h2 className='capitalize text-sm font-bold text-neutral-700 '>{ props.name }</h2>
            <p className='my-1 text-neutral-500 '>{ props.description }</p>
            <p className=' text-accent font-bold'>{ props.price }</p>
            <Button text='add to cart' onClick={() => {addToCart(props.id, 1)}} />
        </div>
    </div>
)};

export default ProductCard;
