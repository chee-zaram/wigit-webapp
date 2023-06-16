// profile pages
"use client"

import { useSignInContext } from '@app/SignInContextProvider';
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';

const Profile = () => {
    const { jwt, setJwt } = useSignInContext();
    if (typeof window !== 'undefined') {
        if (sessionStorage.getItem('jwt')) {
            setJwt(sessionStorage.getItem('jwt'));
        };
    };
    // const headers = {'Authorization': 'Bearer ' + jwt};
    let userObj: string = '';
    if (sessionStorage.getItem('user')) {
        userObj = sessionStorage.getItem('user')!;
    }
    const user: any =  JSON.parse(userObj);
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
    const authorizeUser = () => {
        if (jwt === 'not authorized') {
            router.push('/signin');
        }
    };
    useEffect(authorizeUser, []);

    return (
        <section className='w-[100vw] min-h-screen md:flex'>
            <div className='side_panel bg-dark_bg md:w-1/8 md:mr-4 '>
                <div>
                    side panel
                </div>
            </div>
            <div className='main_content md:w-7/4'>
                <h4 className='p-4 text-lg font-bold text-dark_bg/70 '>Welcome back, {user.first_name}!</h4>
                {/* <p>track an order here... search</p> */}
                <div className='flex gap-4  flex-wrap justify-center'>
                <div onClick={handleAllOrders} className='border p-2 cursor-pointer flex max-w-[150px] flex-col justify-center items-center border-accent hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px]'>
                    <h5 className='mb-4 text-sm uppercase font-bold '>All orders</h5>
                    <span className="material-symbols-outlined">order_approve</span>
                </div>
                <div onClick={handlePendingOrders} className='border p-2 cursor-pointer flex max-w-[150px] flex-col justify-center items-center border-accent hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px]'>
                    <h5 className='mb-4 text-sm uppercase font-bold '>Pending orders</h5>
<svg xmlns="http://www.w3.org/2000/svg" height="40" viewBox="0 -960 960 960" width="40"><path d="M559-97q-18 18-43.5 18T472-97L97-472q-10-10-13.5-21T80-516v-304q0-26 17-43t43-17h304q12 0 24 3.5t22 13.5l373 373q19 19 19 44.5T863-401L559-97ZM245-664q21 0 36.5-15.5T297-716q0-21-15.5-36.5T245-768q-21 0-36.5 15.5T193-716q0 21 15.5 36.5T245-664Z"/></svg>                </div>
                <div onClick={handleConfirmedOrders} className='border p-2 cursor-pointer max-w-[150px] flex flex-col justify-center items-center border-accent hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px]'>
                    <h5 className='mb-4 text-sm uppercase font-bold '>Confirmed orders</h5>
                    <svg xmlns="http://www.w3.org/2000/svg" height="40" viewBox="0 -960 960 960" width="40"><path d="M140-160q-24 0-42-18t-18-42v-520q0-24 18-42t42-18h680q24 0 42 18t18 42v520q0 24-18 42t-42 18H140Zm0-342h680v-129H140v129Z"/></svg>                </div>
                <div onClick={handleEditProfile} className='border p-2 cursor-pointer flex max-w-[150px] flex-col justify-center items-center border-accent hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px]'>
                    <h5 className='mb-4 text-sm uppercase font-bold '>Edit my profile</h5>
                    <svg xmlns="http://www.w3.org/2000/svg" height="40" viewBox="0 -960 960 960" width="40"><path d="M480-481q-66 0-108-42t-42-108q0-66 42-108t108-42q66 0 108 42t42 108q0 66-42 108t-108 42ZM160-160v-94q0-38 19-65t49-41q67-30 128.5-45T480-420q62 0 123 15.5T731-360q31 14 50 41t19 65v94H160Z"/></svg>                </div>
                <div onClick={handleEditProfile} className='border p-2 cursor-pointer flex max-w-[150px] flex-col justify-center items-center border-accent hover:scale-105 duration-300 hover:shadow-accent rounded shadow-md min-h-[150px] min-w-[150px]'>
                    <h5 className='mb-4 text-sm uppercase font-bold '>Settings</h5>
                    <svg xmlns="http://www.w3.org/2000/svg" height="40" viewBox="0 -960 960 960" width="40"><path d="m388-80-20-126q-19-7-40-19t-37-25l-118 54-93-164 108-79q-2-9-2.5-20.5T185-480q0-9 .5-20.5T188-521L80-600l93-164 118 54q16-13 37-25t40-18l20-127h184l20 126q19 7 40.5 18.5T669-710l118-54 93 164-108 77q2 10 2.5 21.5t.5 21.5q0 10-.5 21t-2.5 21l108 78-93 164-118-54q-16 13-36.5 25.5T592-206L572-80H388Zm92-270q54 0 92-38t38-92q0-54-38-92t-92-38q-54 0-92 38t-38 92q0 54 38 92t92 38Zm0-60q-29 0-49.5-20.5T410-480q0-29 20.5-49.5T480-550q29 0 49.5 20.5T550-480q0 29-20.5 49.5T480-410Zm0-70Zm-44 340h88l14-112q33-8 62.5-25t53.5-41l106 46 40-72-94-69q4-17 6.5-33.5T715-480q0-17-2-33.5t-7-33.5l94-69-40-72-106 46q-23-26-52-43.5T538-708l-14-112h-88l-14 112q-34 7-63.5 24T306-642l-106-46-40 72 94 69q-4 17-6.5 33.5T245-480q0 17 2.5 33.5T254-413l-94 69 40 72 106-46q24 24 53.5 41t62.5 25l14 112Z"/></svg>
                </div>
                </div>
            </div>
        </section>
    )
};

export default Profile;
