// search box for products
"use client";

import axios from 'axios';
import { useState } from 'react';
import { NextPage } from 'next';
import { getProducts } from '@app/products/page';
import { Product } from '@app/products/interfaces/product';
import ProductCard from '@app/products/components/ProductCard';
import { ToastContainer, toast } from 'react-toastify';


const ProductSearchBox:NextPage<any> = ( props ) => {

    const [ searchResult, setSearchResult ] = useState<any>(null);
    const [ filterResult, setFilterResult ] = useState<any>(null);
    const [ hideResult, setHideResult ] = useState(false);
    const [searchInput, setSearchInput ] = useState<string>('');

    const handleSearch = async (event: React.FormEvent<HTMLFormElement>) => {

        event.preventDefault();
        try {
            const { data, status } = await axios.get(props.url);
            if ( status == 200 ) {
                setSearchResult(data.data);
                   
                setHideResult(false);
            }
        } catch (error) {
            setHideResult(true);
            toast.info("We didn't find any results for your search", {
                position: "top-center",
                autoClose: 5000,
                hideProgressBar: true,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                theme: "light",
            });
        }
    };

    const handleSearchInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setSearchInput(event.target.value);
    };

    return (
        <section>
            <form onSubmit={handleSearch} className='p-2 '>
                <input onChange={(event: React.ChangeEvent<HTMLInputElement>) => {handleSearchInput(event)}} type='text' placeholder='products/tags' required className='py-2 px-4 text-xs text-dark_bg/90 border border-accent/70' />
                <button className='py-2 px-6 text-xs border border-accent/70 focus:outline-0 text-light_bg font-bold bg-accent/70 hover:bg-accent/90'>Search</button>
            </form>
            <div>
            { filterResult ?
            <div className={hideResult ? 'hidden': 'block'}>
                <div className='mt-4 max-w-max mx-auto '>
                    <p className='font-bold my-2'>Your results</p>
                    <div>
                        {/* {searchResult && searchResult.map ( (item) => (
                            <div key={item.id}>
                                <ProductCard { ...item } />   
                            </div> ))} */}
                    </div>
                </div>
                </div>:
                <p></p>
                // <p className='max-w-[80vw] mx-auto'>We regret that you didn't find a match.</p>
            }
            </div>
            <ToastContainer/>
        </section>
    );
};

export default ProductSearchBox;
