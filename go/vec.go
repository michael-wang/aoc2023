package main

import (
	"math"
	"strconv"
)

type vec2 struct {
	X, Y int
}

func (v *vec2) Parse(str string) {
	bb := []byte(str)
	// Find next non-number
	s := 0
	t := nextNoneNumber(bb, 1)
	x, err := strconv.Atoi(string(bb[s:t]))
	if err != nil {
		panic(err)
	}
	v.X = x

	s = nextNumber(bb, t)
	t = nextNoneNumber(bb, s+1)
	y, err := strconv.Atoi(string(bb[s:t]))
	if err != nil {
		panic(err)
	}
	v.Y = y
}

func (v *vec2) lowerBound(x, y int) {
	if v.X < x {
		v.X = x
	}
	if v.Y < y {
		v.Y = y
	}
}

func (v *vec2) upperBound(x, y int) {
	if x < v.X {
		v.X = x
	}
	if y < v.Y {
		v.Y = y
	}
}

type path []vec2

func (p *path) Build(vv []vec2) {
	if len(vv) == 0 {
		return
	}

	for i := 1; i < len(vv); i++ {
		p.connect(vv[i-1], vv[i])
	}

	*p = append(*p, vv[len(vv)-1])
}

// Add vec2 between [a, b) to path.
// Assume the line between a and b is vertical or horizontal.
func (p *path) connect(a, b vec2) {
	if a.X == b.X {
		if a.Y < b.Y {
			for y := a.Y; y < b.Y; y++ {
				*p = append(*p, vec2{a.X, y})
			}
		} else {
			for y := a.Y; y > b.Y; y-- {
				*p = append(*p, vec2{a.X, y})
			}
		}
	} else {
		// a.Y == b.Y
		if a.X < b.X {
			for x := a.X; x < b.X; x++ {
				*p = append(*p, vec2{x, a.Y})
			}
		} else {
			for x := a.X; x > b.X; x-- {
				*p = append(*p, vec2{x, a.Y})
			}
		}
	}
}

func (p path) max() (x, y int) {
	x, y = math.MinInt, math.MinInt
	for _, v := range p {
		if v.X > x {
			x = v.X
		}
		if v.Y > y {
			y = v.Y
		}
	}
	return
}

func (p path) min() (x, y int) {
	x, y = math.MaxInt, math.MaxInt
	for _, v := range p {
		if v.X < x {
			x = v.X
		}
		if v.Y < y {
			y = v.Y
		}
	}
	return
}

type paths []path

func (pp paths) max() vec2 {
	x, y := pp[0].max()
	for i := 1; i < len(pp); i++ {
		_x, _y := pp[i].max()
		if _x > x {
			x = _x
		}
		if _y > y {
			y = _y
		}
	}
	return vec2{x, y}
}

func (pp paths) min() vec2 {
	x, y := pp[0].min()
	for i := 1; i < len(pp); i++ {
		_x, _y := pp[i].min()
		if _x < x {
			x = _x
		}
		if _y < y {
			y = _y
		}
	}
	return vec2{x, y}
}
