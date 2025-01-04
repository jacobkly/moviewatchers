// Package services provides functions for managing a video library, including populating the
// library, checking if files are hidden, and converting the library to JSON.
package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/google/uuid"
)

// Episode represents an episode of a TV show with a title and video path.
type Episode struct {
	Title     string `json:"title"`
	VideoPath string `json:"videoPath"`
}

// Show represents a TV show with a unique ID, title, image path, and episodes.
type Show struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	ImagePath string    `json:"imagePath"`
	Episodes  []Episode `json:"episodes"`
	ItemType  string    `json:"type"`
}

// Movie represents a movie with a unique ID, title, image path, and video path.
type Movie struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	ImagePath string `json:"imagePath"`
	VideoPath string `json:"videoPath"`
	ItemType  string `json:"type"`
}

// Video is a generic interface that can represent either a Movie or a Show.
type Video interface{}

// library is a global slice that stores all videos in the library.
var library []Video

// PopulateLibrary populates the global video library by reading files and directories
// from the given filePath. It differentiates between movies and shows.
func PopulateLibrary(filePath string) error {
	files, err := os.ReadDir(filePath)
	if err != nil {
		return fmt.Errorf("issue reading directory %s: %v", filePath, err)
	}

	for _, file := range files {
		newFilePath := filepath.Join(filePath, file.Name())
		isSkippable, _ := isSkippable(newFilePath)
		if isSkippable {
			continue
		}

		if file.IsDir() {
			addShow(newFilePath, file.Name())
		} else {
			addMovie(newFilePath, file.Name())
		}
	}
	return nil
}

// JsonLibrary returns the global library as a JSON-encoded byte slice.
// It returns an error if the library is empty or if JSON marshalling fails.
func JsonLibrary() ([]byte, error) {
	if len(library) == 0 {
		return nil, fmt.Errorf("empty library")
	}
	jsonLibrary, err := json.Marshal(library)
	if err != nil {
		return nil, err
	}
	return jsonLibrary, nil
}

// addShow adds a show to the global library by reading its episodes from the given file path.
func addShow(filePath string, showName string) {
	episodes, err := addEpisodes(filePath)
	if err != nil {
		fmt.Printf("error adding episodes: %v", err)
		return
	}
	show := Show{
		Id:        generateID(),
		Title:     RemoveFileExtension(showName),
		ImagePath: "/assets/images/video-placeholder.png", // hard coded for React app
		Episodes:  episodes,
		ItemType:  "show",
	}
	library = append(library, show)
}

// addMovie adds a movie to the global library from the given file path and name.
func addMovie(filePath string, movieName string) {
	movie := Movie{
		Id:        generateID(),
		Title:     RemoveFileExtension(movieName),
		ImagePath: "/assets/images/video-placeholder.png", // hard coded for React app
		VideoPath: filePath,
		ItemType:  "movie",
	}
	library = append(library, movie)
}

// addEpisodes reads and creates a list of episodes from the given file path.
// It returns the list of episodes or an error if reading the directory fails.
func addEpisodes(filePath string) ([]Episode, error) {
	files, err := os.ReadDir(filePath)
	if err != nil {
		return nil, fmt.Errorf("issue reading directory %s: %v", filePath, err)
	}

	var episodes []Episode
	for _, file := range files {
		newFilePath := filepath.Join(filePath, file.Name())
		isSkippable, _ := isSkippable(newFilePath)
		if isSkippable {
			continue
		}

		episode := Episode{
			Title:     RemoveFileExtension(file.Name()),
			VideoPath: newFilePath,
		}
		episodes = append(episodes, episode)
	}
	return episodes, nil
}

// isSkippable checks if a file or directory should be skipped.
func isSkippable(filePath string) (bool, error) {
	// Check for directories that should be skipped
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return true, nil
	}
	if fileInfo.IsDir() {
		files, err := os.ReadDir(filePath)
		if err != nil {
			return true, err
		}
		for _, file := range files {
			if file.Name() == "subtitles" || file.Name() == "raw" {
				return true, nil
			}
		}
	}

	// Check file extension for non-compatible types
	fileExt := filepath.Ext(filePath)
	switch fileExt {
	case ".mp4", ".mov", ".avi", ".mkv", ".webm":
		// if true, continue the function
	default:
		return true, nil
	}

	// Check for hidden and system file attributes (Windows-specific)
	pointer, err := syscall.UTF16PtrFromString(filePath)
	if err != nil {
		return false, err
	}
	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return false, err
	}
	if attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0 ||
		attributes&syscall.FILE_ATTRIBUTE_SYSTEM != 0 {
		return true, nil
	}

	return false, nil
}

// generateID generates a unique ID for a video item by using the UUID package.
// It returns the first 10 characters of the UUID.
func generateID() string {
	return uuid.New().String()[:10]
}
