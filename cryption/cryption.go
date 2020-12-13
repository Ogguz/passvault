package cryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// createHash take a passphrase or any string, hash it, then return the hash as a hexadecimal value.
func createHash(key string) string {
	hasher := md5.New()
    hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte  {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	cipherText := gcm.Seal(nonce, nonce, data, nil)

	return cipherText

}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func EncryptToFile(filename string, data []byte, passphrase string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("File create fail for encryption, %s", err)
	}
	defer f.Close()
	f.Write(encrypt(data,passphrase))
}

func DecryptFromFile(filename string, passphrase string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Reading file ")
	}
	return decrypt(data,passphrase)
}