package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 155; Counting Capacitor Circuits
An electric circuit uses exclusively identical capacitors of the same value C.

The capacitors can be connected in series or in parallel to form sub-units, which can then be connected in series or in parallel with other capacitors
or other sub-units to form larger sub-units, and so on up to a final circuit.

Using this simple procedure and up to n identical capacitors, we can make circuits having a range of different total capacitances.
For example, using up to n=3 capacitors of 60mF each, we can obtain the following 7 distinct total capacitance values:

If we denote by D(n) the number of distinct total capacitance values we can obtain when using up to n equal-valued capacitors
and the simple procedure described above, we have: D(1)=1, D(2)=3, D(3)=7, ...

Find D(18).

Reminder: When connecting capacitors C_1, C_2 etc in parallel, the total capacitance is C_T = C_1 + C_2 + ...,
whereas when connecting them in series, the overall capacitance is given by: \dfrac{1}{C_T} = \dfrac{1}{C_1} + \dfrac{1}{C_2} + ...
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
		limit = 18
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	levels := make([]map[projecteuler.Fraction]struct{}, limit+1)
	levels[1] = map[projecteuler.Fraction]struct{}{projecteuler.NewFraction(1, 1): {}}

	for i := 2; i <= limit; i++ {
		levels[i] = map[projecteuler.Fraction]struct{}{}

		half := (i + 1) / 2
		for upper := i - 1; upper >= half; upper-- {
			lower := i - upper
			for ku := range levels[upper] {
				for kl := range levels[lower] {
					fParallel := projecteuler.AddFractions(ku, kl)
					levels[i][fParallel] = struct{}{}
					fSerial := projecteuler.AddFractions(projecteuler.ReciprocalFraction(ku), projecteuler.ReciprocalFraction(kl))
					fSerial = projecteuler.ReciprocalFraction(fSerial)
					levels[i][fSerial] = struct{}{}
				}
			}
		}
	}

	overall := map[projecteuler.Fraction]struct{}{}
	for i := 1; i <= limit; i++ {
		for k := range levels[i] {
			overall[k] = struct{}{}
		}
	}

	result = strconv.Itoa(len(overall))
	return
}
