import React from 'react'
import { BACK, OPTION } from '../../assets';
import './detailedtask.styles.css';

const Detailedtask = () => {
  return (
    <div className='detailed-task-page'>
        <div className="task-container">
            <div className="task-nav">
                <img src={BACK} alt="" className="back" />
                <img src={OPTION} alt="" className="task-opt" />
            </div>

            <div className="task-tag">
                <p>Office</p>
            </div>
            <h1 className='task-title'>Office Hours Meeting</h1>
            <div className="task-desp-1">
                <div className="task-duration">
                    <p className='task-left'>Time Left</p>
                    <h1>4h 32m</h1>
                    <p>Feb 19, 2022</p>
                </div>
                <div className="task-assignee">

                </div>
            </div>

            <div className="task-desp">
                <p className='task-desp-header'>Additional Description</p>
                <p className='task-desp-content'>
                Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
                </p>
            </div>

            <div className="task-creator">
                <p className='creator-head'>Created</p>
                <p className='created-by'>Feb 19 2022, by Matt</p>
            </div>
        </div>
    </div>
  )
}

export default Detailedtask