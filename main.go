package main

import (
	"adventofcode/solvers"
)

func main() {
	aoc := AdventOfCode{
		session: "",
		year:    2022,
		day:     4,
		level:   2,
		solver:  solvers.Solve4_2,
		submit:  true,
	}

	aoc.SolveTestAndThenReal("4")
	// aoc.Solve()
}
