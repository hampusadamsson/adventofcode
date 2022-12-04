package solvers

import (
	"fmt"
	"strings"
)

func Solve2_1(s string) string {
	a := strings.Split(s, "\n")
	a = a[:len(a)-1] // trainling new line
	arr := make([]int, 0)

	for i := range a {
		b := strings.Split(a[i], " ")
		sum := 0
		// strat
		if b[0] == "A" && b[1] == "X" {
			sum = 3 + 1
		}
		if b[0] == "B" && b[1] == "Y" {
			sum = 3 + 2
		}
		if b[0] == "C" && b[1] == "Z" {
			sum = 3 + 3
		}

		if b[0] == "A" && b[1] == "Y" {
			sum = 6 + 2
		}
		if b[0] == "A" && b[1] == "Z" {
			sum = 0 + 3
		}

		if b[0] == "B" && b[1] == "X" {
			sum = 0 + 1
		}
		if b[0] == "B" && b[1] == "Z" {
			sum = 6 + 3
		}

		if b[0] == "C" && b[1] == "X" {
			sum = 6 + 1
		}
		if b[0] == "C" && b[1] == "Y" {
			sum = 0 + 2
		}
		arr = append(arr, sum)
	}
	max := 0
	for i := range arr {
		max += arr[i]
	}
	return fmt.Sprintln(max)
}

func Solve2_2(s string) string {
	a := strings.Split(s, "\n")
	a = a[:len(a)-1] // trainling new line
	arr := make([]int, 0)

	for i := range a {
		b := strings.Split(a[i], " ")
		sum := 0
		// lose
		if b[1] == "X" {
			if b[0] == "A" {
				sum = 0 + 3
			}
			if b[0] == "B" {
				sum = 0 + 1
			}
			if b[0] == "C" {
				sum = 0 + 2
			}
		}
		// equal
		if b[1] == "Y" {
			if b[0] == "A" {
				sum = 3 + 1
			}
			if b[0] == "B" {
				sum = 3 + 2
			}
			if b[0] == "C" {
				sum = 3 + 3
			}
		}
		// win
		if b[1] == "Z" {
			if b[0] == "A" {
				sum = 6 + 2
			}
			if b[0] == "B" {
				sum = 6 + 3
			}
			if b[0] == "C" {
				sum = 6 + 1
			}
		}
		arr = append(arr, sum)
	}
	max := 0
	for i := range arr {
		max += arr[i]
	}
	return fmt.Sprintln(max)
}
