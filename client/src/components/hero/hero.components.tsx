import React from 'react';
import './hero.styles.css';

const Hero = () => {
  return (
    <div className='hero-section'>
      <div className="hero-container">
        <h1>Productivity at your Service</h1>
      </div>

      <div className="hero-box">
        <div className='hero-content'>
          <p className='hero-content-1'>
            Boost your efficiency with us.
          </p>
          <div className="hero-command-btn">
            <p>
              ctrl+b
            </p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Hero;