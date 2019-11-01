package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb004(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoDpaIdt9T75w2XnwUZ057NtSoFQSYZw==", calc); err != nil {
		t.Errorf(err.Error())
	}
}
