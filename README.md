# easyaes
simple example of AES data encryption

### func DecryptAesCBCStaticIV
    func DecryptAesCBCStaticIV(key, iv, data []byte) ([]byte, error)
DecryptAesCBCStaticIV function get the key, initialization vector, and encrypted data in []byte. Returns decrypted data []byte and error

### func EncryptAesCBCStaticIV
    func EncryptAesCBCStaticIV(key, iv, data []byte) ([]byte, error)
EncryptAesCBCStaticIV function get the key, initialization vector, and data in []byte. Returns encrypted data []byte and error