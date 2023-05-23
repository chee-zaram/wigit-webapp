// The products page for the wigit web app

const url: string = "https://jsonplaceholder.typicode.com/todos"

export const metadata = { title: 'wigit products' };

export default async function Products() {
  const res: Response = await fetch(url);
  const data = await res.json();
  return (
    <div>
      <h1>Our wigs</h1>
        <p>Nothing but class.... { data[1].title }</p>
    </div>
  )
}
