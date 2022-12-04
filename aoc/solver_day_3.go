package aoc

import (
	"fmt"
	"strings"
)

func Solve3_1(s string) string {
	a := strings.Split(s, "\n")
	sum := 0

	for i := range a {
		x := a[i][:len(a[i])/2]
		y := a[i][len(a[i])/2:]

		fmt.Println(x, y)
		for v := range x {
			if strings.Contains(y, string(x[v])) {
				fmt.Println(string(x[v]))
				run := []rune(string(x[v]))
				if int(run[0]) > 96 {
					vv := int(run[0]) - (50 + 38 + 8)
					fmt.Println("RUN:", vv)
					sum += vv
				} else {
					vv := int(run[0]) - (38)
					fmt.Println("RUN:", vv)
					sum += vv
				}
				break
			}
		}
	}
	return fmt.Sprint(sum)
}

func Solve3_2(s string) string {
	b := strings.Split(s, "\n")
	sum := 0

	fmt.Println(len(b))
	var appart = [][]string{}

	for i := 3; i <= len(b); i += 3 {
		appart = append(appart, b[i-3:i])
		fmt.Println()
	}

	for j := range appart {
		a := appart[j]
		fmt.Println("New appart")
		for v := range a[0] {

			if strings.Contains(a[1], string(a[0][v])) && strings.Contains(a[2], string(a[0][v])) {
				fmt.Println(string(a[0][v]))
				run := []rune(string(a[0][v]))
				if int(run[0]) > 96 {
					vv := int(run[0]) - (50 + 38 + 8)
					fmt.Println("RUN:", vv)
					sum += vv
				} else {
					vv := int(run[0]) - (38)
					fmt.Println("RUN:", vv)
					sum += vv
				}
				break
			}
		}

	}
	return fmt.Sprint(sum)
}
