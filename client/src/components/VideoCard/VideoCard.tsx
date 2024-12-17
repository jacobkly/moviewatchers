import React from 'react';
import '../../assets/styles/app.css';
import axios from "axios";

interface VideoCardProps {
    title: string;
    imagePath: string;
    videoPath: string;
}

const handleApiError = (error: any): void => {
    if (axios.isAxiosError(error)) {
        if (error.response) {
            console.error('Failed to play video:', error.response.data);
        } else if (error.request) {
            console.error('No response received:', error.request);
        } else {
            console.error('Error in request setup:', error.message);
        }
    } else {
        console.error('Unexpected error:', error);
    }
};

const VideoCard: React.FC<VideoCardProps> = ({title, imagePath, videoPath}) => {
    const handleClick = async () => {
        try {
            console.log(videoPath)
            await axios.post('http://localhost:8080/play', videoPath,
                {headers: {'Content-Type': 'text/plain'}});
        } catch (error: any) {
            handleApiError(error);
        }
    };

    return (
        <div className="video-card" onClick={handleClick}>
            <img src={imagePath} alt={title}/>
            <h3>{title}</h3>
        </div>
    );
};

export default VideoCard;