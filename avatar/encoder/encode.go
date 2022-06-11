package encoder

import (
	"crypto/sha256"
)

type customEncode interface {
	// EncodeInformation se encarga de transformar la informacion de entrada en una cadena de byte
	//y devolver un error en caso de que algo aya salido mal.
	EncodeInformation(strInformation string) (encodeInformation []byte, err error)
}

// EncodeHash es la estructura utilizada para realizar un encode
// apartir del encriptado por sha256.
type EncodeHash struct {
}

type customEncodeHash struct {
	customEncode
}

// NewEncodeHash es la funcion encargada de generar una nueva orden
// para la creacion de un nuevo hash.
func NewEncodeHash() *customEncodeHash {
	return &customEncodeHash{
		&EncodeHash{}}
}

// EncodeInformation se encarga de transformar la informacion de entrada en una cadena de byte
//y devolver un error en caso de que algo aya salido mal.
func (e *EncodeHash) EncodeInformation(strInformation string) (encodeInformation []byte, err error) {

	miHash := sha256.New()
	miHash.Write([]byte(strInformation))

	return miHash.Sum([]byte{20}), nil
}
