// form component
import Button from '@components/Button';

const signInForm = () => {
    
    const email: string = 'email';
    const password: string = 'password';
    
    return (
        <form>
            <Input type='text' name={email} placeholder='Enter email' id={email}/>
            <Input type='text' name={password} placeholder='Enter password' id={password}/>
            <Button text='sign in' />
            <p>Add sign up / sign in toggle here! :)</p>
        </form>
    )
};

export default signInForm;
