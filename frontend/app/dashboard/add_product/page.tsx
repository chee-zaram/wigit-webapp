// add new products page
"use client";

import axios from 'axios';
import { useState } from 'react';
import Input from '@components/Input';
import Button from '@components/Button';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer, toast } from 'react-toastify';
import { Product } from '@app/products/interfaces/product';

const AddProduct = () => {
    const router = useRouter();
    const [name, setName ] = useState('');
    const [description, setDescription ] = useState('');
    const [price, setPrice ] = useState(0);
    const [stock, setStock ] = useState(0);
    const [category, setCategory ] = useState('');
    const [imageUrl, setImageUrl ] = useState('');
    const [id, setId ] = useState('');
    const [isLoading, setIsLoading] = useState(false);

    const baseUrl = 'https://backend.wigit.com.ng/api/v1/admin';
        
        let jwt: string | null = 'not authorized';
        if (typeof window !== 'undefined') {
            if (sessionStorage.getItem('jwt')) {
                jwt = sessionStorage.getItem('jwt');
            }
    }

    const headers = { "Authorization": "Bearer " + jwt};

    const handleSetName = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setName(event.target.value);
    };
    const handleSetDescription = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setDescription(event.target.value);
    };
    const handleSetPrice = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setPrice(Number(event.target.value));
    };
    const handleSetStock = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setStock(Number(event.target.value));
    };
    const handleSetCategory = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setCategory(event.target.value);
    };
    const handleSetImageUrl = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setImageUrl(event.target.value);
    };

    const handleAddNewProduct = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const productData: Product = {
            category,
            description,
            image_url: imageUrl,
            name,
            price,
            stock,
            id
        };
        try {
            const { status } = await axios.post(baseUrl + '/products', productData, {headers: headers});
            if (status == 201) {
                toast.success('Product created', {
                    position: "top-center",
                    autoClose: 3000,
                    hideProgressBar: false,
                    closeOnClick: true,
                    pauseOnHover: true,
                    draggable: true,
                    progress: undefined,
                    theme: "light",
                }); 
            }
        router.back();
        } catch(error) {
            //
        }
    };
        
        return (
            <section>
                <div onClick={() => {router.back()}} className='hover:bg-accent/80 text-right ml-[10vw] duration-300 rounded-full p-3 max-w-max'>
                    <svg xmlns="http://www.w3.org/2000/svg" height="40" viewBox="0 -960 960 960" width="40"><path d="M480-160 160-480l320-320 42 42-248 248h526v60H274l248 248-42 42Z"/></svg> 
                </div>
                <h2 className='text-2xl font-bold text-dark_bg/70 mt-2 mb-6'>Add a new product</h2>
                <div>
                    <form onSubmit={handleAddNewProduct} className='max-w-[90vw] p-4 md:p-10 shadow-md rounded md:max-w-[60vw] mx-auto bg-dark_bg/10'>
                        <div className='profile_data py-2 px-4'>
                            <label htmlFor='name' className='mr-4 text-sm font-bold capitalize text-dark_bg/60 md:text-md'>product name</label>
                            <Input 
                                placeholder='product name'
                                name='name'
                                onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetName(event)}
                                type='text'
                                id='name'
                                required={ true }
                            />
                        </div>
                        <div className='profile_data py-2 px-4'>
                            <label htmlFor='description' className='mr-4 text-sm font-bold capitalize text-dark_bg/60 md:text-md'>description</label>
                            <Input 
                                placeholder='description'
                                name='description'
                                onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetDescription(event)}
                                type='text'
                                id='description'
                                required={ true }
                            />
                        </div>
                        <div className='profile_data py-2 px-4'>
                            <label htmlFor='price' className='mr-4 text-sm font-bold capitalize text-dark_bg/60 md:text-md'>price</label>
                            <Input 
                                placeholder='price'
                                name='price'
                                onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetPrice(event)}
                                type='number'
                                id='price'
                                required={ true }
                            />
                        </div>
                        <div className='profile_data py-2 px-4'>
                            <label htmlFor='stock' className='mr-4 text-sm font-bold capitalize text-dark_bg/60 md:text-md'>stock</label>
                            <Input 
                                placeholder='stock'
                                name='stock'
                                onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetStock(event)}
                                type='number'
                                id='stock'
                                required={ true }
                            />
                        </div>
                        <div className='profile_data py-2 px-4'>
                            <label htmlFor='category' className='mr-4 text-sm font-bold capitalize text-dark_bg/60 md:text-md'>category</label>
                            <Input 
                                placeholder='category'
                                name='category'
                                onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetCategory(event)}
                                type='text'
                                id='category'
                                required={ true }
                            />
                        </div>
                        <div className='profile_data py-2 px-4'>
                            <label htmlFor='image_url' className='mr-4 text-sm font-bold capitalize text-dark_bg/60 md:text-md'>image url</label>
                            <Input 
                                placeholder='image url'
                                name='image_url'
                                onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetImageUrl(event)}
                                type='text'
                                id='image_url'
                                required={ true }
                            />
                        </div>
                        {! isLoading ?
                            <Button type='submit' text='Create product' /> :
                            <Button disabled={ true } type='button' text='Creating' />
                        }
                    </form>
                </div>
                <ToastContainer />
            </section>
        );
};

export default AddProduct;
