// sign in / sign up page
"use client";
import { useState } from 'react';

const signin = () => {
    // check if user is signed in
    const [ isSignedIn, setIsSignedIn] = useState(false);
    
    
    return (
        <main className='signin_main'>
            {/* take this to rootlayout to conditionally render sign in link  */}
            <div className='welcome_message'>
                { isSignedIn ? 
                <h3>Welcome Ovy</h3> : 
                <h3>Hey there, we're glad you found us, Please sign in</h3> }
            </div>
        </main>
    )
    
};

export default signin;