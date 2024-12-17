export interface Episode {
    title: string;
    path: string;
}

export interface Show {
    id: string;
    title: string;
    imagePath: string;
    episodes: Episode[];
}

export interface Movie {
    id: string;
    title: string;
    imagePath: string;
    videoPath: string;
}

export type Video = Show | Movie;