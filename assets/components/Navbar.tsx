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
const [ isOpen, setIsOpen ] = useState(false);

const handleMobileNav = () => {
    setIsOpen(!isOpen);
};

// const mobileLink = document.querySelectorAll(".mob_nav_link");
// document.getElementById("open")!.addEventListener('click', function(){
//     document.getElementById("nav_mobile")!.style.width = "80vw";
//     document.getElementById("close")!.style.display = "block";
//     document.getElementById("open")!.style.display = "none";
// })
// document.getElementById("close")!.addEventListener('click', function(){
//     document.getElementById("nav_mobile")!.style.width = "0%";
//     document.getElementById("close")!.style.display = "none";
//     document.getElementById("open")!.style.display = "block";

// })

// mobileLink.forEach(link =>{
//     link.addEventListener('click', () =>{
//         document.getElementById("nav_mobile")!.style.width = "0%";
//         document.getElementById("close")!.style.display = "none";
//         document.getElementById("open")!.style.display = "block";
//     })
// })



  return (
    <header id='header'>
        <Logo />
        <nav id="navbar">
            <a href="#welcome_section_wrap">home</a>
            <a href="#projects">projects</a>
            <a href="#certifications">certifications</a>
            <a href="#connect">connect</a>
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
            </nav>
        </div>
        <div id="open" onClick={handleMobileNav} className={isOpen? 'hide_menu' : 'show_menu'}></div>
    </header>
)};

export default Navbar;
