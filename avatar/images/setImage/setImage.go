package setimage

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/LuisCarlos229/avatar/images/colors"
)

type SetImage struct{}

func (s *SetImage) SetImage(encodeInformation uint64, color1 color.RGBA, color2 color.RGBA, color3 color.RGBA) *image.RGBA {
	fmt.Println()
	log.Println("set image in proces...")

	ancho := 4 //width-ancho
	alto := 5  //height-alto

	upLeft := image.Point{0, 0}
	lowRight := image.Point{ancho, alto}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	//color1, color2, color3 := colors.SelectColor((encodeInformation / 1000) % 10)

	var (
		digito       uint64
		interaciones int
		acu          uint64 = 1
	)
	const long = 19

	for x := 0; x < ancho; x++ {
		interaciones = 0
		for y := 0; y < alto; y++ {
			digito = (encodeInformation / acu) % 10
			acu *= 10
			switch {
			case digito%2 == 0 && interaciones < long: // si es par
				if digito != 2 {
					img.Set(x, y, color1)
				} else {
					img.Set(x, y, colors.Gold)
				}
			case digito%2 != 0 && interaciones < long: // si es impar
				if digito != 3 && digito != 5 && digito != 7 {
					img.Set(x, y, color2)
				} else {
					img.Set(x, y, colors.Gold)
				}
			default:
				img.Set(x, y, color3)
			}
			interaciones++
		}
	}

	log.Println("set image is ready...")

	return img
}
