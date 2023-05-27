// input component
import { NextPage } from 'next';
import InputProps from '@interfaces/InputProps';

const Input: NextPage<InputProps> = ({onChange, name, type, placeholder, id} ) => (
    <input
        onChange={ onChange }
        name={ name }
        type={ type }
        placeholder={ placeholder }
        id={ id }
        className='py-1 px-2 outline-none border-b border-accent'
    />
);

export default Input;
