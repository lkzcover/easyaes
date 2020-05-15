package easyaes

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// EncryptAesCBCStaticIV function get the key, initialization vector, and data in []byte. Returns encrypted data []byte and error
func EncryptAesCBCStaticIV(key, iv, data []byte) ([]byte, error) {

	data = fill(data)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encryptData := make([]byte, len(data))

	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(encryptData, data)

	return encryptData, nil

}

// DecryptAesCBCStaticIV function get the key, initialization vector, and encrypted data in []byte. Returns decrypted data []byte and error
func DecryptAesCBCStaticIV(key, iv, data []byte) ([]byte, error) {

	if (len(data) % aes.BlockSize) != 0 {
		return nil, errors.New("incorrect data length")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decryptData := make([]byte, len(data))

	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(decryptData, data)

	return unfill(decryptData)
}
