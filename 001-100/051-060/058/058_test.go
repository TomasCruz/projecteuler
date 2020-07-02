package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb058(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBZCMdN4jAri0SfGtZ3wDKlW0Hr1v", calc, 15000); err != nil {
		t.Errorf(err.Error())
	}
}
