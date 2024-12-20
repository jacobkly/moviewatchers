import React from 'react';
import {useParams} from 'react-router-dom';
import axios from 'axios';
import {useVideo} from '../../contexts/VideoProvider';
import {Episode, Show, Movie, Video} from '../../types/Video';
import './video-page.css';

const isShow = (video: Video): video is Show => 'episodes' in video;

const isMovie = (video: Video): video is Movie => 'videoPath' in video;

const VideoPage = () => {
    const {videoId} = useParams<{ videoId: string }>();
    const {videos} = useVideo();

    const video: Video | undefined = videos.find(v => v.id === videoId);
    if (!video) {
        return <p id="video-not-found">Video not found...</p>
    }

    const handleClick = async (videoPath: string) => {
        try {
            await axios.post('http://localhost:8080/play', videoPath, {
                headers: {'Content-Type': 'text/plain'}
            });
        } catch (error: any) {
            handleApiError(error);
        }
    }

    return (
        <div id="video-page">
            <h1>{video.title}</h1>
            {isMovie(video) ? (
                <>
                    <h3>Movie:</h3>
                    <ul>
                        <li onClick={() => handleClick(video.videoPath)}>
                            <p>{video.title}</p>
                        </li>
                    </ul>
                </>
            ) : isShow(video) ? (
                <>
                    <h3>Episodes:</h3>
                    <ul>
                        {video.episodes.map((episode: Episode, index: number) => (
                            <li key={index} onClick={() => handleClick(episode.videoPath)}>
                                <p>{episode.title}</p>
                            </li>
                        ))}
                    </ul>
                </>
            ) : null}
        </div>
    );
};

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

export default VideoPage;