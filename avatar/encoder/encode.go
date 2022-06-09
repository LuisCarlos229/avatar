package encoder

import (
	"crypto/sha256"
	"fmt"
	"log"
)

type EncodeHash struct {
}

// EncodeInformation se encarga de transformar la informacion de entrada en un numero del tipo uint32
//y devolver un error en caso de que algo aya salido mal.
func (e *EncodeHash) EncodeInformation(strInformation string) (encodeInformation []byte, err error) {
	fmt.Println()
	log.Println("encoding start...")

	miHash := sha256.New()
	miHash.Write([]byte(strInformation))

	log.Print("encoding completed...\n\n")

	return miHash.Sum([]byte{20}), nil
}
