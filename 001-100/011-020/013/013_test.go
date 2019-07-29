package main

import (
	"crypto/cipher"
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb013(t *testing.T) {
	var result string
	var err error
	if result, err = calc(10); err != nil {
		t.Errorf("Problem solution execution broke")
	}

	var gcm cipher.AEAD
	if gcm, err = projecteuler.GenerateGcm(); err != nil {
		t.Errorf(err.Error())
	}

	if result, err = projecteuler.EncryptString(result, gcm); err != nil {
		t.Errorf("Problem solution encryption failed")
	}

	exp := "cGFzc3BocmFzZXdoApONd9xdKHZTB7M6S0w2rb56D0GQ+BIRRu4="
	if result != exp {
		t.Errorf("Result incorrect")
	}
}
