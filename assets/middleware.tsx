// middleware for restricted routes
import { NextResponse } from 'next/server';
import { NextRequest } from 'next/server';

export async function Middleware (request: NextRequest) {
    // pass for now
    if (request.nextUrl.pathname.startsWith('signin')) {
        return NextResponse.rewrite(new URL('mad-loggin', request.url)
    }
    return NextResponse.next;
}
export const config = {
  matcher: ['/signin/(.*)', '/profile/*'],
};