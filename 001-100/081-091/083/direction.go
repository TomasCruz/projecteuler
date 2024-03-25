package main

type direction int

const (
	left  direction = 0
	down  direction = 1
	right direction = 2
	up    direction = 3
)

func (d direction) opposite() direction {
	switch d {
	case left:
		return right
	case down:
		return up
	case right:
		return left
	}

	return down
}

func (d direction) others() []direction {
	dirs := []direction{}

	if d != left {
		dirs = append(dirs, left)
	}
	if d != down {
		dirs = append(dirs, down)
	}
	if d != right {
		dirs = append(dirs, right)
	}
	if d != up {
		dirs = append(dirs, up)
	}

	return dirs
}
