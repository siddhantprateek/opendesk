import React from 'react';
import { BACK, OPTION } from '../../assets';

import './news.styles.css';

const News = () => {
  return (
    <div className='news-page'>
      <div className="news-container">
        <div className="news-nav">
          <img src={BACK} alt="" className="back" />
          <img src={OPTION} alt="" className="task-opt" />
        </div>
        <div className="news-content-box">
          <div className="news-content-ctr">
            <h1>State Management in React with Hookstate</h1>
            <p className='news-by'>@asayer</p>
            <div className="news-topics">
              <p>#react</p>
              <p>#react-hooks</p>
              <p>#axios</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default News;