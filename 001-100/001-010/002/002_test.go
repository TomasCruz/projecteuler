package main

import (
	"crypto/cipher"
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb002(t *testing.T) {
	var result string
	var err error
	if result, err = calc(int64(4000000)); err != nil {
		t.Errorf("Problem solution execution broke")
	}

	var gcm cipher.AEAD
	if gcm, err = projecteuler.GenerateGcm(); err != nil {
		t.Errorf(err.Error())
	}

	if result, err = projecteuler.EncryptString(result, gcm); err != nil {
		t.Errorf("Problem solution encryption failed")
	}

	exp := "cGFzc3BocmFzZXdoA5CPc9hZLHkX0KPIj/7e3zr+FPEeB4A="
	if result != exp {
		t.Errorf("Result incorrect")
	}
}
