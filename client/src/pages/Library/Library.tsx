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
            if (axios.isAxiosError(error)) {
                console.error("Axios error:", error.message);
            } else {
                console.error("Unexpected error:", error);
            }
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
                            />
                        );
                    })
                )}
            </div>
        </div>
    );
};

export default Library;