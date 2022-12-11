package solvers

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
)

type monkey struct {
	name  string
	costf string
	items []int
	// stress    int
	divisible int
	trueCase  int
	falseCase int
	inspects  int
}

func executeEq(eq string, stress int) int {
	expression, _ := govaluate.NewEvaluableExpression(eq)
	parameters := make(map[string]interface{}, 8)
	parameters["old"] = stress
	result, _ := expression.Evaluate(parameters)
	rFloat := result.(float64)
	return int(rFloat)
}

func getTotalDivide(m []monkey) int {
	sum := 1
	for _, v := range m {
		sum *= v.divisible
	}
	return sum
}

func (m *monkey) process(ms []monkey) {
	for _, item := range m.items {
		// inspect
		m.inspects++
		item = executeEq(m.costf, item)
		// item = item / m.divisible
		nyitem := item % getTotalDivide(ms)
		if item%m.divisible == 0 {
			ms[m.trueCase].items = append(ms[m.trueCase].items, nyitem)
		} else {
			ms[m.falseCase].items = append(ms[m.falseCase].items, nyitem)
		}
	}
	m.items = []int{}
}

func Solve11_1(s string) string {
	raw := strings.Split(s, "\n\n")

	// Parse monkey
	var monkeys = []monkey{}
	for _, mraw := range raw {
		minst := strings.Split(mraw, "\n")
		// Start items
		startItem := strings.ReplaceAll(minst[1], "Starting items: ", "")
		var startItems []int
		for _, v := range strings.Split(startItem, ", ") {
			iv, _ := strconv.Atoi(strings.TrimSpace(v))
			startItems = append(startItems, iv)
		}
		// Operations
		costFunc := strings.ReplaceAll(minst[2], "Operation: new = ", "")
		// Test
		testRaw := strings.ReplaceAll(minst[3], "Test: divisible by ", "")
		divisible, _ := strconv.Atoi(strings.TrimSpace(testRaw))
		// True
		trueCase := strings.ReplaceAll(minst[4], "If true: throw to monkey ", "")
		trueCaseInt, _ := strconv.Atoi(strings.TrimSpace(trueCase))
		// False
		falseCase := strings.ReplaceAll(minst[5], "If false: throw to monkey ", "")
		falseCaseInt, _ := strconv.Atoi(strings.TrimSpace(falseCase))

		m := monkey{
			name:      minst[0],
			items:     startItems,
			costf:     costFunc,
			divisible: divisible,
			trueCase:  trueCaseInt,
			falseCase: falseCaseInt,
		}
		monkeys = append(monkeys, m)
	}

	for v := range monkeys {
		fmt.Println(monkeys[v])
	}

	// Play all
	for i := 0; i < 20; i++ {
		for i := range monkeys {
			monkeys[i].process(monkeys)
		}
	}
	fmt.Println("----")
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspects > monkeys[j].inspects
	})

	for v := range monkeys {
		fmt.Println(monkeys[v])
	}

	return fmt.Sprint(monkeys[0].inspects * monkeys[1].inspects)
}

func Solve11_2(s string) string {
	raw := strings.Split(s, "\n\n")
	// Parse monkey
	var monkeys = []monkey{}
	for _, mraw := range raw {
		minst := strings.Split(mraw, "\n")
		// Start items
		startItem := strings.ReplaceAll(minst[1], "Starting items: ", "")
		var startItems []int
		for _, v := range strings.Split(startItem, ", ") {
			iv, _ := strconv.Atoi(strings.TrimSpace(v))
			startItems = append(startItems, iv)
		}
		// Operations
		costFunc := strings.ReplaceAll(minst[2], "Operation: new = ", "")
		// Test
		testRaw := strings.ReplaceAll(minst[3], "Test: divisible by ", "")
		divisible, _ := strconv.Atoi(strings.TrimSpace(testRaw))
		// True
		trueCase := strings.ReplaceAll(minst[4], "If true: throw to monkey ", "")
		trueCaseInt, _ := strconv.Atoi(strings.TrimSpace(trueCase))
		// False
		falseCase := strings.ReplaceAll(minst[5], "If false: throw to monkey ", "")
		falseCaseInt, _ := strconv.Atoi(strings.TrimSpace(falseCase))

		m := monkey{
			name:      minst[0],
			items:     startItems,
			costf:     costFunc,
			divisible: divisible,
			trueCase:  trueCaseInt,
			falseCase: falseCaseInt,
		}
		monkeys = append(monkeys, m)
	}

	for v := range monkeys {
		fmt.Println(monkeys[v])
	}

	// Play all
	for i := 0; i < 10000; i++ {
		for i := range monkeys {
			monkeys[i].process(monkeys)
		}
	}
	fmt.Println("----")
	for v := range monkeys {
		fmt.Println(monkeys[v].inspects, "----", monkeys[v])
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspects > monkeys[j].inspects
	})

	return fmt.Sprint(monkeys[0].inspects * monkeys[1].inspects)
}
