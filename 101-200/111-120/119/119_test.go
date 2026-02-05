package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb119(t *testing.T) {
	if err := projecteuler.FuncForTesting(
		"cGFzc3BocmFzZXdo9RsqSY7vEK/wQ9fXy7RGhLkfBRZHkvuZTf/n9BKpxQ==",
		calc,
		30,
	); err != nil {
		t.Errorf(err.Error())
	}
}
