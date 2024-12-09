package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// get usb directory
	var dirPath string
	fmt.Print("Enter dir path for USB flash drive (ex: \"F:\"): ")
	_, err := fmt.Scan(&dirPath)

	// read and list all files on usb
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirPath, " contained these files below:")
	for _, file := range files {
		fmt.Println(file.Name())
	}

	for {
		// open one of the files in VLC
		var videoFile string
		fmt.Println("Choose the .mp4 file to open in VLC (or \"q\" to quit): ")
		_, err = fmt.Scan(&videoFile)

		if videoFile == "q" {
			break
		}

		cmd := exec.Command("C:\\Program Files\\VideoLAN\\VLC\\vlc.exe", dirPath+"\\"+videoFile)
		if err := cmd.Run(); err != nil { // start command
			fmt.Println(err)
		}
	}

	fmt.Println("Good day sir...")
}
