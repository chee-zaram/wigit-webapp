// search box
"use client";

import axios from 'axios';
import { useState } from 'react';
import { NextPage } from 'next';
import OrderCard from '@components/OrderCard';
import { ToastContainer, toast } from 'react-toastify';

const SearchBox:NextPage<any> = ( props ) => {
   let jwt: string | null = '';
        if (typeof window !== 'undefined') {
            if (sessionStorage.getItem('jwt')) {
                jwt = sessionStorage.getItem('jwt');
            }
    }

    const [ searchResult, setSearchResult ] = useState<any>(null);
    const headers = {'Authorization': 'Bearer ' + jwt};
    const [searchInput, setSearchInput ] = useState<string>('');

    const handleSearch = async (event: React.FormEvent<HTMLFormElement>) => {

        event.preventDefault();
        try {
            const { data, status } = await axios.get(props.url + searchInput, { headers: headers });
            if ( status == 200 ) {
                setSearchResult(data.data);
            }
        } catch (error) {
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
                <input onChange={(event: React.ChangeEvent<HTMLInputElement>) => {handleSearchInput(event)}} type='text' placeholder='order reference' required className='py-2 px-4 text-xs text-dark_bg/90 border border-accent/70' />
                <button className='py-2 px-6 text-xs border border-accent/70 focus:outline-0 text-light_bg font-bold bg-accent/70 hover:bg-accent/90'>Search</button>
            </form>
            { searchResult && 
            <div>
                { searchResult.status === props.status ?
                <div className='mt-4 max-w-max mx-auto min-h-screen'>
                    <p className='font-bold my-2'>Your result</p>
                    <OrderCard { ...searchResult } />
                </div> :
                <p className='max-w-[80vw]'>No match here, check for the order in your <span className='bg-accent'>{searchResult.status}</span> orders section.</p>
                }
            </div>
            }
            <ToastContainer/>
        </section>
    );
};

export default SearchBox;
