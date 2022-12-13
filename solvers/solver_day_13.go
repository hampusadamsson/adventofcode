package solvers

import (
	"encoding/json"
	"fmt"
	"strings"
)

type problemPairs struct {
	p1, p2 []interface{}
}

func (p *problemPairs) inOrder() bool {
	// prove false
	p1order := p.p1 //.([]interface{})
	p2order := p.p2 //.([]interface{})
	for i := range p1order {
		isInOrder := pairCompare(p1order, p2order, i)
		if !isInOrder {
			return false
		}
	}
	return true
}

func pairCompare(p1, p2 []interface{}, i int) bool {
	if len(p2) <= i { // p2 is shorter than cursor index
		return true
	}

	fmt.Println(i, p1, "---", p2)
	fmt.Println(p1[i], "---", p2[i])

	v1, ok1 := p1[i].(float64)
	v2, ok2 := p2[i].(float64)

	fmt.Println("Both are floats", ok1, ok2)
	if ok1 == ok2 && ok1 { // both are floats
		return v1 <= v2
	} else if ok1 != ok2 && ok1 { // p1 val & p2 list
		var castInterface []interface{}
		castInterface = append(castInterface, v1)
		subPair := problemPairs{p1: castInterface, p2: p2[i].([]interface{})}
		return subPair.inOrder()
	} else if ok1 != ok2 && ok2 { // p1 list & p2 val
		var castInterface []interface{}
		castInterface = append(castInterface, v2)
		subPair := problemPairs{p2: castInterface, p1: p1[i].([]interface{})}
		return subPair.inOrder()
	} else if ok1 == ok2 && ok1 == false { // both are lists
		subPair := problemPairs{p1: p1[i].([]interface{}), p2: p2[i].([]interface{})}
		return subPair.inOrder()
	}

	return true
}

func Solve13_1(s string) string {
	// Create maze
	raw := strings.Split(s, "\n\n")

	var ps = []problemPairs{}
	for _, r := range raw[:len(raw)-1] {
		probs := strings.Split(r, "\n")
		p1 := probs[0]
		p2 := probs[1]

		var p1i []interface{}
		var p2i []interface{}

		json.Unmarshal([]byte(p1), &p1i)
		json.Unmarshal([]byte(p2), &p2i)
		ps = append(ps, problemPairs{p1: p1i, p2: p2i})
	}

	// Solve decrypt
	sum := 0
	for i, v := range ps {
		if v.inOrder() {
			sum += i
		}
		fmt.Println("Prob:", v)
	}
	return fmt.Sprint(sum)
}
