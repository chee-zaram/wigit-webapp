import '@styles/globals.css';
import { Inter } from 'next/font/google';
import Footer from '@components/Footer';
import Navbar from '@components/Navbar';
import ReactQueryrovider from './ReactQueryProvider';

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
    <ReactQueryrovider>
      <html lang="en">
        <body className='bg-zinc-100'>
          <Navbar />
          <main className={inter.className}>{children}</main>
          <Footer />
        </body>
      </html>
    </ReactQueryrovider>
  )
}
