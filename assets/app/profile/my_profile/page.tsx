// profile page
"use client";

import { useRouter } from 'next/navigation';
import { useSignInContext } from '@app/SignInContextProvider';
import Input from '@components/Input';
import Button from '@components/Button';
import { useState } from 'react';
import axios from 'axios';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer, toast } from 'react-toastify';


const ProfilePage = () => {
    const { jwt, setJwt } = useSignInContext();
    if (typeof window !== 'undefined') {
        if (sessionStorage.getItem('jwt')) {
            setJwt(sessionStorage.getItem('jwt'));
        };
    };
    const headers = {'Authorization': 'Bearer ' + jwt};
    const [ editProfile, setEditProfile ] = useState(false);
    const [ firstName, setFirstName ] = useState('');
    const [ lastName, setLastName ] = useState('');
    const [ address, setAddress ] = useState('');
    const [ phoneNumber, setPhoneNumber ] = useState('');
    const [ isSaving, setIsSaving ] = useState(false);
    const [ password, setPassword ] = useState('');
    const [ confirmPassword, setConfirmPassword ] = useState('');
    const [ passwordInput, setPasswordInput] = useState(false);
    
    let userObj: string = '';
    if (sessionStorage.getItem('user')) {
        userObj = sessionStorage.getItem('user')!;
    }
    const user: any =  JSON.parse(userObj);
    const email = user.email;
    const userData = { email, first_name: firstName, last_name: lastName, phone: phoneNumber, address };
    const url = 'https://cheezaram.tech/api/v1/users/' + email;
    const passwordUrl = 'https://cheezaram.tech/api/v1/reset_password';
    const router = useRouter();
    
    const handleEditProfile = () => {
        setEditProfile(currValue => !currValue);
    }

    // const handleSetPassword = (event: React.ChangeEvent<HTMLInputElement>) => {
    //     event.preventDefault();
    //     setPassword(event.target.value);
    // };
    // const handleSetConfirmPassword = (event: React.ChangeEvent<HTMLInputElement>) => {
    //     event.preventDefault();
    //     setConfirmPassword(event.target.value);
    // };
    const handleSetFirstName = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setFirstName(event.target.value);
    };
    const handleSetLastName = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setLastName(event.target.value);
    };
    const handleSetAddress = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setAddress(event.target.value);
    };
    const handleSetPhoneNumber = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setPhoneNumber(event.target.value);
    };
    const handleSaveEdit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setIsSaving(true);
        try {
            const { status } = await axios.put(url, userData, {headers: headers});
            if (status == 200) {
                toast.success('Personal information updated successfully!', {
                    position: "top-center",
                    autoClose: 3000,
                    hideProgressBar: false,
                    closeOnClick: true,
                    pauseOnHover: true,
                    draggable: true,
                    progress: undefined,
                    theme: "light",
                });
                // user.email = email;
                // user.first_name = firstName;
                // user.last_name = lastName;
                // user.phone = phoneNumber;
                // user.address = address;
                sessionStorage.setItem('user', JSON.stringify(userData));
                console.log(user);
                console.log(userData);
                router.back();
            }
        } catch(error) {
            console.log(error);
            toast.error('Something went wrong, please try again.', {
                position: "top-center",
                autoClose: 4000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                theme: "colored",
            });
        }
        setIsSaving(false);
    };
    const handleChangePassword = async () => {
        setPasswordInput(true);
    };
    const handleSetPassword = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setPassword(event.target.value);
    };
    const handleSetConfirmPassword = (event: React.ChangeEvent<HTMLInputElement>) => {
        event.preventDefault();
        setConfirmPassword(event.target.value);
    };
    const handleSubmitPassword = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        let token: string = '';
        try {
            const { data, status } = await axios.post(passwordUrl, { email }, {headers: headers});
            if (status == 201) {
                token = data.reset_token;
                console.log(data.reset_token);
           }
        } catch(error) {
            console.log(error);
        }
        const passwordData = {email, new_password: password, repeat_new_password: confirmPassword, reset_token: token}
        
        try {
            const { status } = await axios.put(passwordUrl, passwordData, {headers: headers});
            if (status == 200) {
                toast.success('Password changed successfully!', {
                    position: "top-center",
                    autoClose: 3000,
                    hideProgressBar: false,
                    closeOnClick: true,
                    pauseOnHover: true,
                    draggable: true,
                    progress: undefined,
                    theme: "light",
                }); 
                router.back();
            }
        } catch(error) {
            console.log(error);
            toast.error('Something went wrong, please retype the password fields, and try again.', {
                position: "top-center",
                autoClose: 4000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                theme: "colored",
            });
        }
    };
    
    //update session storage with details
    return (
        <section>
            <div onClick={() => {router.back()}} className='hover:bg-accent/80 text-right ml-[10vw] duration-300 rounded-full p-3 max-w-max'>
               <svg xmlns="http://www.w3.org/2000/svg" height="40" viewBox="0 -960 960 960" width="40"><path d="M480-160 160-480l320-320 42 42-248 248h526v60H274l248 248-42 42Z"/></svg> 
            </div>
            <div>
                <h2 className='text-xl font-bold capitalize text-accent mb-4 md:mb-6'>Personal information</h2>
            </div>
            <button onClick={handleEditProfile} className='hover:bg-accent/80 text-right mb-3 duration-300 rounded-full p-3'>
                <svg xmlns="http://www.w3.org/2000/svg" height="30" viewBox="0 -960 960 960" width="30"><path d="M794-666 666-794l42-42q17-17 42.5-16.5T793-835l43 43q17 17 17 42t-17 42l-42 42Zm-42 42L248-120H120v-128l504-504 128 128Z"/></svg>
            </button>
            {!editProfile ?
                <div className='max-w-[80vw] p-4 md:p-10 shadow-md rounded md:max-w-[60vw] mx-auto bg-dark_bg/10'>
                    <div className='profile_data py-2 px-4'>
                        <svg xmlns="http://www.w3.org/2000/svg" height="20" viewBox="0 -960 960 960" width="20"><path d="M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480v53q0 56-39.5 94.5T744-294q-36 0-68-17.5T627-361q-26 34-65 50.5T480-294q-78 0-132.5-54T293-480q0-78 54.5-133T480-668q78 0 132.5 55T667-480v53q0 31 22.5 52t54.5 21q31 0 53.5-21t22.5-52v-53q0-142-99-241t-241-99q-142 0-241 99t-99 241q0 142 99 241t241 99h214v60H480Zm0-274q53 0 90-36.5t37-89.5q0-54-37-91t-90-37q-53 0-90 37t-37 91q0 53 37 89.5t90 36.5Z"/></svg>
                        <p className='ml-4 font-bold text-dark_bg/80 md:text-m'>{ user.email }</p>
                    </div>
                    <div className='profile_data py-2 px-4'>
                        <svg xmlns="http://www.w3.org/2000/svg" height="20" viewBox="0 -960 960 960" width="20"><path d="M480-481q-66 0-108-42t-42-108q0-66 42-108t108-42q66 0 108 42t42 108q0 66-42 108t-108 42ZM160-160v-94q0-38 19-65t49-41q67-30 128.5-45T480-420q62 0 123 15.5T731-360q31 14 50 41t19 65v94H160Z"/></svg>
                        <p className='ml-4 font-bold text-dark_bg/80 md:text-m'>{ user.first_name }</p>
                    </div>
                    <div className='profile_data py-2 px-4'>
                        <svg xmlns="http://www.w3.org/2000/svg" height="20" viewBox="0 -960 960 960" width="20"><path d="M480-481q-66 0-108-42t-42-108q0-66 42-108t108-42q66 0 108 42t42 108q0 66-42 108t-108 42ZM160-160v-94q0-38 19-65t49-41q67-30 128.5-45T480-420q62 0 123 15.5T731-360q31 14 50 41t19 65v94H160Z"/></svg>
                        <p className='ml-4 font-bold text-dark_bg/80 md:text-md'>{ user.last_name }</p>
                    </div>
                    <div className='profile_data py-2 px-4'>
                        <svg xmlns="http://www.w3.org/2000/svg" height="20" viewBox="0 -960 960 960" width="20"><path d="M260-40q-24 0-42-18t-18-42v-760q0-24 18-42t42-18h440q24 0 42 18t18 42v760q0 24-18 42t-42 18H260Zm0-150v90h440v-90H260Zm220.175 75q12.825 0 21.325-8.675 8.5-8.676 8.5-21.5 0-12.825-8.675-21.325-8.676-8.5-21.5-8.5-12.825 0-21.325 8.675-8.5 8.676-8.5 21.5 0 12.825 8.675 21.325 8.676 8.5 21.5 8.5ZM260-250h440v-520H260v520Zm0-580h440v-30H260v30Zm0 640v90-90Zm0-640v-30 30Z"/></svg>
                        <p className='ml-4 font-bold text-dark_bg/80 md:text-md'>{ user.phone }</p>
                    </div>
                    <div className='profile_data py-2 px-4'>
                        <svg xmlns="http://www.w3.org/2000/svg" height="20" viewBox="0 -960 960 960" width="20"><path d="M480-298q103.95-83.86 156.975-161.43Q690-537 690-604q0-59-21.5-100t-53.009-66.88q-31.51-25.881-68.271-37.5Q510.459-820 480-820q-30.459 0-67.22 11.62-36.761 11.619-68.271 37.5Q313-745 291.5-704T270-604q0 67 53.025 144.57T480-298Zm0 76Q343-325 276.5-419.199q-66.5-94.2-66.5-184.554Q210-672 234.5-723.5T298-810q39-35 86.98-52.5 47.98-17.5 95-17.5T575-862.5q48 17.5 87 52.5t63.5 86.533Q750-671.934 750-603.544 750-513 683.5-419 617-325 480-222Zm.089-318Q509-540 529.5-560.589q20.5-20.588 20.5-49.5Q550-639 529.411-659.5q-20.588-20.5-49.5-20.5Q451-680 430.5-659.411q-20.5 20.588-20.5 49.5Q410-581 430.589-560.5q20.588 20.5 49.5 20.5ZM210-80v-60h540v60H210Zm270-524Z"/></svg>
                        <p className='ml-4 font-bold text-dark_bg/80 md:text-md'>{ user.address }</p>
                    </div>
                </div> :
                <form onSubmit={handleSaveEdit} className='max-w-[80vw] p-4 md:p-10 shadow-md rounded md:max-w-[60vw] mx-auto bg-dark_bg/10'>
                    <div className='profile_data py-2 px-4'>
                        <label htmlFor='first_name' className='mr-4 font-bold capitalize text-dark_bg/60 md:text-md'>first name</label>
                        <Input 
                            placeholder='first name'
                            name='first_name'
                            onLoad={(e:any) => handleSetFirstName(e)}
                            onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetFirstName(event)}
                            type='text'
                            id='first_name'
                            required={ false }
                            // value={user.first_name}
                        />
                    </div>
                    <div className='profile_data py-2 px-4'>
                        <label htmlFor='last_name' className='mr-4 font-bold capitalize text-dark_bg/60 md:text-md'>last name</label>
                        <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetLastName(event)}
                            type='text'
                            name='last name'
                            placeholder='Enter last name'
                            id='last_name'
                            autocomplete='on'
                            required={ false }
                            // value={user.last_name}
                        />
                    </div>
                    <div className='profile_data py-2 px-4'>
                        <label htmlFor='phone_number' className='mr-4 font-bold capitalize text-dark_bg/60 md:text-md'>tele phone</label>
                        <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetPhoneNumber(event)}
                            type='tel'
                            name='phone number'
                            placeholder='Enter phone number'
                            id='phone_number'
                            autocomplete='on'
                            required={ false }
                            // value={user.phone}
                        />
                    </div>
                    <div className='profile_data py-2 px-4'>
                        <label htmlFor='address' className='mr-4 font-bold capitalize text-dark_bg/60 md:text-md'>billing add</label>
                        <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetAddress(event)}
                            type='text'
                            name='address'
                            placeholder='Enter address'
                            id='address'
                            autocomplete='on'
                            required={ false }
                            // value={user.address}
                        />
                    </div>
                    {!isSaving?
                        <Button type='submit' text='edit' /> :
                        <Button type='button' disabled={ true } text='saving...' />
                    }
                </form>
            }
            { passwordInput &&
                <form onSubmit={handleSubmitPassword} className='flex flex-col gap-4 my-4 p-2 bg-accent/80 mx-auto rounded max-w-max'>
                <Input
                    onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetPassword(event)}
                    type='password'
                    name='password'
                    placeholder='new password'
                    id='password'
                    autocomplete='off'
                    required={ true }
                />
                <Input onChange={(event: React.ChangeEvent<HTMLInputElement>) => handleSetConfirmPassword(event)}
                    type='password'
                    name='confirm password'
                    placeholder='confirm password'
                    id='confirm_password'
                    autocomplete='off'
                    required={ true }
                />
                    <Button type='submit' text='reset' />
                </form>
            }
            <div onClick={handleChangePassword} className='profile_data cursor-pointer py-2 max-w-max mx-auto px-4'>
                <svg xmlns="http://www.w3.org/2000/svg" height="20" viewBox="0 -960 960 960" width="20"><path d="M280-412q-28 0-48-20t-20-48q0-28 20-48t48-20q28 0 48 20t20 48q0 28-20 48t-48 20Zm0 172q-100 0-170-70T40-480q0-100 70-170t170-70q72 0 126 34t85 103h356l113 113-167 153-88-64-88 64-75-60h-51q-25 60-78.5 98.5T280-240Zm0-60q58 0 107-38.5t63-98.5h114l54 45 88-63 82 62 85-79-51-51H450q-12-56-60-96.5T280-660q-75 0-127.5 52.5T100-480q0 75 52.5 127.5T280-300Z"/></svg>
                <p className='ml-4 font-bold underline capitalize text-dark_bg/60 md:text-md>Change password'>change password</p>
            </div>
            <ToastContainer />
        </section>
    );
};

export default ProfilePage;
