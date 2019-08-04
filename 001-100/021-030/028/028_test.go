package main

import (
	"crypto/cipher"
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb028(t *testing.T) {
	var result string
	var err error
	if result, err = calc(1001); err != nil {
		t.Errorf("Problem solution execution broke")
	}

	var gcm cipher.AEAD
	if gcm, err = projecteuler.GenerateGcm(); err != nil {
		t.Errorf(err.Error())
	}

	if result, err = projecteuler.EncryptString(result, gcm); err != nil {
		t.Errorf("Problem solution encryption failed")
	}

	exp := "cGFzc3BocmFzZXdoAZCHcdhbLnRRQg5f3LWBP86eaUlsfg9wHA=="
	if result != exp {
		t.Errorf("Result incorrect")
	}
}
