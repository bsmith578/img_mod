package Colors

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png" // import this package to decode PNGs
	"os"
)

func GetRGB() {

	fmt.Println("Processing Pixel Colors")

	reader, err := os.Open("colors.jpg")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Create a new file to save the image
	outputFile, err := os.Create("color_pixel_counts.txt")
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer outputFile.Close()

	fmt.Println("Saving Pixel Values to color_pixel_counts.txt")

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			vals := fmt.Sprintf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8)
			outputFile.WriteString(vals)
		}
	}

	fmt.Print("...done")
}
