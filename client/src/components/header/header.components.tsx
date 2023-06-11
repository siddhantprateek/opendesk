import React from 'react';
import './header.styles.css';

const Header = () => {
  return (
    <div className='header'>
      <div className="left">
        <h2>0pendesk</h2>
      </div>
      <div className="middle">
        <ul className='mid-nav'>
          <li>Solution</li>
          <li>Priciing</li>
          <li>Example</li>
        </ul>
      </div>
      <div className="right">
        <ul className='right-nav'>
          <li>Log in</li>
          <li>Try for free</li>
        </ul>
      </div>
    </div>
  )
}

export default Header