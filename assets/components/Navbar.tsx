// navigation bar
"use client";
import Link from "next/link";
import Logo from '@components/Logo';
import { useState } from 'react';
import { useSignInContext } from '@app/SignInContextProvider';

//type check

const Navbar = () => {
    const { jwt, role, setRole, isSignedIn, setIsSignedIn } = useSignInContext();
//if (typeof window !== 'undefined') {
    if (sessionStorage.getItem('role')) {
        setRole(sessionStorage.getItem('role'));
        setIsSignedIn(sessionStorage.getItem('isSignedIn'));
    } else { setIsSignedIn(false)}
//}
const [ isOpen, setIsOpen ] = useState(false);

const handleMobileNav = () => {
    setIsOpen(!isOpen);
};

  return (
    <header id='header' className='mb-10'>
        <Logo />
        <nav id="navbar" className="flex px-4 gap-2 mr-4">
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
        {
          role === 'customer' ?
          <Link className='mob_nav_link' onClick={handleMobileNav} href='/profile'>Profile</Link> :
          <Link className='mob_nav_link' onClick={handleMobileNav} href='/'>blank</Link>
        }
      </nav>
        <div id="nav_mobile" className={isOpen? 'show_nav' : 'hide_nav'}>
            <nav id="nav_mob_cont">
                <a href="javascript:void(0)" id="close" onClick={handleMobileNav}>&times;</a>
                <p>...{ isSignedIn }</p>
                <Link className='mob_nav_link' onClick={handleMobileNav} href='/'>Home</Link>
                {
                  isSignedIn == false ?
                  <Link className='mob_nav_link' onClick={handleMobileNav} href='/signin'>Sign In</Link> :
                  <Link className='mob_nav_link' onClick={handleMobileNav} href='/signout'>Sign Out</Link>
                }
                <Link className='mob_nav_link' onClick={handleMobileNav} href='/products'>Buy Wig</Link>
                <Link className='mob_nav_link' onClick={handleMobileNav} href='/about'>About</Link>
                {
                  role === 'admin' ?
                  <Link className='mob_nav_link' onClick={handleMobileNav} href='/dashboard'>Dashboard</Link> :
                  <Link className='mob_nav_link' onClick={handleMobileNav} href='/cart'>shopping cart</Link>
                }
                {
                  role === 'customer' ?
                  <Link className='mob_nav_link' onClick={handleMobileNav} href='/profile'>Profile</Link> :
                  <Link className='mob_nav_link' onClick={handleMobileNav} href='/'>blank</Link>
                }
            </nav>
        </div>
        <div id="open" onClick={handleMobileNav} className={isOpen? 'hide_menu' : 'show_menu'}></div>
    </header>
)};

export default Navbar;
