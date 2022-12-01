package main

func main() {
	aoc := AdventOfCode{
		session: "",
		year:    2022,
		day:     2,
		level:   1,
		solver:  solve2_1,
		submit:  true,
	}

	aoc.Solve()
}
