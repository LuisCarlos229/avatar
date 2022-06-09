package encoder

import (
	"fmt"
	"hash/fnv"
	"log"
)

type EncodeHash struct {
}

// EncodeInformation se encarga de transformar la informacion de entrada en un numero del tipo uint32
//y devolver un error en caso de que algo aya salido mal.
func (e *EncodeHash) EncodeInformation(strInformation string) (encodeInformation uint64, err error) {
	fmt.Println()
	log.Println("encoding start...")

	miHash := fnv.New64a()
	miHash.Write([]byte(strInformation))

	log.Print("encoding completed...\n\n")

	return miHash.Sum64(), nil
}
