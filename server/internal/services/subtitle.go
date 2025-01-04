package services

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ProcessSubtitlePath(filePath string) (string, error) {
	fmt.Println("TESTING")

	// need to check if the subtitles directory exists
	_, err := os.ReadDir("F:\\subtitles")
	if err != nil {
		fmt.Printf("issue reading directory %s: %v", filePath, err)
		return "", err
	}

	filePath = filepath.Clean(filePath)
	fileSplit := strings.Split(filePath, "\\")
	fileName := fileSplit[len(fileSplit)-1]
	newFilePath := "F:\\subtitles\\" + RemoveFileExtension(fileName)
	// newFilePath := fmt.Sprintf("F:\\subtitles\\%s", RemoveFileExtension(fileName))

	// fix to either add ".srt" or ".vtt" depending on content type
	processedPath := newFilePath + ".srt"
	// processedPath := fmt.Sprintf("%s.srt", newFilePath)

	// need to check if the processed subtitle path exists

	fmt.Println(processedPath)
	return processedPath, nil
}
