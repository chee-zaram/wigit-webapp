// Product card component 

import Button from '@components/Button';
import { Product } from '../interfaces/product';
import { NextPage } from 'next';
import Image from 'next/image';

const ProductCard: NextPage<Product> = (props) => (
    <div className='bg-white m-3 pb-3 flex flex-col rounded-md shadow-md shadow-zinc-400 border max-w-sm'>
        <div className=' w-32 h-32 object-cover '>
            <Image src={ props.image_url } alt={ props.name } width={70} height={85}/>
        </div>
        <div>
            <h2 className='uppercase text-sm font-bold text-neutral-700'>{ props.name }</h2>
            <p className='my-1 text-neutral-500'>{ props.description }</p>
            <p className=' text-accent font-bold'>{ props.price }</p>
            <Button text='add to cart' />
        </div>
    </div>
);

export default ProductCard;
