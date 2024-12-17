import React from 'react';
import VideoCard from '../components/VideoCard/VideoCard';
import {useVideo} from '../contexts/VideoProvider';
import {Video} from '../types/Video';
import '../assets/styles/app.css';

const Library = () => {
    const {videos} = useVideo();

    return (
        <div id="library">
            <h1>Library</h1>
            <div id="video-library">
                {videos.length === 0 ? (
                    <p>Loading videos...</p>
                ) : (
                    videos.map((video: Video) => (
                        <VideoCard
                            key={video.id}
                            id={video.id}
                            title={video.title}
                            imagePath={video.imagePath}
                        />
                    ))
                )}
            </div>
        </div>
    );
};

export default Library;