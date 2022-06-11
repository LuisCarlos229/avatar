package saveimage

import (
	"image"
	"image/png"
	"os"
)

type SaveImage struct{}

// Saveimage se encarga de guargdar la imagen seteada en el archivo png idicado por el path.
func (S *SaveImage) Saveimage(path string, img *image.RGBA) (err error) {

	file, err := os.OpenFile(path, os.O_RDWR, 0)
	defer file.Close()
	if err != nil {
		return err
	}

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return err
}
