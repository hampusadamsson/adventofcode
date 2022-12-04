package main

import (
	"adventofcode/aoc"
)

func main() {
	aoc := AdventOfCode{
		session: "",
		year:    2022,
		day:     4,
		level:   2,
		solver:  aoc.Solve4_2,
		submit:  true,
	}

	aoc.SolveTestAndThenReal("4")
	// aoc.Solve()
}
