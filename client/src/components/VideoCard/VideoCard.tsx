import React from 'react';
import '../../assets/styles/app.css';

interface VideoCardProps {
    title: string;
    imagePath: string;
}

const VideoCard: React.FC<VideoCardProps> = ({title, imagePath}) => {
    return (
        <div className="video-card">
            <img src={imagePath} alt={title}/>
            <h3>{title}</h3>
        </div>
    );
};

export default VideoCard;