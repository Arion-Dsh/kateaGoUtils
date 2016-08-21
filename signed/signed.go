package signed

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

var secretKey string = fmt.Sprintf("%s", sha256.Sum256([]byte("secret key")))

// setSecretKey sdefualt is "secret key"
func SetSecretKey(key string) {
	secretKey = fmt.Sprintf("%s", sha256.Sum256([]byte(key)))
}

func Decode(t string) string {
	ciphertext, _ := hex.DecodeString(t)

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)
	fmt.Printf("%s", ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

func Encode(t string) string {
	plaintext := []byte(t)

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return fmt.Sprintf("%x", ciphertext)
}
