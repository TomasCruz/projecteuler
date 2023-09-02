package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb050(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/hYlTo7rfMfbmwQLKBsH/1pBmLHdlg==",
		calc,
		1000000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
