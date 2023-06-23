import '@styles/globals.css';
import { Inter } from 'next/font/google';
import Footer from '@components/Footer';
import Navbar from '@components/Navbar';
import { SignInContextProvider } from './SignInContextProvider';

const inter = Inter({ subsets: ['latin'] })

export const metadata = {
  title: 'Wigit Web App',
  description: 'Hair vendor, and wigging service provider',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  
  return (
      <html lang="en">
        <head>
          <link rel="icon" href="/favicon.ico" sizes="any" />
          <script src="https://kit.fontawesome.com/40485fe37e.js" crossOrigin="anonymous"></script>
          <link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200" />
          <link rel="preconnect" href="https://fonts.googleapis.com" />
          <link rel="preconnect" href="https://fonts.gstatic.com" />
          <link href="https://fonts.googleapis.com/css2?family=Playfair+Display:wght@400;500;700;800&family=Poppins:ital,wght@0,300;0,400;0,500;0,600;0,700;0,800;1,500&family=Quicksand:wght@400;500;700&display=swap" rel="stylesheet" />
        </head>
        <SignInContextProvider>
        <body className='bg-white'>
          <Navbar />
          <main>{children}</main>
          <Footer />
        </body>
        </SignInContextProvider>
      </html>
  )
}
