// add new products page
import axios from 'axios';
import { useState, useEffect } from 'react';
import Button from '@components/Button';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer, toast } from 'react-toastify';
import { Product } from '@app/products/interfaces/product';

const AddProduct = () => {
    const router = useRouter();
    const baseUrl = 'https://cheezaram.tech/api/v1/admin';
    const productData: Product = {
            category,
            description,
            image_url,
            name,
            price,
            stock
        };
        
        let jwt: string | null = '';
        if (typeof window !== 'undefined') {
            if (sessionStorage.getItem('jwt')) {
                jwt = sessionStorage.getItem('jwt');
            }
    }

    const headers = { "Authorization": "Bearer " + jwt};


    const handleAddNewProduct = async () => {
        try {
            const { status } = await axios.post(baseUrl + '/products', productData, {headers: headers});
            if (status == 201) {
                toast.success('Password changed successfully!', {
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
                <h2>Add a new product</h2>
                <div>
                    <form onSubmit={handleAddNewProduct}>
                        < Input />
                        <Button type='submit' text='Create product' />
                    </form>
                </div>
            </section>
        );
};

export default AddProduct;
