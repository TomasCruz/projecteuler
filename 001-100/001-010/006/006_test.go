package main

import (
	"crypto/cipher"
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb006(t *testing.T) {
	var result string
	var err error
	if result, err = calc(100); err != nil {
		t.Errorf("Problem solution execution broke")
	}

	var gcm cipher.AEAD
	if gcm, err = projecteuler.GenerateGcm(); err != nil {
		t.Errorf(err.Error())
	}

	if result, err = projecteuler.EncryptString(result, gcm); err != nil {
		t.Errorf("Problem solution encryption failed")
	}

	exp := "cGFzc3BocmFzZXdoBZOPdttbK3R30e6ikjLJ0Op4gq1TrfC8"
	if result != exp {
		t.Errorf("Result incorrect")
	}
}
