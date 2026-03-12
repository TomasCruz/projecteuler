package main

import (
	"math"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 144; Laser Beam Reflections
In laser physics, a "white cell" is a mirror system that acts as a delay line for the laser beam. The beam enters the cell,
bounces around on the mirrors, and eventually works its way back out.

The specific white cell we will be considering is an ellipse with the equation 4x^2 + y^2 = 100.

The section corresponding to -0.01 <= x <= +0.01 at the top is missing, allowing the light to enter and exit through the hole.

The light beam in this problem starts at the point (0.0,10.1) just outside the white cell, and the beam first impacts the mirror at (1.4,-9.6).

Each time the laser beam hits the surface of the ellipse, it follows the usual law of reflection "angle of incidence equals angle of reflection."
That is, both the incident and reflected beams make the same angle with the normal line at the point of incidence.

In the figure on the left, the red line shows the first two points of contact between the laser beam and the wall of the white cell;
the blue line shows the line tangent to the ellipse at the point of incidence of the first bounce.

The slope m of the tangent line at any point (x,y) of the given ellipse is: m = -4x/y.
The normal line is perpendicular to this tangent line at the point of incidence.

How many times does the beam hit the internal surface of the white cell before exiting?
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	b := []point{{x: 0, y: 10.1}, {x: 1.4, y: -9.6}}

	yExit := math.Sqrt(99.9996)

	i := 1
	for ; ; i++ {
		xm := (4*b[i].y*b[i-1].y + 16*b[i].x*b[i-1].x - 3*b[i].y*b[i].y) * b[i].x / (b[i].y*b[i].y + 16*b[i].x*b[i].x)
		ym := (b[i].y*xm + 3*b[i].x*b[i].y) / (4 * b[i].x)

		aPrimeX := 2*xm - b[i-1].x
		aPrimeY := 2*ym - b[i-1].y

		nextRayCoef := (aPrimeY - b[i].y) / (aPrimeX - b[i].x)
		nextRayN := b[i].y - nextRayCoef*b[i].x

		xExit := (yExit - nextRayN) / nextRayCoef
		if -0.01 <= xExit && xExit <= 0.01 {
			break
		}

		// B_nA'
		nextRay := point{x: aPrimeX - b[i].x, y: aPrimeY - b[i].y}

		// t = -(8*B_n.x*B_nA'.x + 2*B_n.y*B_nA'.y) / (B_nA'.y^2 + 4B_nA'.x^2)
		t := -(8*b[i].x*nextRay.x + 2*b[i].y*nextRay.y) / (nextRay.y*nextRay.y + 4*nextRay.x*nextRay.x)

		// B_n + t*B_nA' = B_n+1
		nextX := b[i].x + t*nextRay.x
		nextY := b[i].y + t*nextRay.y

		b = append(b, point{x: nextX, y: nextY})
	}

	result = strconv.Itoa(i)
	return
}

type point struct {
	x, y float64
}

/*
	4x^2 + y^2 = 100
	x^2/25 + y^2/100 = 1
	a = 5, b = 10

	B0(0, 10.1); B1(1.4,-9.6)
	Calculating B_n+1:

	Line passing through B_n-1 parallel to B_n's tangent:
		y_n-1 = -4x_n/y_n * x_n-1 + n => n = y_n-1 + 4*x_n*x_n-1/y_n
		y = -4*x_n/y_n * x + y_n-1 + 4*x_n*x_n-1/y_n
	Normal through B_n:
		y_n = y_n/(4*x_n) * x_n + n => n = 3*y_n/4
		y = y_n/(4*x_n) * x + 3*y_n/4

	Point M (intersection of these two lines):
		(-4*x_n*x + y_n*y_n-1 + 4*x_n*x_n-1)/y_n = (y_n*x + 3*x_n*y_n)/(4*x_n) / * 4*x_n*y_n
		-16*x_n^2*x + 4*x_n*y_n*y_n-1 + 16*x_n^2*x_n-1 = y_n^2*x + 3*x_n*y_n^2
		4*x_n*y_n*y_n-1 + 16*x_n^2*x_n-1 - 3*x_n*y_n^2 = (y_n^2 + 16*x_n^2)*x
		x_n*(4*y_n*y_n-1 + 16*x_n*x_n-1 - 3*y_n^2) = (y_n^2 + 16*x_n^2)*x
		x = (4*y_n*y_n-1 + 16*x_n*x_n-1 - 3*y_n^2)*x_n/(y_n^2 + 16*x_n^2)

	Point A' = 2*M - A
	k_nextRay (A'.y - B_n.y)/(A'.x - B_n.x)
	B_n.y = k_nextRay*B_n.x + n => n = B_n.y - k_nextRay*B_n.x
	Next ray:
		y = k_nextRay * x + B_n.y - k_nextRay*B_n.x

	4*0.0001 + y^2 = 100
	y^2 = 99.9996 => y = sqrt(99.9996) = yExit
	(yExit - nextRay_n)/k_nextRay = xExit

	Intersection of next ray with the ellipse:
		(k_nextRay * x + nextRay_n)^2 + 4x^2 = 100
		(4 + k_nextRay^2)*x^2 + 2*k_nextRay*nextRay_n*x + nextRay_n^2 - 100 = 0

		B_n + t*B_nA' = B_n+1
		4*(B_n.x + t*(A'.x - B_n.x))^2 + (B_n.y + t*(A'.y - B_n.y))^2 = 100
		4*(B_n.x^2 + t*2*B_n.x*(A'.x - B_n.x) + t^2*(A'.x - B_n.x)^2) + B_n.y^2 + t*2*B_n.y*(A'.y - B_n.y) + t^2*(A'.y - B_n.y)^2 = 100
		t^2*((A'.y - B_n.y)^2 + 4*(A'.x - B_n.x)^2) + t*(8*B_n.x*(A'.x - B_n.x) + 2*B_n.y*(A'.y - B_n.y)) + 4*B_n.x^2 + B_n.y^2 - 100 = 0
		t^2*((A'.y - B_n.y)^2 + 4*(A'.x - B_n.x)^2) + t*(8*B_n.x*(A'.x - B_n.x) + 2*B_n.y*(A'.y - B_n.y)) = 0
		t*((A'.y - B_n.y)^2 + 4*(A'.x - B_n.x)^2) + (8*B_n.x*(A'.x - B_n.x) + 2*B_n.y*(A'.y - B_n.y)) = 0
		t*(B_nA'.y^2 + 4B_nA'.x^2) + (8*B_n.x*B_nA'.x + 2*B_n.y*B_nA'.y) = 0
		t = -(8*B_n.x*B_nA'.x + 2*B_n.y*B_nA'.y) / (B_nA'.y^2 + 4B_nA'.x^2)
*/
