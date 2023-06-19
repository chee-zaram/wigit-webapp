// Dashboard
"use client";

// import { useSignInContext } from '@app/SignInContextProvider';
import axios from 'axios';
import { useState, useEffect } from 'react';
import Button from '@components/Button';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import Orders from '@app/dashboard/components/Orders';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer, toast } from 'react-toastify';
import SearchBox from '@components/SearchBox';
import Input from '@components/Input';
import BackButton from '@components/BackButton';
import OrderCard from '@components/OrderCard';


const AdminOrders = async () => {

    const baseUrl = 'https://cheezaram.tech/api/v1/admin';
    const searchUrl = baseUrl + '/orders/';
    const urlObj = {url: searchUrl};
    // const router = useRouter();

    let jwt: string | null = '';
        if (typeof window !== 'undefined') {
            if (sessionStorage.getItem('jwt')) {
                jwt = sessionStorage.getItem('jwt');
            }
    }
    const [orders, setOrders] = useState<any>([]);
    const headers = { "Authorization": "Bearer " + jwt};
    
    // function copy(text:string){
    //   navigator.clipboard.writeText(text);
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
    // }
    
    
    
      
    const [ searchResult, setSearchResult ] = useState<any>(null);
    const [searchInput, setSearchInput ] = useState<string>('');
    
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
        console.log(searchInput);
    };
    
    useEffect(() => {
    async function getOrders() {
        const { data, status } = await axios.get(baseUrl + '/orders', {headers: headers}) 
        if (status == 200) {
            setOrders(data.data);
        }
    };
        getOrders();
    }, []);
    
    return (
        <main className='grid md:grid-rows'>
            <BackButton />
            <SearchBox { ...urlObj} />
            <section>
                <form onSubmit={handleSearch}>
                    <input onChange={(event: React.ChangeEvent<HTMLInputElement>) => {handleSearchInput(event)}}/>
                    <button>Bug</button>
                </form>
                { searchResult &&
                    <div>
                        <p>total: GHS {searchResult.total_amount}</p>
                        <p>With the currency this time, lol... I love you my darling, thank you for staying up with me.</p>
                    </div>
                    // <p>Sorry, we couldn't find a match</p>
                }
            </section>
            <div className='bg-slate-600 my-4 p-4'><h2>All orders </h2></div>
            <div className='w-[80vw] md:w-[70vw] xl:w-[60vw] mx-auto flexbox gap-4'>
                    { orders && orders.map((order: any) => (
                        <OrderCard { ...order } />
                    ))
                    }
            </div>
        </main>
    )
};

export default AdminOrders;