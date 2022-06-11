package setimage

import (
	"image"
	"image/color"

	"github.com/LuisCarlos229/avatar/images/colors"
)

type SetImage struct{}

// Setimage se encarga de setear los pixeles de la imagen creada con los colores y la cadena de bytes pasada.
func (s *SetImage) Setimage(encodeInformation []byte, color1 color.RGBA, color2 color.RGBA, color3 color.RGBA) *image.RGBA {

	long := len(encodeInformation)
	ancho := 60 //width-ancho
	alto := 60  //height-alto

	upLeft := image.Point{0, 0}
	lowRight := image.Point{ancho, alto}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	var (
		mibyte       byte
		interaciones int
	)

	x1, y1 := 0, 0
	for x := 0; x < ancho; x++ {
		x1 = x * 10
		for y := 0; y < alto; y++ {
			y1 = y * 10
			if interaciones > long-1 {
				interaciones = 0
			}
			mibyte = encodeInformation[interaciones]
			switch {
			case mibyte == 0 || mibyte == 85 || mibyte == 170 || mibyte == 255:
				setCuadros(x1, y1, colors.Gold, img)
			case mibyte < 85:
				setCuadros(x1, y1, color1, img)
			case 85 < mibyte && mibyte < 170:
				setCuadros(x1, y1, color2, img)
			case 170 < mibyte && mibyte < 255:
				setCuadros(x1, y1, color3, img)
			default:
				//...
			}
			interaciones++
		}
	}

	return img
}

func setCuadros(x1, y1 int, colorSeleccionado color.RGBA, im *image.RGBA) {
	for i := x1; i < (x1 + 10); i++ {
		for j := y1; j < (y1 + 10); j++ {
			im.Set(i, j, colorSeleccionado)
		}
	}
}
