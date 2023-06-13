// profile pages
"use client"

import { useSignInContext } from '@app/SignInContextProvider';


const Profile = () => {
    //pass
    const { jwt } = useSignInContext();
    const headers = {'Authorization': 'Bearer ' + jwt};

    return (
        <section className='w-[100vw] min-h-screen flex'>
            <div className='side_panel bg-dark_bg md:w-1/4'>side panel</div>
            <div className='main_content bg-red-200 md:w-3/4'>profile section</div>
        </section>
    )
};

export default Profile;
