package colors

import (
	"image/color"
)

// defino los colores con los que podre trabajar
var (
	red        = color.RGBA{255, 0, 0, 0xff}
	anaranjado = color.RGBA{255, 165, 0, 0xff}
	amarillo   = color.RGBA{255, 255, 0, 0xff}

	lime  = color.RGBA{0, 255, 0, 0xff}
	verde = color.RGBA{0, 167, 0, 0xff}
	azul  = color.RGBA{0, 167, 0, 0xff}

	sky_blue    = color.RGBA{135, 206, 235, 0xff}
	aqua        = color.RGBA{0, 255, 255, 0xff}
	aqua_marine = color.RGBA{127, 255, 212, 0xff}

	deep_pink = color.RGBA{255, 20, 147, 0xff}
	pink      = color.RGBA{255, 192, 203, 0xff}
	violeta   = color.RGBA{153, 0, 153, 0xff}

	black  = color.RGBA{0, 0, 0, 0xff}
	gris   = color.RGBA{160, 160, 160, 0xff}
	blanco = color.RGBA{255, 255, 255, 0xff}

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
		c2: anaranjado,
		c3: amarillo,
	}
	paleta2 = paleta{
		c1: lime,
		c2: verde,
		c3: azul,
	}
	paleta3 = paleta{
		c1: sky_blue,
		c2: aqua,
		c3: aqua_marine,
	}
	paleta4 = paleta{
		c1: pink,
		c2: deep_pink,
		c3: violeta,
	}
	paleta5 = paleta{
		c1: black,
	}
	paletas1 = paletas{
		p1: paleta1,
		p2: paleta2,
		p3: paleta3,
		p4: paleta4,
	}
)

type SelectColor struct{}

func (S *SelectColor) SelectColor(digito uint64) (color.RGBA, color.RGBA, color.RGBA) {
	if digito == 0 || digito == 5 {
		return paleta1.c1, paleta1.c2, paleta1.c3

	} else if digito == 1 || digito == 6 {
		return paleta2.c1, paleta2.c2, paleta2.c3

	} else if digito == 2 || digito == 7 {
		return paleta3.c1, paleta3.c2, paleta3.c3

	} else if digito == 3 || digito == 8 {
		return paleta4.c1, paleta4.c2, paleta4.c3
	}
	return paleta5.c1, paleta5.c2, paleta5.c3
}
