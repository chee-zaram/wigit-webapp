// profile page
"use client";

import { useRouter } from 'next/navigation';
import { useSignInContext } from '@app/SignInContextProvider';
import Input from '@components/Input';
import Button from '@components/Button';
import { useState } from 'react';


const ProfilePage = () => {
    const { jwt, setJwt } = useSignInContext();
    const headers = {'Authorization': 'Bearer ' + jwt};
    const [ editProfile, setEditProfile ] = useState(false);
    
    const user =  JSON.parse(sessionStorage.getItem('user'));
    const router = useRouter();
    
    //update session storage with details
    return (
        <section>
            <button>Edit</button>
            {!editProfile ?
                <div>
                    <p>{ user.email }</p>
                    <p>{ user.first_name }</p>
                    <p>{ user.last_name }</p>
                    <p>{ user.address }</p>
                    <p>{ user.phone }</p>
                </div> :
                <form className=' bg-accent'>
                    <Input placeholder='first name' name='first_name' onChange={() => {}} type='text' id='first_name' />
                    <Input name='first_name' onChange={() => {}} type='text' id='first_name' />
                    <Input name='first_name' onChange={() => {}} type='text' id='first_name' />
                    <Input name='first_name' onChange={() => {}} type='text' id='first_name' />
                    <Button type='submit' text='Edit' />
                </form>
            }
        </section>
    );
};

export default ProfilePage;
