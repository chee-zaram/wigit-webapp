// profile pages
"use client"

import { useSignInContext } from '@app/SignInContextProvider';


const Profile = () => {
    //pass
    const { jwt } = useSignInContext();
    const headers = {'Authorization': 'Bearer ' + jwt};
    const user =  JSON.parse(sessionStorage.getItem('user'));

    return (
        <section className='w-[100vw] min-h-screen md:flex'>
            <div className='side_panel bg-dark_bg md:w-1/4'>side panel</div>
            <div className='main_content bg-gray-100 md:w-3/4'>
                <h4 className='text-lg font-bold text-dark_bg/80 '>Welcome {user.first_name} </h4>
                <div className='flex gap-4 flex-wrap justify-center'>
                <div className='border border-accent rounded shadow-md min-h-[150px] min-w-[150px]'>
                    <h5 className='mb-4 text-sm uppercase font-bold text-accent'>My orders</h5>
                    <span className="material-symbols-outlined">order_approve</span>
                </div>
                <div className='bg-slate-300 min-h-[150px] min-w-[150px]'></div>
                <div className='bg-slate-300 min-h-[150px] min-w-[150px]'></div>
                <div className='bg-slate-300 min-h-[150px] min-w-[150px]'></div>
                </div>
            </div>
        </section>
    )
};

export default Profile;
