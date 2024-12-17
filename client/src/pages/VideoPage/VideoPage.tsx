import React from 'react';
import {useParams} from 'react-router-dom';
import {useVideo} from '../../contexts/VideoProvider';

interface Video {
    id: string;
    title: string;
    imagePath: string;
    videoPath: string;
}

const VideoPage = () => {
    const {videoId} = useParams<{ videoId: string }>();
    const {videos} = useVideo();

    const video: Video | undefined = videos.find(v => v.id === videoId);
    if (!video) {
        return <p>Video not found...</p>
    }

    return (
        <div id="video-page">
            <h1>{video.title}</h1>
            <h3>{video.videoPath}</h3>
        </div>
    );
};

export default VideoPage;