// search box
"use client";
import axios from 'axios';
import { useState } from 'react';
import { NextPage } from 'next';
import { useSignInContext } from '@app/SignInContextProvider';
import Link from 'next/link';
import OrderCard from '@components/OrderCard';


const SearchBox:NextPage<any> = (props) => {
    const { jwt, setJwt } = useSignInContext();
    if (typeof window !== 'undefined') {
        if (sessionStorage.getItem('jwt')) {
            setJwt(sessionStorage.getItem('jwt'));
        };
    };
    
    const [ searchResult, setSearchResult ] = useState<any>(null);
    const headers = {'Authorization': 'Bearer ' + jwt};
    const [searchInput, setSearchInput ] = useState<string>('');
    // const baseUrl = 'https://cheezaram.tech/api/v1/admin';
    // const searchUrl = baseUrl + '/orders/';
    
    const handleSearch = async (event: React.FormEvent<HTMLFormElement>) => {
        
        event.preventDefault();
        try {
            const { data, status } = await axios.get(props.url + searchInput, { headers: headers });
            if ( status == 200 ) {
                setSearchResult(data.data);
            }
        } catch (error) {
            alert(error);
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
                <button className='py-2 px-6 text-xs border border-accent/70 focus:outline-0 text-light_bg font-bold bg-accent/70'>Search</button>
            </form>
            { searchResult &&
                <div>
                        <OrderCard { ...searchResult } />
                </div>
            }
        </section>
    );
};

export default SearchBox;
