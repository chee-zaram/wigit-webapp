import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';
import { verify } from 'jsonwebtoken';

export function middleware(request: NextRequest) {
    //pass
    if (request.nextUrl.pathname.startsWith('/about')) {
        return NextResponse.rewrite(new URL('/signin', request.url));
    }
    return NextResponse.next();
} 

export const config = {
  matcher: '/about/:path*',
};