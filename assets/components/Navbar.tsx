// navigation bar
"use client";
import Link from "next/link";
import Logo from '@components/Logo';
import { useState, useEffect } from 'react';


//type check

const Navbar = () => (
  <header className='flex justify-between h-12 bg-neutral-900 text-white items-center'>
    <Logo />
    <nav className="flex px-4 gap-2">
      <Link className='nav_link font-thin' href='/'>Home</Link>
      <Link className='nav_link' href='/signin'>Sign In</Link>
      <Link className='nav_link' href='/products'>Buy Wig</Link>
      <Link className='nav_link' href='/about'>About</Link>
      <Link className='nav_link' href='/cart'>shopping cart</Link>
    </nav>
  </header>
);

export default Navbar;
