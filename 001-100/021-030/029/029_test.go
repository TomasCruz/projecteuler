package main

import (
	"crypto/cipher"
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb029(t *testing.T) {
	var result string
	var err error
	if result, err = calc(101); err != nil {
		t.Errorf("Problem solution execution broke")
	}

	var gcm cipher.AEAD
	if gcm, err = projecteuler.GenerateGcm(); err != nil {
		t.Errorf(err.Error())
	}

	if result, err = projecteuler.EncryptString(result, gcm); err != nil {
		t.Errorf("Problem solution encryption failed")
	}

	exp := "cGFzc3BocmFzZXdoDpeGc8c7r9+90zh3T9l4gL+To6g="
	if result != exp {
		t.Errorf("Result incorrect")
	}
}