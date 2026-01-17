package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 102; Triangle Containment
Three distinct points are plotted at random on a Cartesian plane, for which -1000 <= x, y <= 1000, such that a triangle is formed.

Consider the following two triangles:
A(-340,495), B(-153,-910), C(835,-947)
X(-175,41), Y(-421,-714), Z(574,-645)

It can be verified that triangle ABC contains the origin, whereas triangle XYZ does not.
Using triangles.txt(right click and 'Save Link/Target As...'), a 27K text file containing the co-ordinates of one thousand "random" triangles,
find the number of triangles for which the interior contains the origin.

NOTE: The first two examples in the file represent the triangles in the example given above.
*/

func main() {
	var limit int

	if len(os.Args) > 1 {
		limit64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		limit = int(limit64)
	} else {
		limit = 1
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	var fileName string
	if limit == 2 {
		fileName = "example.txt"
	} else {
		fileName = "0102_triangles.txt"
	}

	var rowStrings []string
	if rowStrings, err = projecteuler.FileToStrings(fileName); err != nil {
		return
	}

	sum := 0
	for _, s := range rowStrings {
		coordsStr := strings.Split(s, ",")
		coords := make([]int, 6)

		for i := 0; i < 6; i++ {
			coords[i], err = strconv.Atoi(coordsStr[i])
			if err != nil {
				return
			}
		}

		a := point{x: coords[0], y: coords[1]}
		b := point{x: coords[2], y: coords[3]}
		c := point{x: coords[4], y: coords[5]}

		if triangleContainsOrigin(a, b, c) {
			sum++
		}
	}

	result = strconv.Itoa(sum)
	return
}

func triangleContainsOrigin(a, b, c point) bool {
	return triangleContainsPoint(a, b, c, point{})
}
