//wavy hairs

//accessories page
import { getProducts } from '@app/products/page';
import { Product } from '@app/products/interfaces/product';
import ProductCard from '@app/products/components/ProductCard';
import { ToastContainer } from 'react-toastify';
import BackButton from '@components/BackButton';

export const metadata = { title: 'Wigit Wavy Wigs' };

const Wavy = async () => {
    const wavyUrl = "https://cheezaram.tech/api/v1/products/categories/wavy";
    const wavyProducts = await getProducts(wavyUrl);
    
    return (
        <section>
            <BackButton />
            <div>
                <h3 className='border-b border-accent text-2xl font-bold text-dark_bg/80 mb-4'>Wavy</h3>
            </div>
            <div className='max-w-[100vw] mx-auto'>
                <div className=' flex justify-center flex-wrap max-w-full gap-4 p-4'>
                {
                wavyProducts && wavyProducts.map((item: Product) => (
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

export default Wavy;
