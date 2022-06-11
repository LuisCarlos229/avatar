package createImage

import (
	"log"
	"os"
	"strconv"
)

type Createimage struct{}

// Createimage se encarga de la generacion del archivo png vacio para su posterior seteo, y devuelve el path del archivo
// y un error en caso de que algo aya salido mal.
func (C *Createimage) Createimage(encodeInformation []byte) (path string, err error) {

	path = ""
	for i, byte := range encodeInformation {
		path += strconv.Itoa(int(byte))
		if i > 8 {
			break
		}
	}

	path += ".png"

	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		log.Println("El directorio o archivo no puede ser creado.")
		panic(err)
	}

	return path, nil
}
