// pending_orders card 
import { NextPage } from 'next';
import Link from 'next/link';


const PendingOrdersCard: NextPage<any> = (order)=> {
    return (
      <Link href={'/profile/' + order.id} key={ order.id } className='border border-accent w-full py-3 px-6'>
                        <h3>Reference: 
                        <span
                        className=' px-2 text-accent text-sm underline font-bold'
                        onClick={() => copy(order.id.split('-')[0])}>{ order.id.split('-')[0]}</span>
                        <span className={order.status === 'pending' ? 'bg-red-500 px-3 py-1 rounded text-light_bg' : 'bg-green-500 px-3 py-1 rounded text-light_bg'}>{ order.status }</span>

                        </h3>
                        <div>
                            <p>Items: <span className='font-bold text-sm'>{ order.items.length }</span></p>
                            <p>Total: <span className='font-bold text-sm'>GHS { order.total_amount }</span></p>
                            <p>Delivery method: <span className='font-bold text-sm'>{ order.delivery_method }</span></p>
                        </div>
                    </Link>  
    );
};

export default PendingOrdersCard;
