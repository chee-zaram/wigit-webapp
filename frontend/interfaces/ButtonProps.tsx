// Button props interface

import { MouseEventHandler } from "react";

export default interface ButtonProps {
    type?: "reset" | "submit" | "button";
    text: string;
    onClick?: MouseEventHandler<HTMLButtonElement>;
    disabled?: boolean;
}
