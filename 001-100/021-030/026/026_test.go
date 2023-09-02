package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb026(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo/hchw96FKnZOUdw20yImNF34iQ==",
		calc,
		1000,
	); err != nil {
		t.Errorf(err.Error())
	}
}
