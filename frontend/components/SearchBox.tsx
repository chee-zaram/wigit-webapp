// search box
"use client";
import axios from 'axios';
import { useState } from 'react';
// import { NextPage } from 'next';
import { useSignInContext } from '@app/SignInContextProvider';


const SearchBox = () => {
    const { jwt, setJwt } = useSignInContext();
    if (typeof window !== 'undefined') {
        if (sessionStorage.getItem('jwt')) {
            setJwt(sessionStorage.getItem('jwt'));
        };
    };
    
    const [ searchResult, setSearchResult ] = useState<any>(null);
    const headers = {'Authorization': 'Bearer ' + jwt};
    const [searchInput, setSearchInput ] = useState<string>('');
    const baseUrl = 'https://cheezaram.tech/api/v1/admin';
    const searchUrl = baseUrl + '/orders/';
    
    const handleSearch = async (event: React.FormEvent<HTMLFormElement>) => {
        
        event.preventDefault();
        try {
            const { data, status } = await axios.get(searchUrl + searchInput, { headers: headers });
            if ( status == 200 ) {
                setSearchResult(data.data);
            }
        } catch (error) {
            //
        }
    };
    
    const handleSearchInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setSearchInput(event.target.value);
    };
    
    return (
        <section>
            <form onSubmit={handleSearch}>
                <input onChange={(event: React.ChangeEvent<HTMLInputElement>) => {handleSearchInput(event)}} type='text' placeholder='order reference' />
                <button>Search</button>
            </form>
            { searchResult &&
                <div>
                    <p>{searchResult.id}</p>
                </div>
            }
        </section>
    );
};

export default SearchBox;
