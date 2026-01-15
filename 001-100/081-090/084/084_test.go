package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb084(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9h8jTYnuWkCHozyPedFfK3CTPVEjtQ==",
		calc,
		4,
	); err != nil {
		t.Errorf(err.Error())
	}
}
