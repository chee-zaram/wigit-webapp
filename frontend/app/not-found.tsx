// 404 page
"use client";

import { useRouter } from 'next/navigation';

const NotFound = () => {
  
  const router = useRouter();
  return (
  <div className='not_found min-w-[80vw] min-h-[60vh]'>
    <h3>You seem to have wandered into the woods.</h3>
    <p>No worries, we'll hold your hands back <button onClick={() => router.push('/')} className='underline hover:text-accent'>home</button></p>
  </div>
);
};

export default NotFound;
