import React from 'react';
import { Hero } from '../../components';
import './home.styles.css';

const Home = () => {
  return (
    <div className='home-page'>
      <div className="home-container">
        <Hero/>
      </div>
    </div>
  )
}

export default Home;