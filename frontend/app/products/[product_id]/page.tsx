// products details page
// "use client";
import axios from 'axios';
// import { useState, useEffect } from 'react';
import Image from 'next/image';
import { Product } from '@app/products/interfaces/product';
import { getProducts } from '@app/products/page';

// import { useRouter } from 'next/navigation';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const ProductDetails = async ({params}: {params: {product_id: string} }) => {
    // const router = useRouter();
    const url = 'https://cheezaram.tech/api/v1/products/' + params.product_id;
    let product;
    try {
        product = await getProducts(url);
    } catch (error) {
        toast.error("Unfortunately, you can't view this product now. Please try again.", {
            position: "top-center",
            autoClose: 5000,
            hideProgressBar: true,
            closeOnClick: true,
            pauseOnHover: true,
            draggable: true,
            progress: undefined,
            theme: "colored",
        });
    }
    
    
    return (
        <section>
            { product &&
                <div>
                    <div className='flexbox gap-3 max-w-[80vw] mx-auto'>
                        <h2 className='font-bold text-xl text-dark_bg/80 '>{ product.name }</h2>
                        <div className='border-3 p-3 border-accent bg-dark_bg'>
                            <Image src={product.image_url} alt={product.name} height={400} width={150} />
                        </div>
                        <p>{ product.description }</p>
                        <p className='text-accent font-bold'>{ product.price }</p>
                    </div>
                    <ToastContainer/>
                </div>
            }
        </section>
    );

};

export default ProductDetails;
