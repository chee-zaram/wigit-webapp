// search box
"use client";
import axios from 'axios';
import { useState } from 'react';
import { NextPage } from 'next';
import { useSignInContext } from '@app/SignInContextProvider';
import Link from 'next/link';


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
            //
        }
    };
    
    function copy(text:string){
      navigator.clipboard.writeText(text);
    //   toast.info('Reference number copied!', {
    //     position: "top-center",
    //     autoClose: 500,
    //     hideProgressBar: true,
    //     closeOnClick: true,
    //     pauseOnHover: true,
    //     draggable: true,
    //     progress: undefined,
    //     theme: "light",
    //     });
    }
    
    
    const handleSearchInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setSearchInput(event.target.value);
    };
    
    return (
        <section>
            <form onSubmit={handleSearch} className='p-2'>
                <input onChange={(event: React.ChangeEvent<HTMLInputElement>) => {handleSearchInput(event)}} type='text' placeholder='order reference' className='py-2 px-4 text-xs text-dark_bg/90 outline-accent' />
                <button className='py-2 px-6 text-xs bg-accent/70'>Search</button>
            </form>
            { searchResult &&
                <div>
                        <p>total: GHS {searchResult.total_amount}</p>
                        <p>With the currency this time, lol... I love you my darling, thank you for staying up with me.</p>
                        <p>This one doesn't break baby, praise be</p>
                        <Link href={'/dashboard/' + searchResult.id} key={ searchResult.id } className='bsearchResult bsearchResult-accent w-full py-3 px-6'>
                            <h3>Reference: 
                            <span
                            className=' px-2 text-accent text-sm underline font-bold'
                            onClick={() => copy(searchResult.id.split('-')[0])}>{ searchResult.id.split('-')[0]}</span>
                            <span className={searchResult.status === 'pending' ? 'bg-red-500 px-3 py-1 rounded text-light_bg' : 'bg-green-500 px-3 py-1 rounded text-light_bg'}>{ searchResult.status }</span>
                            </h3>
                            <div>
                                <p>Items: <span className='font-bold text-sm'>{ searchResult.items.length }</span></p>
                                <p>Total: <span className='font-bold text-sm'>GHS { searchResult.total_amount }</span></p>
                                <p>Delivery method: <span className='font-bold text-sm'>{ searchResult.delivery_method }</span></p>
                            </div>
                        </Link>
                </div>
            }
        </section>
    );
};

export default SearchBox;
