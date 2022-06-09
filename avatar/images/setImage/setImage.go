package setimage

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/LuisCarlos229/avatar/images/colors"
)

type SetImage struct{}

func (s *SetImage) SetImage(encodeInformation []byte, color1 color.RGBA, color2 color.RGBA, color3 color.RGBA) *image.RGBA {
	fmt.Println()
	log.Println("set image in proces...")

	long := len(encodeInformation)
	ancho := long / 5 //width-ancho
	alto := long / 6  //height-alto

	upLeft := image.Point{0, 0}
	lowRight := image.Point{ancho, alto}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	var (
		mibyte       byte
		interaciones int
	)

	for x := 0; x < ancho; x++ {
		if interaciones > long {
			interaciones = 0
		}
		for y := 0; y < alto; y++ {
			mibyte = encodeInformation[interaciones]
			switch {
			case mibyte == 0 || mibyte == 85 || mibyte == 170 || mibyte == 255:
				img.Set(x, y, colors.Gold)

			case mibyte < 85:
				img.Set(x, y, color1)

			case 85 < mibyte && mibyte < 170:
				img.Set(x, y, color2)

			case 170 < mibyte && mibyte < 255: //&& interaciones < long
				img.Set(x, y, color3)

			default:
				img.Set(x, y, color.White)
			}
			interaciones++
		}
	}

	log.Println("set image is ready...")

	return img
}
