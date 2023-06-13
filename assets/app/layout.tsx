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
          <script src="https://kit.fontawesome.com/40485fe37e.js" crossOrigin="anonymous"></script>
          <link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200" />
        </head>
        <SignInContextProvider>
        <body className='bg-white'>
          <Navbar />
          <main className={inter.className}>{children}</main>
          <Footer />
        </body>
        </SignInContextProvider>
      </html>
  )
}
