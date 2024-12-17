import React from 'react';
import {useState, useEffect} from 'react';
import axios, {AxiosResponse} from "axios";
import VideoCard from '../../components/VideoCard/VideoCard';
import '../../assets/styles/app.css';

interface Video {
    title: string;
    imagePath: string;
    videoPath: string;
}

const handleApiError = (error: any): void => {
    if (axios.isAxiosError(error)) {
        if (error.response) {
            switch (error.response.status) {
                case 404:
                    console.error('No JSON library found (404):', error.response.data);
                    break;
                case 500:
                    console.error('Error fetching JSON library (500):', error.response.data);
                    break;
                default:
                    console.error(`Unexpected error (${error.response.status}):`, error.response.data);
            }
        } else if (error.request) {
            console.error('No response received from server:', error.request);
        } else {
            console.error('Error setting up request:', error.message);
        }
    } else {
        console.error('Unexpected error:', error);
    }
};

const Library = () => {
    const [library, setLibrary] = useState<Video[] | null>(null);

    const fetchLibrary = async () => {
        try {
            const response: AxiosResponse = await axios.get('http://localhost:8080/');

            let videoLibrary: Video[] = [];
            for (const key of Object.keys(response.data)) {
                videoLibrary.push({
                    title: key,
                    imagePath: "/assets/images/video-placeholder.png",
                    videoPath: response.data[key]
                });
            }

            setLibrary(videoLibrary);
        } catch (error) {
            handleApiError(error);
        }
    };

    useEffect(() => {
        void fetchLibrary();
    }, []);

    return (
        <div id="library">
            <h1>Library</h1>
            <div id="video-library">
                {library === null ? (
                    <p>Loading videos...</p>
                ) : library.length === 0 ? (
                    <p>No videos available.</p>
                ) : (
                    library.map((video: Video) => {
                        return (
                            <VideoCard
                                key={video.title}
                                title={video.title}
                                imagePath={video.imagePath}
                                videoPath={video.videoPath}
                            />
                        );
                    })
                )}
            </div>
        </div>
    );
};

export default Library;