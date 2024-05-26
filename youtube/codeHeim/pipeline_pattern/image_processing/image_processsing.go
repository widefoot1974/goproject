package imageprocessing

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/nfnt/resize"
)

func ReadImage(path string) image.Image {
	inputFile, err := os.Open(path)
	if err != nil {
		log.Panicf("os.Open(%v) Failed: %v\n", path, err)
	}
	defer inputFile.Close()

	// Decode the image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		log.Panicf("image.Decode(%v) Failed: %v\n", path, err)
	}

	return img
}

func WriteImage(path string, img image.Image) {
	outputFile, err := os.Create(path)
	if err != nil {
		log.Panicf("os.Create(%v) Failed: %v\n", path, err)
	}
	defer outputFile.Close()

	// Encode the image to the new file
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		log.Panicf("jpeg.Encode(%v) Failed: %v\n", path, err)
	}
}

func Grayscale(img image.Image) image.Image {
	// Crate a new grayscale image
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}
	return grayImg
}

func Resize(img image.Image) image.Image {
	newWidth := uint(500)
	newHeight := uint(500)
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	return resizedImg
}
