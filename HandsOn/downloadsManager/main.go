package main

import (
	"fmt"
	"math/rand"
	"time"
)

func downloadFile(fileId int, fileName string, downloadCompleteCh chan string) {
	fileDownloadTime := time.Duration(rand.Intn(5000)) * time.Millisecond
	steps := 10
	stepTime := fileDownloadTime / time.Duration(steps)

	fmt.Printf("Download started for file: %s\n", fileName)
	for i := 1; i <= steps; i++ {
		time.Sleep(stepTime)
		progress := float64(i) / float64(steps) * 100
		fmt.Printf("Progress for %s: %.1f%%\n", fileName, progress)
	}
	downloadCompleteCh <- fmt.Sprintf("File: %s", fileName)
}

func main() {

	files := [5]string{"game", "pdf", "excel", "jpg", "docx"}
	numFiles := len(files)
	downloadCompleteCh := make(chan string, numFiles)

	for i := range numFiles {
		go downloadFile(i+1, files[i], downloadCompleteCh)
	}

	for range numFiles {
		select {
		case file := <-downloadCompleteCh:
			fmt.Printf("Download complete for %s\n", file)
		case <-time.After(1 * time.Second):
			fmt.Println("Download is taking too long")
		}
	}
}
