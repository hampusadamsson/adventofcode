package solvers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solve9_1(s string) string {
	hist := make(map[pos]int)
	hist[pos{0, 0}] = 1

	h := &pos{0, 0}
	t := &pos{0, 0}

	rows := strings.Split(s, "\n")
	for _, insraw := range rows[:len(rows)-1] {
		fmt.Println(insraw)
		ins := strings.Split(insraw, " ")
		dir := ins[0]
		post, _ := strconv.Atoi(ins[1])

		for i := 0; i < post; i++ {
			if dir == "R" {
				h.x++
			} else if dir == "L" {
				h.x--
			} else if dir == "U" {
				h.y++
			} else if dir == "D" {
				h.y--
			}

			// Move tail
			man_dist := int(math.Abs(float64(h.x-t.x))) + int(math.Abs(float64(h.y-t.y)))
			if man_dist <= 1 {
				//adjacent
			} else if int(math.Abs(float64(h.x-t.x))) == 1 && int(math.Abs(float64(h.y-t.y))) == 1 {
				// diag
			} else {
				// move
				if h.x > t.x {
					t.x++
				} else if h.x < t.x {
					t.x--
				}
				if h.y > t.y {
					t.y++
				} else if h.y < t.y {
					t.y--
				}
				hist[pos{x: t.x, y: t.y}] = 1
			}
			fmt.Println(h, t)
		}

	}

	return fmt.Sprint(len(hist))
}

func Solve9_2(s string) string {
	hist := make(map[pos]int)
	hist[pos{0, 0}] = 1

	h := &pos{0, 0}
	ts := []*pos{}

	ts = append(ts, h)
	for i := 0; i < 9; i++ {
		ts = append(ts, &pos{0, 0})
	}

	rows := strings.Split(s, "\n")
	for _, insraw := range rows[:len(rows)-1] {
		fmt.Println(insraw)
		ins := strings.Split(insraw, " ")
		dir := ins[0]
		post, _ := strconv.Atoi(ins[1])

		for i := 0; i < post; i++ {
			if dir == "R" {
				h.x++
			} else if dir == "L" {
				h.x--
			} else if dir == "U" {
				h.y++
			} else if dir == "D" {
				h.y--
			}

			// Move tail
			for i := 1; i < len(ts); i++ {
				man_dist := int(math.Abs(float64(ts[i-1].x-ts[i].x))) + int(math.Abs(float64(ts[i-1].y-ts[i].y)))
				if man_dist <= 1 {
					//adjacent
				} else if int(math.Abs(float64(ts[i-1].x-ts[i].x))) == 1 && int(math.Abs(float64(ts[i-1].y-ts[i].y))) == 1 {
					// diag
				} else {
					// move
					if ts[i-1].x > ts[i].x {
						ts[i].x++
					} else if ts[i-1].x < ts[i].x {
						ts[i].x--
					}
					if ts[i-1].y > ts[i].y {
						ts[i].y++
					} else if ts[i-1].y < ts[i].y {
						ts[i].y--
					}

					if i == 9 {
						nx := ts[i].x
						ny := ts[i].y
						hist[pos{x: nx, y: ny}] = 1
					}
				}
				// fmt.Println(h, t)
			}

			fmt.Println("---")
			// fmt.Print(h)
			for _, v := range ts {
				fmt.Print(v)
			}
			fmt.Println()
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					found := false
					for pp, _ := range ts {
						if !found && ts[pp].x == j && ts[pp].y == i {
							ci := fmt.Sprint(pp)
							if ci == "0" {
								ci = "H"
							}
							fmt.Print(ci + " ")
							found = true
						}
					}
					if !found {
						fmt.Print(". ")
					}
				}
				fmt.Println()
			}

		}

	}

	return fmt.Sprint(len(hist))
}
