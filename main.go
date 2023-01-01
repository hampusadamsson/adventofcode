package main

import "adventofcode/solvers"

func main() {
	aoc := AdventOfCode{
		session: "",
		year:    2022,
		day:     14,
		level:   2,
		solver:  solvers.Solve14_2,
		submit:  true,
	}

	aoc.SolveTestAndThenReal("93")
}
