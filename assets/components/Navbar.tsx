// navigation bar
"use client";
import Link from "next/link";
import Logo from '@components/Logo';
import { useState, useEffect } from 'react';
import { useSignInContext } from '@app/SignInContextProvider';

//type check

const Navbar = () => {
    const { jwt, role, setRole, isSignedIn, setIsSignedIn } = useSignInContext();
//if (typeof window !== 'undefined') {
    if (window.sessionStorage.getItem('role')) {
        setRole(window.sessionStorage.getItem('role'));
        setIsSignedIn(window.sessionStorage.getItem('isSignedIn'));
    } else { setIsSignedIn(false)}
//}

  return (
    <header className='flex justify-between h-12 bg-neutral-900 text-white items-center'>
      <Logo />
      <nav className="flex px-4 gap-2">
        <p>...{ isSignedIn }</p>
        <Link className='nav_link font-thin' href='/'>Home</Link>
        {
          isSignedIn == false ?
          <Link className='nav_link' href='/signin'>Sign In</Link> :
          <Link className='nav_link' href='/signout'>Sign Out</Link>
        }
        <Link className='nav_link' href='/products'>Buy Wig</Link>
        <Link className='nav_link' href='/about'>About</Link>
        {
          role === 'admin' ?
          <Link className='nav_link' href='/dashboard'>Dashboard</Link> :
          <Link className='nav_link' href='/cart'>shopping cart</Link>
        }
      </nav>
    </header>
)};

export default Navbar;
