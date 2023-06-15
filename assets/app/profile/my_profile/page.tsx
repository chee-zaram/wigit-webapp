// profile page
"use client";

import { useRouter } from 'next/navigation';
import { useSignInContext } from '@app/SignInContextProvider';


const ProfilePage = () => {
    const { jwt, setJwt } = useSignInContext();
    const headers = {'Authorization': 'Bearer ' + jwt};
    
    const user =  JSON.parse(sessionStorage.getItem('user'));
    const router = useRouter();
    
    //update session storage with details
    return (
        <section>
            <div>
                <p>{ user.email }</p>
                <p>{ user.first_name }</p>
                <p>{ user.last_name }</p>
                <p>{ user.address }</p>
                <p>{ user.phone }</p>
            </div>
        </section>
    );
};

export default ProfilePage;
