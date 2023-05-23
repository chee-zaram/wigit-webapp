import '@styles/globals.css';
import { Inter } from 'next/font/google';
import Footer from "@components/Footer";

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
      <body>
        <main className={inter.className}>{children}</main>
        <Footer />
      </body>
    </html>
  )
}
