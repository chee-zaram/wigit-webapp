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
