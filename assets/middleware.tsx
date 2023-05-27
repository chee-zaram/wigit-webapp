// middleware for restricted routes
import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

export function Middleware(request: NextRequest) {
    // pass for now
    if (request.nextUrl.pathname.startsWith('signin')) {
        return NextResponse.rewrite(new URL('/mad-login', request.url));
    }
    return NextResponse.next;
}
export const config = {
  matcher: ['/signin/(.*)', '/profile/*'],
};