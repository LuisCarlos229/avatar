package saveimage

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

type SaveImage struct{}

func (S *SaveImage) SaveImage(path string, img *image.RGBA) error {
	fmt.Println()
	log.Println("save image in proces...")

	file, err := os.OpenFile(path, os.O_RDWR, 0)
	defer file.Close()
	if err != nil {
		return err
	}
	//ima := *img
	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	log.Println("save image in ready...")

	return err
}
