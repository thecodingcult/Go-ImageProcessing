package main

import (
	"fmt"
	"image/jpeg"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	path := "image/googlelogo.png"

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding original image: %v\n", err)
		return
	}

	mat, err := gocv.ImageToMatRGBA(img)
	if err != nil {
		fmt.Printf("Error converting image to Mat: %v\n", err)
		return
	}

	gray := gocv.NewMat()
	gocv.CvtColor(mat, &gray, gocv.ColorBGRToGray)

	gocv.Canny(gray, &gray, 50, 150)

	out, err := os.Create("result.jpg")
	if err != nil {
		fmt.Printf("Error creating result file: %v\n", err)
		return
	}
	defer out.Close()

	if err = jpeg.Encode(out, gray, &jpeg.Options{Quality: jpeg.DefaultQuality}); err != nil {
		fmt.Printf("Error encoding image: %v\n", err)
		return
	}

	// Close the mats
	mat.Close()
	gray.Close()
}
