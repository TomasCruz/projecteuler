package main

type point struct {
	x, y int
}

func dotProduct(p1, p2 point) int {
	return p1.x*p2.x + p1.y*p2.y
}

func triangleContainsPoint(a, b, c, p point) bool {
	// using barycentric coordinates
	v0 := point{x: c.x - a.x, y: c.y - a.y}
	v1 := point{x: b.x - a.x, y: b.y - a.y}
	v2 := point{x: p.x - a.x, y: p.y - a.y}

	d00 := dotProduct(v0, v0)
	d01 := dotProduct(v0, v1)
	d11 := dotProduct(v1, v1)
	d02 := dotProduct(v0, v2)
	d12 := dotProduct(v1, v2)

	denom := d00*d11 - d01*d01
	l2 := float64(d11*d02-d01*d12) / float64(denom)
	l3 := float64(d00*d12-d01*d02) / float64(denom)
	l1 := 1 - l2 - l3

	if l1 >= 0.0 && l1 <= 1.0 && l2 >= 0.0 && l2 <= 1.0 && l3 >= 0.0 && l3 <= 1.0 {
		return true
	}

	return false
}
