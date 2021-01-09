package cryption_test

import (
	"github.com/Ogguz/passvault/cryption"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO naming
func TestEncrypt(t *testing.T) {
	var data = []byte("texttoaes")
	var pass = "parola"
	encryptedData := cryption.Encrypt(data,pass)

	assert.Equal(t, data, cryption.Decrypt(encryptedData,pass))

}

