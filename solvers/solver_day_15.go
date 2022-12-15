package solvers

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type beacon struct {
	x, y, sigStrength int
	symbol            string
}

func Solve15_1(s string) string {
	//create
	maze := []beacon{}

	//Parse input
	lines := strings.Split(s, "\n")
	for _, l := range lines[:len(lines)-1] {
		r := regexp.MustCompile(`Sensor at x=(.*), y=(.*): closest beacon is at x=(.*), y=(.*)`)
		sol := r.FindStringSubmatch(l)

		x1, _ := strconv.Atoi(sol[1])
		y1, _ := strconv.Atoi(sol[2])

		x2, _ := strconv.Atoi(sol[3])
		y2, _ := strconv.Atoi(sol[4])

		sig := int(math.Abs((float64(x1) - float64(x2)))) + int(math.Abs((float64(y1) - float64(y2))))
		maze = append(maze, beacon{x: x1, y: y1, symbol: "S", sigStrength: sig})
		maze = append(maze, beacon{x: x2, y: y2, symbol: "B", sigStrength: 0})
	}
	fmt.Println(maze)

	// // Print maze
	// for y := 0; y < 23; y++ {
	// 	for x := -2; x < 26; x++ {
	// 		k, ok := maze[pos{x, y}]
	// 		if ok {
	// 			fmt.Print(k)
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	// //create
	visited := make(map[int]int)

	searchDepth := 2000000
	for _, k := range maze {
		if k.symbol == "S" {
			if k.y+k.sigStrength > searchDepth && k.y-k.sigStrength < searchDepth {
				delta := int(math.Abs(float64(k.y) - float64(searchDepth)))
				rem := k.sigStrength - delta

				for i := k.x - rem; i <= k.x+rem; i++ {
					// if k.x == 20 && k.y == 14 {
					// 	fmt.Println(i)
					// }
					visited[i] = i
				}
				// if k.x == 20 && k.y == 14 {
				// fmt.Println(k, "---", searchDepth, k.y > searchDepth, k.y-10 < searchDepth, delta, rem)
				// fmt.Println(len(visited), visited)
				// }
			}
		}
	}

	// // Print maze
	// for y := 0; y < 23; y++ {
	// 	for x := -2; x < 26; x++ {
	// 		found := false
	// 		for _, v := range maze {
	// 			if v.x == x && v.y == y {
	// 				fmt.Print(v.symbol)
	// 				found = true
	// 				break
	// 			}
	// 		}
	// 		if found == false {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	for _, v := range maze {
		if v.y == searchDepth && v.symbol == "B" {
			for _, k := range visited {
				if v.x == k {
					delete(visited, k)
					break
				}
			}
		}
	}

	return fmt.Sprint(len(visited))
}

func Solve15_2(s string) string {
	fmt.Println("Running...")
	//create
	maze := []beacon{}

	//Parse input
	lines := strings.Split(s, "\n")
	for _, l := range lines[:len(lines)-1] {
		r := regexp.MustCompile(`Sensor at x=(.*), y=(.*): closest beacon is at x=(.*), y=(.*)`)
		sol := r.FindStringSubmatch(l)

		x1, _ := strconv.Atoi(sol[1])
		y1, _ := strconv.Atoi(sol[2])

		x2, _ := strconv.Atoi(sol[3])
		y2, _ := strconv.Atoi(sol[4])

		sig := int(math.Abs((float64(x1) - float64(x2)))) + int(math.Abs((float64(y1) - float64(y2))))
		maze = append(maze, beacon{x: x1, y: y1, symbol: "S", sigStrength: sig})
		// maze = append(maze, beacon{x: x2, y: y2, symbol: "B", sigStrength: 0})
	}
	// fmt.Println(beacon{8, 7, 9, "S"}.getNextX(7, -1)) //10
	// fmt.Println(beacon{8, 7, 9, "S"}.getNextX(7, 0))  //11
	// fmt.Println(beacon{8, 7, 9, "S"}.getNextX(7, 14)) //10
	// fmt.Println(beacon{8, 7, 9, "S"}.getNextX(5, 14)) // false
	// fmt.Println(beacon{8, 7, 9, "S"}.getNextX(7, -2)) // false
	// fmt.Println(beacon{8, 7, 9, "S"}.getNextX(8, -2)) // 9, true
	// fmt.Println(beacon{8, 7, 9, "S"}.getNextX(9, -2)) // false

	// fmt.Println(beacon{8, 7, 9, "S"}.getNextX(8, 17)) // false
	// Prinnt
	// fmt.Println(visited)
	// Print maze
	maxr := 4000000
	for y := 0; y < maxr; y++ {
		for x := 0; x < maxr; x++ {
			goAgain := true
			for goAgain == true {
				for _, v := range maze {
					goAgain = false
					nx, ok := v.getNextX(x, y)
					// fmt.Println(x, y, "-->", nx, ok)
					if ok {
						x = nx
						goAgain = true
						break
					}
				}
				if goAgain == false && maxr > x {
					fmt.Println(x, y)
					return fmt.Sprint(x*4000000 + y)
				}
			}
		}
	}
	panic("Im here")
}

func (b beacon) getNextX(x, y int) (int, bool) {
	diffY := int(math.Max(float64(b.y), float64(y)) - math.Min(float64(b.y), float64(y)))
	delta := b.sigStrength - diffY + 1
	// fmt.Println(diffY, delta, b, (b.x - delta), (b.x + delta), x, y)
	if (b.x-delta) < x && (b.x+delta) > x {
		return b.x + b.sigStrength - diffY + 1, true
	}
	return -1, false
}

func (b beacon) adj() []beacon {
	fmt.Println(b.x, b.y, b.sigStrength)
	bs := []beacon{}
	// for x := 0; x <= b.sigStrength; x++ {
	// 	for y := 0 - b.sigStrength + x; y <= b.sigStrength-x; y++ {
	// 		bs = append(bs, beacon{x: b.x - x, y: b.y + y, sigStrength: 0, symbol: "#"})
	// 		bs = append(bs, beacon{x: b.x + x, y: b.y + y, sigStrength: 0, symbol: "#"})
	// 	}
	// }

	for x := b.x - b.sigStrength; x <= b.x; x++ {
		for y := b.y - b.sigStrength + (b.x - x); y <= b.y+b.sigStrength-(b.x-x); y++ {
			bs = append(bs, beacon{x: x, y: y, sigStrength: 0, symbol: "#"})
		}
	}
	for x := b.x; x <= b.x+b.sigStrength; x++ {
		for y := b.y - b.sigStrength + (x - b.x); y <= b.y+b.sigStrength-(x-b.x); y++ {
			bs = append(bs, beacon{x: x, y: y, sigStrength: 0, symbol: "#"})
		}
	}

	return bs
}
