package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb017(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBZePctv7nCdgpTVz057OGbw7Oo8n", calc); err != nil {
		t.Errorf(err.Error())
	}
}
