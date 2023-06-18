// products details page
import axios from 'axios';
import BackButton from '@components/BackButton';
import Image from 'next/image';
import { Product } from '@app/products/interfaces/product';
import { getProducts } from '@app/products/page';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
const ProductDetails = async ({params}: {params: {product_id: string} }) => {
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
                    <BackButton />
                    <div className='flexbox gap-3 max-w-[80vw] mx-auto'>
                        <h2 className='font-bold capitalize text-xl text-dark_bg/80 '>{ product.name }</h2>
                        <div className='border-3 p-3 border-accent bg-dark_bg'>
                            <Image src={product.image_url} alt={product.name} height={400} width={250} />
                        </div>
                        <p>{ product.description }</p>
                        <p className='text-accent font-bold'>GHS { product.price }</p>
                    </div>
                    <ToastContainer/>
                </div>
            }
        </section>
    );

};

export default ProductDetails;
