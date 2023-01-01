package solvers

import (
	"fmt"
	"strconv"
	"strings"
)

func getPosBetween(fromr, tor pos) []pos {
	ps := []pos{}

	// assert order that from is smaller
	var to pos
	var from pos
	if fromr.x > tor.x || fromr.y > tor.y {
		to = fromr
		from = tor
	} else {
		to = tor
		from = fromr
	}

	if from.x == to.x {
		for i := from.y; i <= to.y; i++ {
			ps = append(ps, pos{x: from.x, y: i})
		}
	} else {
		for i := from.x; i <= to.x; i++ {
			ps = append(ps, pos{x: i, y: from.y})
		}
	}
	fmt.Println(ps)
	return ps
}

func Solve14_1(s string) string {
	maze := make(map[pos]string)
	maze[pos{x: 500, y: 0}] = "+"

	raw := strings.Split(s, "\n")
	for _, ins := range raw[:len(raw)-1] {
		ridge := strings.Split(ins, " -> ")
		for i := range ridge[:len(ridge)-1] {
			point1 := strings.Split(ridge[i], ",")
			point2 := strings.Split(ridge[i+1], ",")
			x1, _ := strconv.Atoi(point1[0])
			y1, _ := strconv.Atoi(point1[1])
			x2, _ := strconv.Atoi(point2[0])
			y2, _ := strconv.Atoi(point2[1])
			p1 := pos{x: x1, y: y1}
			p2 := pos{x: x2, y: y2}
			betweens := getPosBetween(p1, p2)
			for j := range betweens {
				c := betweens[j]
				maze[c] = "#"
			}
		}
	}

	fmt.Println(maze)
	// Print maze
	printMaze(maze)
	// Simulate
	ended := 1
	for ended == 1 {
		// one sand
		sand := pos{x: 500, y: 0}
		x := 1
		for x == 1 {
			if _, ok := maze[pos{x: sand.x, y: sand.y + 1}]; ok == false {
				sand = pos{x: sand.x, y: sand.y + 1}
			} else if _, ok := maze[pos{x: sand.x - 1, y: sand.y + 1}]; ok == false {
				sand = pos{x: sand.x - 1, y: sand.y + 1}
			} else if _, ok := maze[pos{x: sand.x + 1, y: sand.y + 1}]; ok == false {
				sand = pos{x: sand.x + 1, y: sand.y + 1}
			} else {
				maze[sand] = "O"
				x = 0
			}
			// printMaze(maze)

			if sand.y > 1000 {
				ended = 0
				x = 0
			}
		}
	}
	sum := 0
	for v := range maze {
		if maze[v] == "O" {
			sum++
		}
	}
	return fmt.Sprint(sum)
}

func printMaze(maze map[pos]string) {
	for y := 0; y <= 15; y++ {
		for x := 480; x <= 510; x++ {
			if v, ok := maze[pos{x: x, y: y}]; ok {
				fmt.Print(v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func Solve14_2(s string) string {
	maze := make(map[pos]string)
	maze[pos{x: 500, y: 0}] = "+"

	raw := strings.Split(s, "\n")
	for _, ins := range raw[:len(raw)-1] {
		ridge := strings.Split(ins, " -> ")
		for i := range ridge[:len(ridge)-1] {
			point1 := strings.Split(ridge[i], ",")
			point2 := strings.Split(ridge[i+1], ",")
			x1, _ := strconv.Atoi(point1[0])
			y1, _ := strconv.Atoi(point1[1])
			x2, _ := strconv.Atoi(point2[0])
			y2, _ := strconv.Atoi(point2[1])
			p1 := pos{x: x1, y: y1}
			p2 := pos{x: x2, y: y2}
			betweens := getPosBetween(p1, p2)
			for j := range betweens {
				c := betweens[j]
				maze[c] = "#"
			}
		}
	}

	fmt.Println(maze)

	max := 0
	for v := range maze {
		if v.y > max {
			max = v.y
		}
	}
	max += 2

	// Print maze
	printMaze(maze)
	// Simulate
	ended := 1
	for ended == 1 {
		// one sand
		sand := pos{x: 500, y: 0}
		x := 1
		for x == 1 {
			if sand.y+1 == max {
				maze[sand] = "O"
				x = 0
				break
			}
			if _, ok := maze[pos{x: sand.x, y: sand.y + 1}]; ok == false {
				sand = pos{x: sand.x, y: sand.y + 1}
			} else if _, ok := maze[pos{x: sand.x - 1, y: sand.y + 1}]; ok == false {
				sand = pos{x: sand.x - 1, y: sand.y + 1}
			} else if _, ok := maze[pos{x: sand.x + 1, y: sand.y + 1}]; ok == false {
				sand = pos{x: sand.x + 1, y: sand.y + 1}
			} else {
				maze[sand] = "O"
				x = 0
			}
			// printMaze(maze)

			if sand.y == 0 && sand.x == 500 {
				ended = 0
				x = 0
			}
		}
	}
	sum := 0
	for v := range maze {
		if maze[v] == "O" {
			sum++
		}
	}
	return fmt.Sprint(sum)
}
