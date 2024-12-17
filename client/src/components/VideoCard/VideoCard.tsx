import React from 'react';
import {useNavigate, NavigateFunction} from 'react-router-dom';
import {Video} from '../../types/Video';
import './video-card.css';

const VideoCard: React.FC<Video> = ({id, title, imagePath}: Video) => {
    const navigate: NavigateFunction = useNavigate();

    const handleClick = async () => {
        navigate(`/video/${id}`);
    };

    return (
        <div className="video-card" onClick={handleClick}>
            <img src={imagePath} alt={title}/>
            <h3>{title}</h3>
        </div>
    );
};

export default VideoCard;