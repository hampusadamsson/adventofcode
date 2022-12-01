package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Solver has to implement following methods
type Solver interface {
	solve1(string) string
	solve2(string) string
}

// AdventOfCode can download the problem, and upload a solution
type AdventOfCode struct {
	session string
	year    int
	solver  func(string) string
	day     int
	level   int
	submit  bool
}

// Solve - retrieve, solve, and parse the response
func (aoc *AdventOfCode) Solve() bool {
	problem := aoc.getProblem()
	sol := aoc.solver(problem)
	fmt.Println(fmt.Sprintf("Day: %d\t Solution: %s", aoc.day, sol))
	if aoc.submit {
		resp := aoc.submitSolution(aoc.level, sol)
		parsedResponse := aoc.parseAnswer(resp)
		fmt.Println(fmt.Sprintf("Solution valid: %t", parsedResponse))
		return parsedResponse
	} else {
		fmt.Println("Won't submit")
		return false
	}
}

// getProblem - retrieve the problem for the given advent-of-code
func (aoc *AdventOfCode) getProblem() string {
	fmt.Println(fmt.Sprintf("Retrieving problem for day: %d", aoc.day))
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", aoc.year, aoc.day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Cookie", "session="+aoc.session)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	s := string(resBody)

	fmt.Println(fmt.Sprintf("Problem length: %d runes\nProblem: %s...%s", len(s), s[:3], s[len(s)-3:]))
	if len(s) < 3 {
		panic("Problem retrieval problems...")
	}

	return s
}

// submitSolution - level is either 1 or 2, ans is the solution as string
func (aoc *AdventOfCode) submitSolution(level int, ans string) string {
	fmt.Println(fmt.Sprintf("Submitting day:%d, level:%d, sol:%s", aoc.day, level, ans))

	params := url.Values{}
	params.Add("level", fmt.Sprintf("%d", level))
	params.Add("answer", ans)
	body := strings.NewReader(params.Encode())

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", aoc.year, aoc.day)
	req, err := http.NewRequest("POST", url, body)

	if err != nil {
		panic(err)
	}
	req.Header.Set("Cookie", "session="+aoc.session)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)

	return string(resBody)
}

func (aoc *AdventOfCode) parseAnswer(ans string) bool {
	if strings.Contains(ans, "That's the right answer!") {
		fmt.Println("You solved the problem")
		return true
	}

	if strings.Contains(ans, "You don't seem to be solving the right level.  Did you already complete it") {
		fmt.Println("Already solved the problem")
		return true
	}

	if strings.Contains(ans, "before it unlocks") {
		fmt.Println("Problem not yet available")
		return false
	}

	if strings.Contains(ans, "You're posting too much data") {
		fmt.Println("You're posting too much data")
		return false
	}

	if strings.Contains(ans, "That's not the right answer") {
		fmt.Println("Wrong answer")
		return false
	}

	if strings.Contains(ans, "You gave an answer too recently; you have to wait after submitting an answer before trying again.  You have") {
		fmt.Println("Too recent")
		return false
	}
	fmt.Println(ans)
	panic("You should not be here")
}
