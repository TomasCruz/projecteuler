package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb084(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RglSo2oX90PgayZK2UqN0Ou0bA=",
		calc,
		4,
	); err != nil {
		t.Errorf(err.Error())
	}
}
