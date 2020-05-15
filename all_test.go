package easyaes

import (
	"encoding/base64"
	"log"
	"testing"
)

func TestEncrypt(t *testing.T) {
	iv := []byte("ivTest1234567890")
	key := []byte("1234567890123456")
	data := []byte("test")

	encrypted, err := EncryptAesCBCStaticIV(key, iv, data)
	if err != nil {
		t.Error(err)
	}

	log.Println(base64.RawStdEncoding.EncodeToString(encrypted))

	decrypted, err := DecryptAesCBCStaticIV(key, iv, encrypted)
	if err != nil {
		t.Error(err)
	}

	if string(decrypted) != "test" {
		t.Error("Expected 'test', got ", string(decrypted))
	}

}
