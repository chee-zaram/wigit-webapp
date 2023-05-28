// input interfaces
import { ChangeEventHandler } from "react";

export default interface InputProps {
    name: string;
    type: string;
    placeholder?: string;
    id: string;
    onChange: ChangeEventHandler<HTMLInputElement>;
    required?: boolean;
    autocomplete?: 'on' | 'off';
}