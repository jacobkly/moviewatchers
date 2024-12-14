package services

import (
	"encoding/json"
	"fmt"
)

// idea from early on in main:
// adding JSON object to req body which frontend can use to populate
// categorize with "Anime", "Movies", "Shows"

// example: adding to library (easy map operations)
//	movieLibrary.FilesJSON["Anime"] = map[string]interface{}{
//		"Kids on the Slope": map[string]interface{}{
//			"Episode 1": "ep1.mp4",
//			"Episode 2": "ep2.mp4",
//		},
//	}
//	movieLibrary.FilesJSON["Movies"] = map[string]interface{}{
//		"The Pianists (2002)": "thepianist.mp4",
//	}

type MovieLibrary struct {
	Library map[string]interface{}
}

var lib = MovieLibrary{
	Library: make(map[string]interface{}),
}

func PopulateJSON(storagePath string) {
	lib.Library["Anime"] = "Kids on the Slope" // test
}

func JsonMovieLibrary() ([]byte, error) {
	if len(lib.Library) == 0 {
		return nil, fmt.Errorf("Empty library")
	}

	jsonLib, err := json.Marshal(lib.Library)
	if err != nil {
		return nil, err
	}
	return jsonLib, nil
}

// might move to "media_player" services
func playVideo() {

}
