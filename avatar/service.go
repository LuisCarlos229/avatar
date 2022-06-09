package avatar

import (
	"github.com/LuisCarlos229/avatar/encoder"
	"github.com/LuisCarlos229/avatar/images"
)

type cryptoEncoder interface {
	EncodeInformation(strInformation string) (encodeInformation uint64, err error)
}

type imageGenerator interface {
	BuildAndSaveImage(encodeInformation uint64) error
}

//Service se encarga de resibir la Informacion y darla a las funciones que
//las trabajaran.
type Service struct {
	encoder   cryptoEncoder
	generator imageGenerator
}

// Information resibe el correo del usuario como string.
type Information struct {
	MAIL string
}

// NewService crea un nuevo servicio para la carga de la Informacion.
func NewService() *Service {
	return &Service{
		encoder:   &encoder.EncodeHash{},
		generator: &images.BuildImage{},
	}
}

// GenerateAndSaveAvatar se encarga de la creacion, encode y guardado del Avatar.
func (s *Service) GenerateAndSaveAvatar(information Information) error {
	encodeInformation, err := s.encoder.EncodeInformation(information.MAIL)
	if err != nil {
		return err
	}
	err = s.generator.BuildAndSaveImage(encodeInformation)
	return err
}
