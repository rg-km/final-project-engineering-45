import React from 'react';
import Button from '../../../component/button/button';
import { Link } from 'react-router-dom';
import HeroImage from '../../../component/image/heroimage.png';
import './heroSection.css';

const HeroSection = ({ headLine, description, buttonLabel, alt }) => {
    return (
        <div className="hero-section">
            <div className="hero-section--text-box">
                <h1 className="heading-primary">
                    {headLine}
                </h1>
                <p className="hero-section--text">
                    {description}
                </p>
                <Link to="/RightWay">
                    <Button buttonStyle="btn--primary" buttonSize="btn--medium">
                        {buttonLabel}
                    </Button>
                </Link>
            </div>
            <div className="hero-section--image">
                <img src={HeroImage} alt={alt} className='hero-image' />
            </div>
        </div>
    )
}

export default HeroSection;