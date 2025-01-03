import React from 'react';
import { useState } from 'react';
import { useParams } from 'react-router-dom';
import axios, { AxiosResponse } from "axios";
import { useVideo } from '../../contexts/VideoProvider';
import { Episode, Movie, Show, Video } from '../../types/Video';
import Plyr from 'plyr-react';
import './video-page.css';
import 'plyr/dist/plyr.css';

const isShow = (video: Video): video is Show => 'episodes' in video;
const isMovie = (video: Video): video is Movie => 'videoPath' in video;

const VideoPage = () => {
    const { videoId } = useParams<{ videoId: string }>();
    const { videos } = useVideo();
    const [videoSource, setVideoSource] = useState<string | null>(null);
    const [videoType, setVideoType] = useState<string | null>(null);
    const [subtitleSource, setSubtitleSource] = useState<string | null>(null);

    const video: Video | undefined = videos.find(v => v.id === videoId);
    if (!video) {
        return <p id="video-not-found">Video not found...</p>;
    }

    const handleClick = async (videoPath: string) => {
        try {
            const videoResponse: AxiosResponse = await axios.get(
                `http://localhost:8080/video?path=${encodeURIComponent(videoPath)}`, {
                responseType: 'blob' // response as a blob (not string or json) to handle video
            });

            const subtitleResponse: AxiosResponse = await axios.get(
                "http://localhost:8080/subtitle", {
                responseType: 'text'
            });

            setVideoSource(URL.createObjectURL(videoResponse.data));
            setVideoType(videoResponse.headers['content-type']);
            setSubtitleSource(URL.createObjectURL(new Blob([subtitleResponse.data], { type: 'text/srt' })));
        } catch (error: any) {
            handleApiError(error);
        }
    };

    return (
        <div id="video-page">
            <h1>{video.title}</h1>
            <div className="video-player">
                <Plyr
                    source={{
                        type: "video",
                        sources: [{
                            src: videoSource || '',
                            type: videoType || '',
                        }],
                        tracks: subtitleSource
                            ? [{
                                kind: "captions",
                                srcLang: "en",
                                label: "English",
                                src: subtitleSource,
                                default: true,
                            }]
                            : [],
                    }}
                    options={{
                        autoplay: false,
                        controls: ['play', 'progress', 'current-time', 'mute', 'volume', 'captions', 'fullscreen'],
                        captions: {
                            active: true,
                            update: true,
                        },
                    }}
                />
            </div>
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