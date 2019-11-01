package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb030(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoA5KNeNxTz99VuoEJp8+1xzn0xp+Grw==", calc, 5); err != nil {
		t.Errorf(err.Error())
	}
}
