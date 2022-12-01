package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solve2019_1_1(s string) string {
	a := strings.Split(s, "\n")
	a = a[:len(a)-1]
	arr := make([]int, 0)

	for i := range a {
		v, _ := strconv.Atoi(a[i])
		arr = append(arr, (v/3)-2)
	}

	max := 0
	for i := range arr {
		max += arr[i]
	}

	return fmt.Sprintln(max)
}
