import React from 'react';
import {createContext, useContext, useState, useEffect, ReactNode} from 'react';
import axios, {AxiosResponse} from "axios";
import {Video} from '../types/Video';

interface VideoContextType {
    videos: Video[];
    setVideos: (videos: Video[]) => void;
}

const VideoContext = createContext<VideoContextType | undefined>(undefined);

interface VideoProviderProps {
    children: ReactNode;
}

export const VideoProvider: React.FC<VideoProviderProps> = ({children}: VideoProviderProps) => {
    const [videos, setVideos] = useState<Video[]>([]);

    const fetchLibrary = async () => {
        try {
            const response: AxiosResponse = await axios.get("http://localhost:8080/");
            setVideos(response.data);
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
    const context: VideoContextType | undefined = useContext(VideoContext);
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
