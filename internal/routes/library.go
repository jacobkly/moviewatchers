package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jacobkly/moviewatchers/internal/services"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", libraryDisplayHandler)
	mux.HandleFunc("/play", playVideoHandler)

	return mux
}

func libraryDisplayHandler(w http.ResponseWriter, r *http.Request) {
	jsonLibrary, err := services.JsonMovieLibrary()
	if err != nil {
		if err.Error() == "Empty library" {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintln(w, "No JSON library found")
		} else {
			w.WriteHeader(http.StatusInternalServerError) // 500
			fmt.Fprintln(w, "Error fetching JSON library: ", err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, writeErr := w.Write(jsonLibrary)
	if writeErr != nil {
		fmt.Println("Error writing response: ", writeErr)
	}
}

// might move to "media_player" route
func playVideoHandler(w http.ResponseWriter, r *http.Request) {
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Error reading body: ", err)
		return
	}
	err = services.PlayVideo(rBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Error playing video: ", err)
	}
}
