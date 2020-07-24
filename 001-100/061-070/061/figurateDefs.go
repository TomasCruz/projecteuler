package main

const (
	figurateTriangle   = 0x1
	figurateSquare     = 0x2
	figuratePentagonal = 0x4
	figurateHexagonal  = 0x8
	figurateHeptagonal = 0x10
	figurateOctagonal  = 0x20
)

type (
	figurateNumber struct {
		x, t int
	}

	figurateMap map[int]figurateNumber
)
