import React from "react";
import './button.css';

const STYLE = [
    'btn--primary',
    'btn--outline',
]

export const Button = ({ children, type, onClick, buttonStyle, buttonSize }) => {
    const checkButtonStyle = STYLE.includes(buttonStyle) ? buttonStyle : STYLE[0];
    return (
        <button
            className={`btn ${checkButtonStyle} btn--medium`}
            onClick={onClick}
            type={type}
        >
            {children}
        </button>
    )
}

export default Button;