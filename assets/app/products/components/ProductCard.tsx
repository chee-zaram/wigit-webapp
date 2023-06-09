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
        <div className='bg-white w-150 hover:scale-105 duration 300'>
            <div className=' w-full w-150 overflow-hidden bg-blue-300 '>
                {/* <Image src={ props.image_url } alt={ props.name } width={150} height={80}
                 className='object-cover'
                /> */}
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
