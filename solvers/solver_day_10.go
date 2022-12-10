package solvers

import (
	"fmt"
	"strconv"
	"strings"
)

type ins struct {
	v    int
	what string
}

func Solve10_1(s string) string {
	raw := strings.Split(s, "\n")

	// Parse inst
	in := []ins{}
	for i := range raw[:len(raw)-1] {
		req := strings.Split(raw[i], " ")
		if len(req) == 2 {
			v, _ := strconv.Atoi(req[1])
			in = append(in, ins{v: v, what: req[0]})
		} else {
			in = append(in, ins{v: 0, what: req[0]})
		}
	}

	next1 := 0
	next2 := 0
	sum := 1
	count := 0
	tot := 0
	for i := range in {
		if in[i].what == "addx" {
			for j := 0; j < 2; j++ {
				count++
				sum += next1
				next1 = next2
				if j == 0 {
					next2 = in[i].v
				} else {
					next2 = 0
				}
				fmt.Println(i, count, in[i], sum, sum*(count-1))
				if (count%40)-20 == 0 {
					fmt.Println("TOT", count, sum, sum*(count-1))
					tot += (sum * count)
				}
			}
		} else {
			count++
			sum += next1
			next1 = next2
			next2 = in[i].v
			if (count%40)-20 == 0 {
				fmt.Println("TOT", count, sum, sum*(count-1))
				tot += (sum * count)
			}
			fmt.Println(i, count, in[i], sum, sum*count)
		}
	}

	return fmt.Sprint(tot)
}

func Solve10_2(s string) string {
	raw := strings.Split(s, "\n")

	// Parse inst
	in := []ins{}
	for i := range raw[:len(raw)-1] {
		req := strings.Split(raw[i], " ")
		if len(req) == 2 {
			v, _ := strconv.Atoi(req[1])
			in = append(in, ins{v: v, what: req[0]})
		} else {
			in = append(in, ins{v: 0, what: req[0]})
		}
	}

	next1 := 0
	next2 := 0
	sum := 1
	count := 0
	draw := []string{}
	for i := range in {
		if in[i].what == "addx" {
			for j := 0; j < 2; j++ {
				count++
				sum += next1
				next1 = next2
				if j == 0 {
					next2 = in[i].v
				} else {
					next2 = 0
				}
				fmt.Println(i, count, in[i], sum, sum*count)
				if sum == (count%40)-2 || sum == (count%40)-1 || sum == (count%40) {
					draw = append(draw, "#")
				} else {
					draw = append(draw, ".")
				}
				fmt.Println(draw)
			}
		} else {
			count++
			sum += next1
			next1 = next2
			next2 = in[i].v
			if sum == (count%40)-2 || sum == (count%40)-1 || sum == (count%40)+1 {
				draw = append(draw, "#")
			} else {
				draw = append(draw, ".")
			}
			fmt.Println(i, count, in[i], sum)
			fmt.Println(draw)
		}
	}
	fmt.Println(len(draw))
	for j, v := range draw {
		if j%40 == 0 {
			fmt.Println()
		}
		fmt.Print(v)
	}
	fmt.Println()
	return fmt.Sprint("tot")
}
