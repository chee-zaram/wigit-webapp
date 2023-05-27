// form component
import Button from '@components/Button';
import Input from '@components/Input';

const signInForm = () => {
    
    const email: string = 'email';
    const password: string = 'password';
    
    return (
        <form className='flex flex-col gap-2 p-4 center max-w-max sm:max-w-l'>
            <Input type='text' name={email} placeholder='Enter email' id={email}/>
            <Input type='text' name={password} placeholder='Enter password' id={password}/>
            <Button text='sign in' />
            <p>Add sign up / sign in toggle here! :)</p>
        </form>
    )
};

export default signInForm;
