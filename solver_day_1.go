package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func solve1_1(s string) string {
	a := strings.Split(s, "\n\n")
	arr := make([]int, 0)

	for i := range a {
		b := strings.Split(a[i], "\n")
		sum := 0
		for j := range b {
			ia, _ := strconv.Atoi(strings.TrimSpace(b[j]))
			sum += ia
		}
		arr = append(arr, sum)
	}
	max := 0
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return fmt.Sprintln(max)
}

func solve1_2(s string) string {
	a := strings.Split(s, "\n\n")
	arr := make([]int, 0)

	for i := range a {
		b := strings.Split(a[i], "\n")
		sum := 0
		for j := range b {
			ia, _ := strconv.Atoi(strings.TrimSpace(b[j]))
			sum += ia
		}
		arr = append(arr, sum)
	}

	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	return fmt.Sprintln(arr[0] + arr[1] + arr[2])
}
