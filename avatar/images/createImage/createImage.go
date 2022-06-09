package createImage

import (
	"log"
	"os"
	"strconv"
)

type Createimage struct{}

func (C *Createimage) Createimage(encodeInformation []byte) (string, error) {

	log.Println("creating image in proces...")

	path := ""
	for i, byte := range encodeInformation {
		path += strconv.Itoa(int(byte))
		if i > 8 {
			break
		}
	}

	//path := strconv.Itoa(int(encodeInformation / 2))
	path += ".png"

	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return "", err
	}

	log.Println("creating image ready...")

	return path, nil
}
