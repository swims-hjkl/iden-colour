package png

import (
	"errors"
)

type PNG struct {
	bytesContent []byte
	fileName string
}

func NewPNG(bytesContent []byte, fileName string) (*PNG, error) {
	if (len(bytesContent) == 0) || (bytesContent == nil) {
		return nil, errors.New("bytesContent cannot be empty or nil")
	}
	if fileName == "" {
		return nil, errors.New("fileName cannot be empty")
	}
	return &PNG{bytesContent, fileName}, nil
}
