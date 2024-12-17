import React from 'react';
import {useParams} from 'react-router-dom';
import {useVideo} from '../../contexts/VideoProvider';
import {Video, Show, Movie} from '../../types/Video';
import './video-page.css';

const isShow = (video: Video): video is Show => {
    return (video as Show).episodes !== undefined;
}

const isMovie = (video: Video): video is Movie => {
    return (video as Movie).videoPath !== undefined;
}

const VideoPage = () => {
    const {videoId} = useParams<{ videoId: string }>();
    const {videos} = useVideo();

    const video: Video | undefined = videos.find(v => v.id === videoId);
    if (!video) {
        return <p id="video-not-found">Video not found...</p>
    }

    if (isMovie(video)) {
        console.log("MOVIE")
        return (
            <div id="video-page">
                <h1>{video.title}</h1>
                <h3>Movie:</h3>
                <ul>
                    <li>
                        <p>{video.title}</p>
                    </li>
                </ul>
            </div>
        );
    } else if (isShow(video)) {
        console.log("SHOW")
        return (
            <div id="video-page">
                <h1>{video.title}</h1>
                <h3>Episodes:</h3>
                <ul>
                    {video.episodes.map((episode, index) => (
                        <li key={index}>
                            <p>{episode.title}</p>
                        </li>
                    ))}
                </ul>
            </div>
        );
    }

    return null;
};

export default VideoPage;