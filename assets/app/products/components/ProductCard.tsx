// Product card component 

import Button from '@components/Button';
import { Product } from '../interface/product';
import { NextPage } from 'next';
import Image from 'next/image';

const ProductCard: NextPage<Product> = (props) => (
    <div className='m-2 pb-3 flex flex-col rounded shadow-md shadow-zinc-400 border border-zinc-500 max-w-sm'>
        <div className=' w-32 h-32 object-cover '>
            <img src='https://images.pexels.com/photos/6383212/pexels-photo-6383212.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1' />
            {/* <Image src='https://images.pexels.com/photos/6383212/pexels-photo-6383212.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1' alt={ props.name } width={70} height={85}/> */}
        </div>
        <div className='bg-white'>
            <h2>{ props.name }</h2>
            <p>{ props.description }</p>
            <p>{ props.price }</p>
            <Button text='add to cart' />
        </div>
    </div>
);

export default ProductCard;
