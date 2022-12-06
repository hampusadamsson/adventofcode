package solvers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseProb(s string) []Stack {
	rows := strings.Split(s, "\n")

	lastRow := rows[len(rows)-1]
	lastRowN := strings.Split(lastRow, "   ")
	rows = rows[:len(rows)-1]

	stackCount, _ := strconv.Atoi(strings.TrimSpace((lastRowN[len(lastRowN)-1])))
	stacks := make([]Stack, stackCount)

	for _, v := range rows {
		for i := 0; i < stackCount; i++ {
			// fmt.Println(v, len(v), i*4, (i*4)+4)
			ss := v[i*4 : (i*4)+3]
			if strings.TrimSpace(ss) != "" {
				stacks[i].add(ss[1:2])
			}
		}
	}
	return stacks
}

type Stack struct {
	s []string
}

func (s *Stack) add(sol string) {
	s.s = append(s.s, sol)

}

func (s *Stack) prepend(sol []string) {
	// for i := len(sol) - 1; i > 0; i-- {
	// 	sss := sol[i]
	// 	s.s = append(s.s, sss)
	// }
	s.s = append(sol, s.s...)
}

func (s *Stack) remove(c int) []string {
	v := s.s[:c]
	s.s = s.s[c:]

	var ret = []string{}
	for _, v := range v {
		ret = append(ret, v)
	}
	return ret
}

func move(count int, from, to *Stack) {
	// for i := 0; i < count; i++ {
	// 	r := from.remove()
	// 	to.prepend(r)
	// }
	r := from.remove(count)
	to.prepend(r)
}

func Solve5_1(s string) string {
	a := strings.Split(s, "\n\n")

	prob := parseProb(a[0])

	inst := a[1]
	inss := strings.Split(inst, "\n")
	for _, v := range inss {
		fmt.Println(prob)
		if v == "" {
			break
		}

		r := regexp.MustCompile(` \d+`)
		insR := r.FindAllString(v, 3)

		quant, _ := strconv.Atoi(strings.TrimSpace(insR[0]))
		from, _ := strconv.Atoi(strings.TrimSpace(insR[1]))
		to, _ := strconv.Atoi(strings.TrimSpace(insR[2]))

		fmt.Println(v, quant, from, to)
		move(quant, &prob[from-1], &prob[to-1])
		fmt.Println(" ")
	}

	sol := ""
	for _, v := range prob {
		sol += v.s[0]
	}

	return fmt.Sprint(sol)
}

func Solve5_2(s string) string {
	a := strings.Split(s, "\n\n")

	prob := parseProb(a[0])

	inst := a[1]
	inss := strings.Split(inst, "\n")
	for _, v := range inss {
		fmt.Println(prob)
		if v == "" {
			break
		}

		r := regexp.MustCompile(` \d+`)
		insR := r.FindAllString(v, 3)

		quant, _ := strconv.Atoi(strings.TrimSpace(insR[0]))
		from, _ := strconv.Atoi(strings.TrimSpace(insR[1]))
		to, _ := strconv.Atoi(strings.TrimSpace(insR[2]))

		fmt.Println(v, quant, from, to)
		move(quant, &prob[from-1], &prob[to-1])

	}

	sol := ""
	for _, v := range prob {
		sol += v.s[0]
	}

	return fmt.Sprint(sol)
}
