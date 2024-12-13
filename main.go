package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var dirPath string
	fmt.Print("Enter dir path for USB flash drive (ex: \"F:\\\"): ")
	_, err := fmt.Scanln(&dirPath)
	if err != nil {
		log.Fatal(err)
	}

	dirPath = filepath.Clean(dirPath) // normalize path for cross-platform compatibility
	fmt.Println()
	printFiles(dirPath)
	fmt.Println()

	reader := bufio.NewReader(os.Stdin) // bufio for reading input with spaces
	for {
		fmt.Println("Choose the file path to open in VLC (or \"q\" to quit): ")
		filePath, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		filePath = filepath.Clean(filePath[:len(filePath)-1]) // trim null char
		if filePath == "q" {
			break
		}

		fmt.Println("Running file path:", filePath)
		cmd := exec.Command("C:\\Program Files\\VideoLAN\\VLC\\vlc.exe", filePath)
		if err := cmd.Run(); err != nil {
			fmt.Println("Error running VLC:", err)
		}
	}

	fmt.Println("Good day sir")
}

// Recursively print files
func printFiles(dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Printf("Error reading directory '%s': %v\n", dirPath, err)
		return
	}

	for _, file := range files {
		temp := filepath.Join(dirPath, file.Name())
		fmt.Println(temp)

		if file.IsDir() {
			printFiles(temp)
		}
	}
}
