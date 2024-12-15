// Package main including initial testing of moviewatchers project concept.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// console prompts the user to input a directory path for a USB flash drive, then displays the
// files in that directory. The user is then prompted to choose a file path to open in VLC. The
// function will continue to allow the user to open files until the user enters "q" to quit.
// If an error occurs, it will log the error and continue the process.
func console() {
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

// printFiles recursively prints the paths of all files in a directory, including files in
// subdirectories. It handles errors that occur during the reading of directories and logs them.
// The function will traverse the directory structure and print the path of each file.
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
