import React from 'react';
import {createContext, useContext, useState, useEffect, ReactNode} from 'react';
import axios, {AxiosResponse} from "axios";
import {sha256} from 'js-sha256';

interface Video {
    id: string;
    title: string;
    imagePath: string;
    videoPath: string;
}

interface VideoContextType {
    videos: Video[];
    setVideos: (videos: Video[]) => void;
}

const VideoContext = createContext<VideoContextType | undefined>(undefined);

interface VideoProviderProps {
    children: ReactNode;
}

export const VideoProvider: React.FC<VideoProviderProps> = ({children}) => {
    const [videos, setVideos] = useState<Video[]>([]);

    const fetchLibrary = async () => {
        try {
            const response: AxiosResponse = await axios.get("http://localhost:8080/");

            const videoLibrary: Video[] = Object.keys(response.data).map((key: string) => {
                const videoId: string = sha256(key).slice(0, 10);

                return {
                    id: videoId,
                    title: key,
                    imagePath: '/assets/images/video-placeholder.png',
                    videoPath: response.data[key],
                };
            });

            setVideos(videoLibrary);
        } catch (error) {
            handleApiError(error);
        }
    };

    useEffect(() => {
        void fetchLibrary();
    }, []);

    return (
        <VideoContext.Provider value={{videos, setVideos}}>
            {children}
        </VideoContext.Provider>
    );
}

export const useVideo = (): VideoContextType => {
    const context = useContext(VideoContext);
    if (!context) {
        throw new Error("useVideo must be used within a VideoProvider");
    }
    return context;
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