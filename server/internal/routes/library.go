// Package routes provides the router and routes for the moviewatcher application.
package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jacobkly/moviewatchers/server/internal/services"
)

// NewRouter creates and returns a new HTTP router with two routes:
// - "/" to display the movie library in JSON format
// - "/play" to play a video
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", libraryDisplayHandler)
	mux.HandleFunc("/play", playVideoHandler)
	return mux
}

// libraryDisplayHandler handles the "/" route, which returns the user's movie library in JSON format.
// If the library is empty, it returns a 404 status code with an appropriate message.
// If there is an error fetching the library, it returns a 500 status code and the error message.
func libraryDisplayHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	jsonLibrary, err := services.JsonLibrary()
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

// playVideoHandler handles the "/play" route, which attempts to play a video file specified in the
// request body. It reads the video file path from the request body and attempts to play it using
// the PlayVideo service. If an error occurs while reading the body or playing the video, it
// returns an appropriate status code (such as 400, 404, or 500) along with an error message.
func playVideoHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintln(w, "Error reading body: ", err)
		return
	}
	err = services.PlayVideo(rBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintln(w, "Error playing video: ", err)
	}
}

// enableCors sets the "Access-Control-Allow-Origin" header to "*" to allow  cross-origin requests.
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
