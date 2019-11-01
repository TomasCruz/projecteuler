package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb011(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoAJaIcN9cKXBjspBQlHlOGFa+9fvcVoaS", calc, 20, 4); err != nil {
		t.Errorf(err.Error())
	}
}
