// navigation bar
"use client";
import Link from "next/link";
import Logo from '@components/Logo';
import { useState, useEffect } from 'react';


//type check

const Navbar = () => (
  <header >
    <Logo />
    <nav className="flex-between w-full pt-3">
      <Link className='nav_link' href='/'>Home</Link>
      <Link className='nav_link' href='/signin'>Sign In</Link>
      <Link className='nav_link' href='/products'>Buy Wig</Link>
      <Link className='nav_link' href='/about'>About</Link>
      <Link className='nav_link' href='/'>shopping cart</Link>
    </nav>
  </header>
);

export default Navbar;
