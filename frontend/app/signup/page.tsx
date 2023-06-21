// Sign up page
import SignUpForm from '@app/signup/components/SignUpForm';

export const metadata = {
  title: 'Wigit Sign up'
}

const SignUp = () => {
    return (
        <main className='signin_main flex flex-col justify-around items-center'>
            <div className='mb-6 capitalize font-extrabold text-dark_bg'>
                <h2>We're glad you found us, please sign up</h2>
            </div>
            <SignUpForm />
        </main>
    )
};

export default SignUp;
