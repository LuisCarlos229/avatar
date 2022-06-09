package createImage

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Createimage struct{}

func (C *Createimage) Createimage(encodeInformation uint64) (string, error) {
	fmt.Println()
	log.Println("creating image in proces...")

	path := strconv.Itoa(int(encodeInformation / 2))
	path += ".png"

	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return "", err
	}

	log.Println("creating image ready...")

	return path, nil
}
