# Adventofcode helper

The helper will help you with the following following 
1) Automatically retrieve the problem for a given day
2) Automatically submit, and parse the result, the solution for a given day


## Usage

```go

// dummy solver
def solver(inp string) string {
   return "some solution"
}

aoc := AdventOfCode{
    session: "53616c74asdadasdaasdadssda", // Session cookie taken from devtools in Chrome
    year:    2022,   // What year to solve
    day:     2,      // What day to solve
    level:   1,      // What level (every day has 2 levels)
    solver:  solver, // func(string)string
    submit:  true,   // Should you submit the solution (false won't submit, only print)
}

aoc.Solve() // This will trigger the actions

```

See main.go for how it's used
