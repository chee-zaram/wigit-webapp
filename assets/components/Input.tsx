// input component
import { NextPage } from 'next';
import InputProps from '@interfaces/InputProps';

const Input: NextPage<InputProps> = ({ onChange, name, type, placeholder, id, required, autocomplete } ) => (
    <input
        onChange={ onChange }
        name={ name }
        type={ type }
        placeholder={ placeholder }
        id={ id }
        className='py-1 px-2 outline-none rounded-md'
        required={ required }
        autoComplete={ autocomplete }
    />
);

export default Input;
