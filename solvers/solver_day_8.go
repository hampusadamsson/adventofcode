package solvers

import (
	"fmt"
	"strconv"
	"strings"
)

type pos struct {
	x, y int
}

type cell struct {
	c int
	p pos
}

func (c cell) visible(cells map[pos]cell, rows, cols int) bool {
	for i := 1; i < 10000; i++ {
		if v, ok := cells[pos{x: c.p.x + i, y: c.p.y}]; ok {
			if v.c >= c.c {
				break
			}
		} else {
			return true
		}
	}
	for i := 1; i < 10000; i++ {
		if v, ok := cells[pos{x: c.p.x - i, y: c.p.y}]; ok {
			if v.c >= c.c {
				break
			}
		} else {
			return true
		}
	}
	for i := 1; i < 10000; i++ {
		if v, ok := cells[pos{x: c.p.x, y: c.p.y - i}]; ok {
			if v.c >= c.c {
				break
			}
		} else {
			return true
		}
	}
	for i := 1; i < 10000; i++ {
		if v, ok := cells[pos{x: c.p.x, y: c.p.y + i}]; ok {
			if v.c >= c.c {
				break
			}
		} else {
			return true
		}
	}
	fmt.Println(c)
	return false
}

func (c cell) treeHouse(cells map[pos]cell, rows, cols int) int {
	totSum := 1
	sum := 0
	for i := 1; i < 10000; i++ {
		if v, ok := cells[pos{x: c.p.x + i, y: c.p.y}]; ok {
			sum++
			if v.c >= c.c {
				break
			}
		}
	}
	if sum == 0 {
		sum = 1
	}
	totSum *= sum
	sum = 0
	for i := 1; i < 10000; i++ {
		if v, ok := cells[pos{x: c.p.x - i, y: c.p.y}]; ok {
			sum++
			if v.c >= c.c {
				break
			}
		}
	}
	if sum == 0 {
		sum = 1
	}
	totSum *= sum
	sum = 0
	for i := 1; i < 10000; i++ {
		if v, ok := cells[pos{x: c.p.x, y: c.p.y - i}]; ok {
			sum++
			if v.c >= c.c {
				break
			}
		}
	}
	if sum == 0 {
		sum = 1
	}
	totSum *= sum
	sum = 0
	for i := 1; i < 10000; i++ {
		if v, ok := cells[pos{x: c.p.x, y: c.p.y + i}]; ok {
			sum++
			if v.c >= c.c {
				break
			}
		}
	}
	if sum == 0 {
		sum = 1
	}
	totSum *= sum
	sum = 0
	return totSum
}

func stringToMaze(s string) map[pos]cell {
	rows := strings.Split(s, "\n")
	cells := make(map[pos]cell)
	for y, row := range rows {
		for x, c := range row {
			cost, _ := strconv.Atoi(string(c))
			p := pos{x: x, y: y}
			cells[p] = cell{cost, p}
		}
	}
	return cells
}

func Solve8_1(s string) string {
	cells := stringToMaze(s)
	rowone := strings.Split(s, "\n")
	sum := 0
	for p2, _ := range cells {
		if cells[p2].visible(cells, len(rowone), len(rowone[0])) {
			sum++
		}
	}
	return fmt.Sprint(sum)
}

func Solve8_2(s string) string {
	cells := stringToMaze(s)
	rowone := strings.Split(s, "\n")
	sum := 0
	for p2, _ := range cells {
		if cells[p2].p.x != len(rowone)-1 && cells[p2].p.x != 0 && cells[p2].p.y != 0 && cells[p2].p.y != len(rowone[0])-1 {
			th := cells[p2].treeHouse(cells, len(rowone), len(rowone[0]))
			fmt.Println(p2, cells[p2].c, "treehouse:", th)
			if th > sum {
				sum = th
			}
		}
	}
	return fmt.Sprint(sum)
}
