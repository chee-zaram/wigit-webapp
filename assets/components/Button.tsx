// button component 
import { NextPage } from "next";
import ButtonProps from "@interfaces/ButtonProps";

const Button: NextPage<ButtonProps> = ({ text }) => {
    return(
    <button className='py-1 capitalize px-4 text-sm text-slate-900 rounded border border-slate-900 hover:bg-dark_bg hover:text-slate-50 transition-all duration-300'>{ text }</button>
    )
};

export default Button;