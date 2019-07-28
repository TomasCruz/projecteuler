package main

import (
	"crypto/cipher"
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb015(t *testing.T) {
	var result string
	var err error
	if result, err = calc(20); err != nil {
		t.Errorf("Problem solution execution broke")
	}

	var gcm cipher.AEAD
	if gcm, err = projecteuler.GenerateGcm(); err != nil {
		t.Errorf(err.Error())
	}

	if result, err = projecteuler.EncryptString(result, gcm); err != nil {
		t.Errorf("Problem solution encryption failed")
	}

	exp := "cGFzc3BocmFzZXdoBpWJeNtcK3ZYD4Lj3c9IutobxuR8/AQ2pVTAGQ=="
	if result != exp {
		t.Errorf("Result incorrect")
	}
}
