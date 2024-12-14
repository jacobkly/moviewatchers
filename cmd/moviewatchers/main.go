package main

import (
	"encoding/json"
	"fmt"
)

type MovieLibrary struct {
	FilesJSON map[string]interface{}
}

var movieLibrary = MovieLibrary{
	FilesJSON: make(map[string]interface{}),
}

func main() {
	movieLibrary.FilesJSON["Anime"] = map[string]interface{}{
		"Kids on the Slope": map[string]interface{}{
			"Episode 1": "ep1.mp4",
			"Episode 2": "ep2.mp4",
		},
	}
	movieLibrary.FilesJSON["Movies"] = map[string]interface{}{
		"The Pianists (2002)": "thepianist.mp4",
	}

	jsonMovieLibrary, _ := json.Marshal(movieLibrary)
	fmt.Printf("json movie library: %s\n", jsonMovieLibrary)
}
