// profile pages
"use client"

import { useSignInContext } from '@app/SignInContextProvider';
import { useRouter } from 'next/navigation';

const Profile = () => {
    //pass
    const { jwt, setJwt } = useSignInContext();
    const headers = {'Authorization': 'Bearer ' + jwt};
    
    const user =  JSON.parse(sessionStorage.getItem('user'));
    const router = useRouter();
    
    const handleAllOrders = () => {
        router.push('/profile/all_orders');
    };
    const handlePendingOrders = () => {
        router.push('profile/pending_orders');
    };
    const handleConfirmedOrders = () => {
        router.push('profile/confirmed_orders');
    };
    const handleEditProfile = () => {
        //submit pre-filled form
        router.push('profile/my_profile');
    };

    return (
        <section className='w-[100vw] min-h-screen md:flex'>
            <div className='side_panel bg-dark_bg md:w-1/4'>side panel</div>
            <div className='main_content bg-gray-100 md:w-3/4'>
                <h4 className='p-4 text-lg font-bold text-dark_bg/70 '>Welcome {user.first_name} </h4>
                {/* <p>track an order here... search</p> */}
                <div className='flex gap-4  flex-wrap justify-center'>
                <div onClick={handleAllOrders} className='border p-2 flex max-w-[150px] flex-col justify-center items-center border-accent hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px]'>
                    <h5 className='mb-4 text-sm uppercase font-bold '>All orders</h5>
                    <span className="material-symbols-outlined">order_approve</span>
                </div>
                <div onClick={handlePendingOrders} className='border p-2 flex max-w-[150px] flex-col justify-center items-center border-accent hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px]'>
                    <h5 className='mb-4 text-sm uppercase font-bold '>Pending orders</h5>
                    <span className="material-symbols-outlined">order_approve</span>
                </div>
                <div onClick={handleConfirmedOrders} className='border p-2 max-w-[150px] flex flex-col justify-center items-center border-accent hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px]'>
                    <h5 className='mb-4 text-sm uppercase font-bold '>Confirmed orders</h5>
                    <span className="material-symbols-outlined">order_approve</span>
                </div>
                <div onClick={handleEditProfile} className='border p-2 flex max-w-[150px] flex-col justify-center items-center border-accent hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px]'>
                    <h5 className='mb-4 text-sm uppercase font-bold '>Edit my profile</h5>
                    <span className="material-symbols-outlined">order_approve</span>
                </div>
                </div>
            </div>
        </section>
    )
};

export default Profile;
