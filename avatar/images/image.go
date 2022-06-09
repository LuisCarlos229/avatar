package images

import (
	"image"
	"image/color"

	"github.com/LuisCarlos229/avatar/images/colors"
	createimage "github.com/LuisCarlos229/avatar/images/createImage"
	saveimage "github.com/LuisCarlos229/avatar/images/saveImage"
	setimage "github.com/LuisCarlos229/avatar/images/setImage"
)

/*
	"github.com/LuisCarlos229/avatar/images/colors"
	"github.com/LuisCarlos229/avatar/images/createimage"
	"github.com/LuisCarlos229/avatar/images/setImage"
	"github.com/LuisCarlos229/avatar/images/saveImage"
*/

type selectColor interface {
	SelectColor(mibyte byte) (color.RGBA, color.RGBA, color.RGBA)
}
type createimages interface {
	Createimage(encodeInformation []byte) (string, error)
}
type setImage interface {
	SetImage(encodeInformation []byte, color1 color.RGBA, color2 color.RGBA, color3 color.RGBA) *image.RGBA
}
type saveImage interface {
	SaveImage(path string, img *image.RGBA) error
}

type BuildImage struct {
	selectColor
	createimages
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

func (B *BuildImage) BuildAndSaveImage(encodeInformation []byte) error {
	imaProses1 := newBuildImage()
	path, err := imaProses1.Createimage(encodeInformation)
	if err != nil {
		return err
	}

	color1, color2, color3 := imaProses1.SelectColor(encodeInformation[10])
	img := imaProses1.SetImage(encodeInformation, color1, color2, color3)

	err = imaProses1.SaveImage(path, img)
	if err != nil {
		return err
	}

	return nil
}
