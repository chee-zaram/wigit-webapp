// Button props interface

import { MouseEventHandler } from "react";

export default interface ButtonProps {
    text: string;
    onClick: MouseEventHandler<HTMLButtonElement>;
}
