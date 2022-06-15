import React, { useState } from "react";
import Button from "../button/button";
import { MenuItems } from "./menuItem";
import { Link } from "react-router-dom";
import './navbar.css';

const Navbar = () => {

    return (
        <nav className='NavbarItems'>
            <Link to='/' className='navbar-logo'>RightWay<i className="fa-solid fa-bacon"></i></Link>
            <ul className='nav-menu'>
                {MenuItems.map((item, index) => (
                    <li key={index}>
                        <Link className={item.cName} to={item.url}>
                            {item.title}
                        </Link>
                    </li>
                ))}
            </ul>
            <Button children='Sign Up' />
        </nav>
    )
}

export default Navbar;