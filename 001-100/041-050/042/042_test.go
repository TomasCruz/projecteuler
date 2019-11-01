package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb042(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpCM+SoYuL1GMAHu59jj+BF6YQ==", calc); err != nil {
		t.Errorf(err.Error())
	}
}
