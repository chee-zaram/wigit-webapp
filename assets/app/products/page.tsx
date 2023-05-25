// The products page for the wigit web app
"use client";

import useFetch from "@app/hooks/useFetch";
import { ReactElement, JSXElementConstructor, ReactFragment, ReactPortal, PromiseLikeOfReactNode } from "react";

const url: string = "https://jsonplaceholder.typicode.com/todos"

// export const metadata = { title: 'wigit products' };

export default async function Products() {
  // const res: Response = await fetch(url);
  // const data = await res.json();
  const data = useFetch(url);
  return (
    <div>
      <h1>Our wigs</h1>
        <p>Nothing but class....</p>
        { data && data.map((item: any) => {
          <p key={ item.id }>{ item.title }</p>
        }) }
    </div>
  )
}
