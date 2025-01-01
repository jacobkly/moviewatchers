// Package routes provides the router and routes for the moviewatcher application.
package routes

import (
	"fmt"
	"github.com/jacobkly/moviewatchers/server/internal/services"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// NewRouter creates and returns a new HTTP router with two routes.
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", libraryDisplayHandler)
	mux.HandleFunc("/video", videoFileHandler)
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

// videoFileHandler handles the "/video" route, which serves video files based on the "path" query parameter.
// It checks if the file exists and serves it with the appropriate content type (MP4 or MKV).
// If the content type is unsupported or the file doesn't exist, it returns an appropriate error.
func videoFileHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	videoPath := r.URL.Query().Get("path")
	filePath := filepath.Clean(videoPath)
	if _, err := os.Stat(filePath); err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	contentType := getContentType(filePath)
	if contentType == "video/mp4" || contentType == "video/x-matroska" {
		w.Header().Set("Content-Type", contentType)
		http.ServeFile(w, r, filePath)
	} else {
		http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
	}
}

// getContentType determines the content type based on the file extension.
// It supports MP4 and MKV file types. If the file has an unsupported extension,
// it returns "application/octet-stream".
func getContentType(filePath string) string {
	var contentType string
	if strings.HasSuffix(filePath, ".mp4") {
		contentType = "video/mp4"
	} else if strings.HasSuffix(filePath, ".mkv") {
		contentType = "video/x-matroska"
	} else {
		contentType = "application/octet-stream"
	}
	return contentType
}

// enableCors sets the "Access-Control-Allow-Origin" header to "*" to allow cross-origin requests.
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
