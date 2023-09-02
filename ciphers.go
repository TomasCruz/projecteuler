package projecteuler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

// GenerateGcm creates cipher key
func GenerateGcm() (gcm cipher.AEAD, err error) {
	passPhrase := os.Getenv("PROJECT_EULER_PSWD")
	pswd := []byte(passPhrase)
	md5Sum := md5.Sum(pswd)

	var block cipher.Block
	if block, err = aes.NewCipher(md5Sum[:]); err != nil {
		fmt.Println(err)
		return
	}

	if gcm, err = cipher.NewGCM(block); err != nil {
		fmt.Println(err)
		return
	}

	return
}

// EncryptString encrypts
func EncryptString(txt string, gcm cipher.AEAD) (encryptedText string, err error) {
	nonce := make([]byte, gcm.NonceSize())
	r := strings.NewReader("passphrasewhichneedstobe32bytes!")

	if _, err = io.ReadFull(r, nonce); err != nil {
		fmt.Println("ReadFull", err)
		return
	}

	encryptedText = base64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(txt), nil))
	return
}

// DecryptString decrypts
func DecryptString(encryptedText string, gcm cipher.AEAD) (text string, err error) {
	nonceSize := gcm.NonceSize()

	var encryptedBytes []byte
	if encryptedBytes, err = base64.StdEncoding.DecodeString(encryptedText); err != nil {
		fmt.Println(err)
		return
	}

	if len(encryptedBytes) < nonceSize {
		fmt.Println(err)
		return
	}

	nonce, encryptedBytes := encryptedBytes[:nonceSize], encryptedBytes[nonceSize:]

	var textBytes []byte
	if textBytes, err = gcm.Open(nil, nonce, encryptedBytes, nil); err != nil {
		fmt.Println(err)
		return
	}

	text = string(textBytes)
	return
}

// funcToTest is the type of 'calc' function to test
type funcToTest func(...interface{}) (string, error)

// FuncForTesting is generic integration test function
func FuncForTesting(expectedEncryptedString string, f funcToTest, args ...interface{}) (err error) {
	var result string
	if result, err = f(args...); err != nil {
		err = fmt.Errorf("solution execution broke: %v", err)
		return
	}

	var gcm cipher.AEAD
	if gcm, err = GenerateGcm(); err != nil {
		return
	}

	if result, err = EncryptString(result, gcm); err != nil {
		err = fmt.Errorf("solution encryption failed: %v", err)
		return
	}

	if result != expectedEncryptedString {
		err = fmt.Errorf("Result incorrect")
		return
	}

	return
}
