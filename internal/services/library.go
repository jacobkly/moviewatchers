package services

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

var movieLibrary = make(map[string]interface{})

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

		if file.IsDir() {
			movieLibrary[file.Name()], _ = populateMap(newFullPath)
		} else {
			movieLibrary[file.Name()] = newFullPath
		}
	}
	return nil
}

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
		mapResult[file.Name()] = newFullPath
	}
	return mapResult, nil
}

// only windows compatible
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

func JsonMovieLibrary() ([]byte, error) {
	if len(movieLibrary) == 0 {
		return nil, fmt.Errorf("empty library")
	}

	jsonLibrary, err := json.Marshal(movieLibrary)
	if err != nil {
		return nil, err
	}
	return jsonLibrary, nil
}

func PlayVideo(filePath []byte) error {
	filePathString := string(filePath)

	cmd := exec.Command("C:\\Program Files\\VideoLAN\\VLC\\vlc.exe", filePathString)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error playing video: %v", err)
	}
	return nil
}
