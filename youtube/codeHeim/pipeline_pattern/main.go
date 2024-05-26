package main

import (
	"image"
	"log"
	"strings"
	"time"
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
			log.Printf("(loadImage) path = [%v]\n", p)
			job := Job{
				InputPath: p,
				OutPath:   strings.Replace(p, "images/", "images/output/", 1),
			}
			job.Image = imageprocessing.ReadImage(p)
			out <- job
		}
		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input { // Read from the channel
			log.Printf("(resize) path = [%v]\n", job.InputPath)
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input { // Read from the channel
			log.Printf("(convertToGrayscale) path = [%v]\n", job.InputPath)
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input { // Read from the channel
			log.Printf("(saveImage) path = [%v]\n", job.InputPath)
			imageprocessing.WriteImage(job.OutPath, job.Image)
			out <- true
		}
		close(out)
	}()
	return out
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		log.Printf("duration = %v msec\n", duration.Milliseconds())
	}()

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
			log.Println("Success!")
		} else {
			log.Println("Failed!")
		}
	}
}
