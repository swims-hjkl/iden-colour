package png


import (
	"errors"
	"github.com/swims-hjkl/iden-colour/utils"
)

type PNG struct {
	bytesContent []uint8
	fileName string
}

func NewPNG(bytesContent []uint8, fileName string) (*PNG, error) {
	if (len(bytesContent) == 0) || (bytesContent == nil) {
		return nil, errors.New("bytesContent cannot be empty or nil")
	}
	if fileName == "" {
		return nil, errors.New("fileName cannot be empty")
	}
	if validatePng(bytesContent) != nil{
		return nil, validatePng(bytesContent)
	}
	return &PNG{bytesContent, fileName}, nil
}

func validatePng(pngContent []uint8) error {
	var headerBytes = []byte{0X89, 0X50, 0X4E, 0X47, 0X0D, 0X0A, 0X1A, 0x0A}
	if pngContent == nil {
		return errors.New("png content cannot be nil")
	}
	if len(pngContent) < 8 || !utils.BytesEqual(headerBytes, pngContent[0:8]) {
		return errors.New("png header is not valid")
	}
	return nil
}
