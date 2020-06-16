package main

import (
	"testing"

	"github.com/TomasCruz/projecteuler"
)

func TestProb045(t *testing.T) {
	if err := projecteuler.FuncForTesting("cGFzc3BocmFzZXdoBpONc9hdKHxQAvYJcQ4OPUOwJbsJwK22EBM=", calc, 1000000); err != nil {
		t.Errorf(err.Error())
	}
}
