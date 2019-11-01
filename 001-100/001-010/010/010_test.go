package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb010(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpKMed5ZJnZYDoLh5wRh3Vla49VcKZXKLh/VlQ==", calc, 2000000); err != nil {
		t.Errorf(err.Error())
	}
}
