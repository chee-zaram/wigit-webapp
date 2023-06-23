// straight page

//wavy hairs

//accessories page
import { getProducts } from '@app/products/page';
import { Product } from '@app/products/interfaces/product';
import ProductCard from '@app/products/components/ProductCard';
import Link from 'next/link';
import { ToastContainer } from 'react-toastify';
import BackButton from '@components/BackButton';

export const metadata = { title: 'Wigit Wavy Wigs' };

const Straight = async () => {
    const straightUrl = "https://cheezaram.tech/api/v1/products/categories/straight";
    const straightProducts = await getProducts(straightUrl);
    
    return (
        <section>
            <BackButton />
            <div>
                <h3 className='border-b border-accent text-2xl font-bold text-dark_bg/80 mb-4'>straight</h3>
            </div>
            <div className='max-w-[80vw] mx-auto'>
                <div className=' flex justify-center gap-4 p-4'>
                {
                straightProducts && straightProducts.map((item: Product) => (
                  <div key={item.id}>
                    < ProductCard { ...item } />
                  </div>
                ))
            }
            </div>
        </div>
        <ToastContainer />
        </section>
    );
    
};

export default Straight;
