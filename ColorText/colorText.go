package ColorText

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fogleman/gg"
	"golang.org/x/image/font/gofont/goregular"
)

func ColorText() {
	const W = 500
	const H = 300

	fmt.Println("Adding Text")

	// Create a temporary file and write the byte slice to it
	tempFile, err := ioutil.TempFile("", "font-*.ttf")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(goregular.TTF); err != nil {
		panic(err)
	}

	dc := gg.NewContext(W, H)

	if err := dc.LoadFontFace(tempFile.Name(), 72); err != nil {
		panic(err)
	}

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	fmt.Print("Adding the string \"COLORS\" to the image")

	dc.SetRGB(.5, 0, 0)
	dc.DrawStringAnchored("COLORS", W/2, H/2, 0.5, 0.5)
	dc.Stroke()

	fmt.Println("...done")
	fmt.Print("Saving labeled image to colors_labeled.jpg")

	dc.SavePNG("colors_labeled.jpg")

	fmt.Println("...done")
}
