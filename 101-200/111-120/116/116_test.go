package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb116(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9R8mQYnvEKf5Q9gM/CIqmz9ln7a8pXX8s5Kr",
		calc,
		50,
	); err != nil {
		t.Errorf(err.Error())
	}
}
