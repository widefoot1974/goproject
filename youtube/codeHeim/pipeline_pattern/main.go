package main

import (
	"fmt"
	"image"
	"strings"
	imageprocessing "youtube/codeHeim/pipeline_pattern/image_processing"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input path create a job abd add it to
		// the out channel
		for _, p := range paths {
			job := Job{
				InputPath: p,
				OutPath:   strings.Replace(p, "images/", "images/output/", 1),
			}
			job.Image = imageprocessing.ReadImage(p)
		}
	}()
}

func main() {
	imagePaths := []string{
		"images/cat1.jpg",
		"images/cat2.jpg",
		"images/cat3.jpg",
		"images/cat4.jpg",
	}

	channel1 := loadImage(imagePaths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)
	writeResults := saveImage(channel3)

	for success := range writeResults {
		if success {
			fmt.Println("Success!")
		} else {
			fmt.Println("Failed!")
		}
	}
}
