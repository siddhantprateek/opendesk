import React from 'react'
import './motivation.styles.css';
import { BACK, OPTION } from '../../assets';
const Motivation = () => {
  return (
    <div className='motivation-page'>
      <div className="motivation-container">
        <div className="task-nav">
          <img src={BACK} alt="" className="back" />
          <img src={OPTION} alt="" className="task-opt" />
        </div>

        <div className="motivation-content">
          <div className='mvt-container'>
            <h1 className='mvt-content'>
              "Success is often achieved by those who don't know that failure is inevitable."
            </h1>
            <p className='mvt-author'>
            — Coco Chanel, fashion designer 
            </p>

          </div>

        </div>
      </div>
    </div>
  )
}

export default Motivation;