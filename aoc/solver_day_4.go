package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve4_1(s string) string {
	a := strings.Split(s, "\n")
	a = a[:len(a)-1]
	sum := 0

	for i := range a {
		s := strings.Split(a[i], ",")

		x := strings.Split(s[0], "-")
		y := strings.Split(s[1], "-")

		x1, _ := strconv.Atoi(x[0])
		x2, _ := strconv.Atoi(x[1])
		y1, _ := strconv.Atoi(y[0])
		y2, _ := strconv.Atoi(y[1])

		if x1 >= y1 && x2 <= y2 || x1 <= y1 && x2 >= y2 {
			sum++
		}
	}

	return fmt.Sprint(sum)
}

type TwoPair struct {
	a, b *Pair
}

func (p *TwoPair) overlap() bool {
	if p.a.x <= p.b.y && p.b.x <= p.a.y {
		return true
	}
	if p.b.x <= p.a.y && p.a.x <= p.b.y {
		return true
	}
	return false
}

type Pair struct {
	x, y   int
	solved bool
}

func (p *Pair) contain(p2 Pair) bool {
	if p.x >= p2.x && p.y <= p2.y {
		if p.solved {
		} else {
			p.solved = true
			return true
		}
	}
	if p.x <= p2.x && p.y >= p2.y {
		if p2.solved {
		} else {
			p2.solved = true
			return true
		}
	}
	return false
}

func Solve4_2(s string) string {
	// a := strings.Split(s, "\n")
	a := strings.FieldsFunc(s, func(r rune) bool {
		return r == '\n'
	})
	// a = a[:len(a)-1]

	var ps = []*TwoPair{}
	for i := range a {

		s := strings.Split(a[i], ",")

		x := strings.Split(s[0], "-")
		y := strings.Split(s[1], "-")

		x1, _ := strconv.Atoi(x[0])
		x2, _ := strconv.Atoi(x[1])
		y1, _ := strconv.Atoi(y[0])
		y2, _ := strconv.Atoi(y[1])

		tp := &TwoPair{
			a: &Pair{x1, x2, false},
			b: &Pair{y1, y2, false},
		}
		ps = append(ps, tp)
	}

	sum := 0
	for i := range ps {
		if ps[i].overlap() {
			sum++
		}
	}

	return fmt.Sprint(sum)
}
