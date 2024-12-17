// Package services provides functions for managing a video library, including populating the
// library, checking if files are hidden, converting the library to JSON, and playing videos.
package services

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

// videoLibrary holds the video library as a map of file names to their paths.
var videoLibrary = make(map[string]interface{})

// PopulateJSON populates the videoLibrary map by recursively reading a directory and its contents.
// It excludes hidden files and directories and stores file paths in the videoLibrary map.
func PopulateJSON(filePath string) error {
	files, err := os.ReadDir(filePath)
	if err != nil {
		return fmt.Errorf("issue reading directory %s: %v", filePath, err)
	}

	for _, file := range files {
		newFullPath := filepath.Join(filePath, file.Name())
		hidden, _ := isHidden(newFullPath)
		if hidden {
			continue
		}

		fileName := removeFileExtension(file.Name())
		if file.IsDir() {
			videoLibrary[fileName], _ = populateMap(newFullPath)
		} else {
			videoLibrary[fileName] = newFullPath
		}
	}
	return nil
}

// populateMap is a helper function that populates a map with the file paths from a given directory.
// It excludes hidden files and directories. Note that it is fully compatible with Windows.
func populateMap(filePath string) (map[string]interface{}, error) {
	mapResult := make(map[string]interface{})

	files, err := os.ReadDir(filePath)
	if err != nil {
		return nil, fmt.Errorf("issue reading directory %s: %v", filePath, err)
	}

	for _, file := range files {
		newFullPath := filepath.Join(filePath, file.Name())
		hidden, _ := isHidden(newFullPath)
		if hidden {
			continue
		}
		mapResult[removeFileExtension(file.Name())] = newFullPath
	}
	return mapResult, nil
}

// isHidden checks if a given file or directory is hidden on the Windows filesystem.
// It returns true if the file is hidden or a system file, and false otherwise.
func isHidden(filePath string) (bool, error) {
	pointer, err := syscall.UTF16PtrFromString(filePath)
	if err != nil {
		return false, err
	}
	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return false, err
	}
	return (attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0) ||
		(attributes&syscall.FILE_ATTRIBUTE_SYSTEM != 0), nil
}

// removeFileExtension removes the file extension from a given filename string.
// It returns the filename without the extension by trimming the suffix returned by filepath.Ext.
func removeFileExtension(file string) string {
	return strings.TrimSuffix(file, filepath.Ext(file))
}

// JsonVideoLibrary returns the videoLibrary map as a JSON-encoded byte slice.
// It returns an error if the library is empty or if marshalling fails.
func JsonVideoLibrary() ([]byte, error) {
	if len(videoLibrary) == 0 {
		return nil, fmt.Errorf("empty library")
	}

	jsonLibrary, err := json.Marshal(videoLibrary)
	if err != nil {
		return nil, err
	}
	return jsonLibrary, nil
}

// PlayVideo attempts to play a video using VLC from the provided file path.
// It returns an error if there is an issue running the VLC command.
func PlayVideo(filePath []byte) error {
	filePathString := string(filePath)

	cmd := exec.Command("C:\\Program Files\\VideoLAN\\VLC\\vlc.exe", filePathString)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error playing video: %v", err)
	}
	return nil
}
