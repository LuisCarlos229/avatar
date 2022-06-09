package colors

import (
	"image/color"
)

// defino los colores con los que podre trabajar
var (
	red        = color.RGBA{255, 0, 0, 0xff}
	dark       = color.RGBA{0, 0, 0, 0xff}
	dark_green = color.RGBA{0, 100, 0, 0xff}

	lime              = color.RGBA{0, 255, 0, 0xff}
	verde             = color.RGBA{0, 128, 0, 0xff}
	medium_slate_blue = color.RGBA{123, 104, 238, 0xff}

	deep_sky_blue = color.RGBA{0, 191, 255, 0xff}
	aqua          = color.RGBA{0, 255, 255, 0xff}
	steel_blue    = color.RGBA{70, 130, 180, 0xff}

	dark_violet   = color.RGBA{148, 0, 211, 0xff}
	pink          = color.RGBA{255, 192, 203, 0xff}
	lemon_chiffon = color.RGBA{255, 250, 205, 0xff}

	black = color.RGBA{0, 0, 0, 0xff}
	gris  = color.RGBA{160, 160, 160, 0xff}
	azure = color.RGBA{240, 255, 255, 0xff}

	// te color of te winers
	Gold = color.RGBA{255, 215, 0, 0xff}
)

type paleta struct {
	c1 color.RGBA
	c2 color.RGBA
	c3 color.RGBA
}

type paletas struct {
	p1 paleta
	p2 paleta
	p3 paleta
	p4 paleta
}

var (
	paleta1 = paleta{
		c1: red,
		c2: dark,
		c3: dark_green,
	}
	paleta2 = paleta{
		c1: lime,
		c2: verde,
		c3: medium_slate_blue,
	}
	paleta3 = paleta{
		c1: deep_sky_blue,
		c2: aqua,
		c3: steel_blue,
	}
	paleta4 = paleta{
		c1: pink,
		c2: dark_violet,
		c3: lemon_chiffon,
	}
	paleta5 = paleta{
		c1: black,
		c2: gris,
		c3: azure,
	}
	paletas1 = paletas{
		p1: paleta1,
		p2: paleta2,
		p3: paleta3,
		p4: paleta4,
	}
)

type SelectColor struct{}

func (S *SelectColor) SelectColor(mibyte byte) (color.RGBA, color.RGBA, color.RGBA) {
	if mibyte <= 76 {
		return paleta1.c1, paleta1.c2, paleta1.c3

	} else if 76 < mibyte && mibyte <= 102 {
		return paleta2.c1, paleta2.c2, paleta2.c3

	} else if 102 < mibyte && mibyte <= 153 {
		return paleta3.c1, paleta3.c2, paleta3.c3

	} else if 153 < mibyte && mibyte <= 204 {
		return paleta4.c1, paleta4.c2, paleta4.c3
	}
	return paleta5.c1, paleta5.c2, paleta5.c3
}
