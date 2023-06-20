// Dashboard
"use client";

import 'react-toastify/dist/ReactToastify.css';
import { useRouter } from 'next/navigation';

const Dashboard = async () => {
    const router = useRouter();
    let userObj: string = '';
    if (sessionStorage.getItem('user')) {
        userObj = sessionStorage.getItem('user')!;
    }
    const user: any =  JSON.parse(userObj);
    
    const handleAllOrders = () => {
        router.push('/dashboard/orders');
    };
    const handlePendingOrders = () => {
        router.push('dashboard/pending_orders');
    };
    const handlePaidOrders = () => {
        router.push('dashboard/paid_orders');
    }
    const handleAddProduct = () => {
        router.push('dashboard/add_product');
    };
    
    return (
        <section className='bg-neutral-400 max-w-[80vw] py-6 px-8 rounded-md shadow-md mx-auto'>
            <div>
                <h2 className='p-4 text-2xl font-bold mb-4 text-dark_bg/70'>Welcome to the dashboard, {user.first_name}</h2>
            </div>
            <div className='flexbox_row gap-6 md:gap-8'>
                <div onClick={handleAllOrders} className='bg-neutral-800 text-light_bg text-xl font-medium border p-2 cursor-pointer flex max-w-[150px] flex-col justify-center items-center border-dark_bg hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px]'>
                    <h2>View all orders</h2>
                </div>
                <div onClick={handleAddProduct} className='bg-neutral-800 text-light_bg text-xl font-medium border p-2 cursor-pointer flex max-w-[150px] flex-col justify-center items-center border-dark_bg hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px] '>
                    <h2>Add new product</h2>
                </div>
                <div className='bg-neutral-800 text-light_bg text-xl font-medium border p-2 cursor-pointer flex max-w-[150px] flex-col justify-center items-center border-dark_bg hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px] '>
                    <h2>Manage bookings</h2>
                </div>
                <div className='bg-neutral-800 text-light_bg text-xl font-medium border p-2 cursor-pointer flex max-w-[150px] flex-col justify-center items-center border-dark_bg hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px] '>
                    <h2>Add new slot</h2>
                </div>
                <div className='bg-neutral-800 text-light_bg text-xl font-medium border p-2 cursor-pointer flex max-w-[150px] flex-col justify-center items-center border-dark_bg hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px] '>
                    <h2>settings</h2>
                </div>
                <div onClick={handlePendingOrders} className='bg-yellow-600 bg-neutral-00 text-light_bg text-xl font-medium border p-2 cursor-pointer flex max-w-[150px] flex-col justify-center items-center border-dark_bg hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px] '>
                    <h2>View pending orders</h2>
                </div>
                <div onClick={handlePaidOrders} className='bg-green-700 bg-neutral-00 text-light_bg text-xl font-medium border p-2 cursor-pointer flex max-w-[150px] flex-col justify-center items-center border-dark_bg hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px] '>
                    <h2>View paid orders</h2>
                </div>
            </div>
        </section>
    );
};

export default Dashboard;