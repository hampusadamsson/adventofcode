package solvers

import (
	"fmt"
	"strings"
)

func Solve12_1(s string) string {
	// Create maze
	raw := strings.Split(s, "\n")
	var start coord
	var end coord
	maze := make(map[coord]int)
	for y, r := range raw[:len(raw)-1] {
		for x, c2 := range r {
			c := c2
			if c == 'S' {
				start = coord{x: x, y: y}
				c = 'a'
			}
			if c == 'E' {
				end = coord{x: x, y: y}
				c = 'z'
			}
			maze[coord{x: x, y: y}] = int(c)
		}
	}

	// Solve maze
	hist := make(map[coord]int)
	hist[start] = 0
	pf := pathfidner{
		maze: maze,
		hist: hist,
	}
	fmt.Println(pf.hist)
	pf.find(start, end)
	fmt.Println(pf.hist)
	return fmt.Sprint(pf.hist[end])
}

func Solve12_2(s string) string {
	// Create maze
	raw := strings.Split(s, "\n")
	var starts = []coord{}
	var end coord
	maze := make(map[coord]int)
	for y, r := range raw[:len(raw)-1] {
		for x, c2 := range r {
			c := c2
			if c == 'S' {
				c = 'a'
			}
			if c == 'a' {
				starts = append(starts, coord{x: x, y: y})
			}
			if c == 'E' {
				end = coord{x: x, y: y}
				c = 'z'
			}
			maze[coord{x: x, y: y}] = int(c)
		}
	}

	// Solve maze
	best := 9999999999
	fmt.Println("N:", len(starts))
	for i, start := range starts {
		hist := make(map[coord]int)
		hist[start] = 0
		pf := pathfidner{
			maze: maze,
			hist: hist,
		}
		pf.find(start, end)
		b := pf.hist[end]
		fmt.Println(i, "/", len(starts), start, best, b)
		if b < best && b != 0 {
			best = b
		}
	}
	return fmt.Sprint(best)
}

func (p *pathfidner) find(from, goal coord) {
	if from == goal {
		return // reached dest
	}
	adj := from.adj()
	for i := range adj {
		nextCoord := adj[i]
		if p.valid(nextCoord) { // Exist in maze ?
			if (p.maze[from] + 1) >= p.maze[nextCoord] { // is this a valid move? c > b
				if t, ok := p.hist[nextCoord]; ok { // visited before?
					if (p.hist[from] + 1) < t {
						p.hist[nextCoord] = p.hist[from] + 1
						p.find(nextCoord, goal)
					}
				} else { // never seen before
					p.hist[adj[i]] = p.hist[from] + 1
					p.find(adj[i], goal)
				}
			}
		}
	}
}

func (p *pathfidner) valid(coord coord) bool {
	if _, ok := p.maze[coord]; ok {
		return true
	}
	return false
}

type coord struct {
	x int
	y int
}

func (p *coord) adj() []coord {
	return []coord{
		{x: p.x - 1, y: p.y},
		{x: p.x + 1, y: p.y},
		{x: p.x, y: p.y - 1},
		{x: p.x, y: p.y + 1},
	}
}

type pathfidner struct {
	maze map[coord]int
	hist map[coord]int
}
