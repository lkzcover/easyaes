package easyaes

import (
	"bytes"
	"crypto/aes"
	"errors"
)

func fill(data []byte) []byte {
	filling := aes.BlockSize - len(data)%aes.BlockSize
	filltext := bytes.Repeat([]byte{byte(filling)}, filling)
	return append(data, filltext...)
}

func unfill(data []byte) ([]byte, error) {
	length := len(data)
	unfilling := int(data[length-1])

	if unfilling > length {
		return nil, errors.New("error: check encryption key")
	}

	return data[:(length - unfilling)], nil
}
