// input component
import { NextPage } from 'next';
import InputProps from '@interfaces/InputProps';

const Input: NextPage<InputProps> = ({ onChange, name, type, placeholder, id, required, autocomplete, value } ) => (
    <input
        onChange={ onChange }
        name={ name }
        type={ type }
        placeholder={ placeholder }
        id={ id }
        className='py-2 px-4 outline-none text-sm text-bg-dark_bg rounded-md'
        required={ required }
        autoComplete={ autocomplete }
        value={ value }
    />
);

export default Input;
