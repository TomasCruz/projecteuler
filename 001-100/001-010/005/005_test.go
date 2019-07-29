package main

import (
	"crypto/cipher"
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb005(t *testing.T) {
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

	exp := "cGFzc3BocmFzZXdoBZWMd9ZYK3JQ2WsS+RE0r9QsiWRuNSN0Sg=="
	if result != exp {
		t.Errorf("Result incorrect")
	}
}
