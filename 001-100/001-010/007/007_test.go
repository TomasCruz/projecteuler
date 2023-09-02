package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb007(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h8mT4/pZuo9Xb5HE1lXc5MdNTvK9w==",
		calc,
		10001,
	); err != nil {
		t.Errorf(err.Error())
	}
}
