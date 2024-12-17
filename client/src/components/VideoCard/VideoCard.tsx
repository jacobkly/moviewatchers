import React from 'react';
import {useNavigate, NavigateFunction} from 'react-router-dom';
import './video-card.css';

interface VideoCardProps {
    id: string;
    title: string;
    imagePath: string;
}

const VideoCard: React.FC<VideoCardProps> = ({id, title, imagePath}: VideoCardProps) => {
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