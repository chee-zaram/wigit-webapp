//order card for profile

// Order card
import { NextPage} from 'next';
import Link from 'next/link';
import { ToastContainer, toast } from 'react-toastify';


const ProfileOrderCard: NextPage<any> = (searchResult) => {
    
    function copy(text:string){
      navigator.clipboard.writeText(text);
      toast.info('Reference number copied!', {
        position: "top-center",
        autoClose: 500,
        hideProgressBar: true,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "light",
        });
    }
    
    return (
        <Link href={'/profile/' + searchResult.id} key={ searchResult.id } className='border block border-accent w-max py-3 px-6'>
            <h3>Reference: 
            <span
            className=' px-2 text-accent text-sm underline font-bold'
            onClick={() => copy(searchResult.id.split('-')[0])}>{ searchResult.id.split('-')[0]}</span>
            <span className={searchResult.status === 'pending' || searchResult.status === 'cancelled' ? 'bg-red-500 px-3 py-1 rounded text-light_bg' : 'bg-green-500 px-3 py-1 rounded text-light_bg'}>{ searchResult.status }</span>
            </h3>
            <div>
                <p>Items: <span className='font-bold text-sm'>{ searchResult.items.length }</span></p>
                <p>Total: <span className='font-bold text-sm'>GHS { searchResult.total_amount }</span></p>
                <p>Delivery method: <span className='font-bold text-sm'>{ searchResult.delivery_method }</span></p>
            </div>
        </Link>
    );
};

export default ProfileOrderCard;
