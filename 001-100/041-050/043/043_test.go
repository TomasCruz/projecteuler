package main

import (
	"crypto/cipher"
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb043(t *testing.T) {
	var result string
	var err error
	if result, err = calc(); err != nil {
		t.Errorf("Problem solution execution broke")
	}

	var gcm cipher.AEAD
	if gcm, err = projecteuler.GenerateGcm(); err != nil {
		t.Errorf(err.Error())
	}

	if result, err = projecteuler.EncryptString(result, gcm); err != nil {
		t.Errorf("Problem solution encryption failed")
	}

	exp := "cGFzc3BocmFzZXdoBpCIedpZLXBYDoB/G4w4f92JutU0uJOZkD1A"
	if result != exp {
		t.Errorf("Result incorrect")
	}
}
