package images

import (
	"image"
	"image/color"
	"log"

	"github.com/LuisCarlos229/avatar/images/colors"
	createimage "github.com/LuisCarlos229/avatar/images/createImage"
	saveimage "github.com/LuisCarlos229/avatar/images/saveImage"
	setimage "github.com/LuisCarlos229/avatar/images/setImage"
)

type selectColors interface {
	// SelectColor resibe un digito por el cual ara la seleccion de una paleta de 3 colores para el usuario.
	SelectColor(mibyte byte) (color.RGBA, color.RGBA, color.RGBA)
}
type sreateImage interface {
	// Createimage se encarga de la generacion del archivo png vacio para su posterior seteo, y devuelve el path del archivo
	// y un error en caso de que algo aya salido mal.
	Createimage(encodeInformation []byte) (path string, err error)
}
type setImage interface {
	// Setimage se encarga de setear los pixeles de la imagen creada con los colores y la cadena de bytes pasada.
	Setimage(encodeInformation []byte, color1 color.RGBA, color2 color.RGBA, color3 color.RGBA) *image.RGBA
}
type saveImage interface {
	// Saveimage se encarga de guargdar la imagen seteada en el archivo png idicado por el path.
	Saveimage(path string, img *image.RGBA) error
}

type BuildImage struct {
	selectColors
	sreateImage
	setImage
	saveImage
}

func newBuildImage() *BuildImage {
	return &BuildImage{
		&colors.SelectColor{},
		&createimage.Createimage{},
		&setimage.SetImage{},
		&saveimage.SaveImage{},
	}
}

// NewCustomImage genera una nueva estructura para que el usuario
//pueda crear una imagen personalizada.
func NewCustomImage() *BuildImage {
	return &BuildImage{
		&colors.SelectColor{},
		&createimage.Createimage{},
		&setimage.SetImage{},
		&saveimage.SaveImage{},
	}
}

//BuildAndSaveImage se encarga de la creacion de la imagen,
// eleccion de colores, seteo y guardado de esta.
func (B *BuildImage) BuildAndSaveImage(encodeInformation []byte) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error en tiempo de ejecucion==>", r)
		}
	}()

	imaProses1 := newBuildImage()
	path, err := imaProses1.Createimage(encodeInformation)
	if err != nil {
		return err
	}

	color1, color2, color3 := imaProses1.SelectColor(encodeInformation[10])
	img := imaProses1.Setimage(encodeInformation, color1, color2, color3)

	err = imaProses1.Saveimage(path, img)
	if err != nil {
		return err
	}

	return nil
}
