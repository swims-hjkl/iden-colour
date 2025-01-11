package png


import (
	"errors"
	"github.com/swims-hjkl/iden-colour/utils"
)

type PNG struct {
	bytesContent []uint8
	fileName string
}

func (png PNG) validatePNG() error {
	var headerBytes = []byte{0X89, 0X50, 0X4E, 0X47, 0X0D, 0X0A, 0X1A, 0x0A}
	if png.bytesContent == nil {
		return errors.New("png content cannot be nil")
	}
	if len(png.bytesContent) < 8 || !utils.BytesEqual(headerBytes, png.bytesContent[0:8]) {
		return errors.New("png header is not valid")
	}
	return nil
}

func NewPNG(bytesContent []uint8, fileName string) (*PNG, error) {
	if (len(bytesContent) == 0) || (bytesContent == nil) {
		return nil, errors.New("bytesContent cannot be empty or nil")
	}
	if fileName == "" {
		return nil, errors.New("fileName cannot be empty")
	}
	png := &PNG{bytesContent, fileName} 
	if png.validatePNG() != nil {
		return nil, png.validatePNG()
	}
	return png, nil
}

