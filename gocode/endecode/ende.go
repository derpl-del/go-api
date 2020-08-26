package endecode

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var key = "a very very very very secret key"

//GenerateEn func
func GenerateEn(input string) string {
	ciphertext := Encrypt(key, input)
	fmt.Println(ciphertext)
	return ciphertext
}

//GenerateDe func
func GenerateDe(input string) string {
	result := Decrypt(key, input)
	fmt.Println(result)
	return result
}

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

//Encrypt func
func Encrypt(key, text string) string {
	fmt.Println(text)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return encodeBase64(ciphertext)
}

//Decrypt func
func Decrypt(key, text string) string {
	fmt.Println(text)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	ciphertext := decodeBase64(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext)
}

//encodeBase64 func
func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

//decodeBase64 func
func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}
