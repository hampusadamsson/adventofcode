package solvers

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strings"
)

type problemPairs struct {
	p1, p2 []interface{}
}

func (p *problemPairs) inOrder() (bool, bool) {
	// prove false
	p1order := p.p1 //.([]interface{})
	p2order := p.p2 //.([]interface{})
	rng := math.Max(float64(len(p1order)), float64(len(p2order)))
	for i := 0; i < int(rng); i++ {
		isInOrder, over := pairCompare(p1order, p2order, i)
		if over {
			return isInOrder, over
		}
	}
	return false, false
}

func pairCompare(p1, p2 []interface{}, i int) (bool, bool) { //order, over
	log := false

	if log {
		fmt.Println(i, p1, "---", p2)
	}
	// fmt.Println(p1[i], "---", p2[i])
	// fmt.Println("Both are floats", ok1, ok2)

	if len(p2) <= i { // p2 runs out
		if log {
			fmt.Println("OOL", len(p2), i, "-X-", p1, p2)
		}
		return false, true
	}
	if len(p1) <= i { // p1 runs out
		if log {
			fmt.Println("OOL", len(p1), i, "-X-", p1, p2)
		}
		return true, true
	}
	if log {
		fmt.Println(p1[i], "-?-", p2[i])
	}

	v1, ok1 := p1[i].(float64)
	v2, ok2 := p2[i].(float64)

	// fmt.Println(">>>", p1, p2, "--", p1[i], p2[i])

	if ok1 == ok2 && ok1 { // both are floats
		if log {
			fmt.Println("2 floats: ", v1 <= v2)
		}
		if v1 == v2 {
			return false, false
		}
		return v1 < v2, true
	} else if ok1 != ok2 && ok1 { // p1 val & p2 list
		if log {
			fmt.Println("1 float, 2 list")
		}
		var castInterface []interface{}
		castInterface = append(castInterface, v1)
		subPair := problemPairs{p1: castInterface, p2: p2[i].([]interface{})}
		return subPair.inOrder()
	} else if ok1 != ok2 && ok2 { // p1 list & p2 val
		if log {
			fmt.Println("1 list, 2 float")
		}
		var castInterface []interface{}
		castInterface = append(castInterface, v2)
		subPair := problemPairs{p2: castInterface, p1: p1[i].([]interface{})}
		return subPair.inOrder()
	} else if ok1 == ok2 && ok1 == false { // both are lists
		if log {
			fmt.Println("1 list, 2 list")
		}
		subPair := problemPairs{p1: p1[i].([]interface{}), p2: p2[i].([]interface{})}
		return subPair.inOrder()
	}
	panic("NOT HERE")
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

	// r, o := ps[3].inOrder()
	// fmt.Println("Prob:", r, o, ps[3])

	// return ""

	// Solve decrypt
	sum := 0
	for i, v := range ps {
		r, _ := v.inOrder()
		if r {
			sum += i + 1
		}
		fmt.Println("Prob:", r, v)
	}
	return fmt.Sprint(sum)
}

func Solve13_2(s string) string {
	// Create maze
	raw := strings.Split(s, "\n")

	raw = append(raw, "[[2]]")
	raw = append(raw, "[[6]]")

	ps := make([][]interface{}, 0)

	for i := range raw {
		if raw[i] != "" {
			var p1i []interface{}
			json.Unmarshal([]byte(raw[i]), &p1i)
			ps = append(ps, p1i)
		}
	}

	sort.Slice(ps, func(i, j int) bool {
		pp := problemPairs{p1: ps[i], p2: ps[j]}
		k, _ := pp.inOrder()
		return k
	})
	for _, v := range ps {
		fmt.Println(v)
	}

	f1 := 0
	s1 := 0
	for i := range ps {
		str := fmt.Sprintf("%v", ps[i])
		if str == "[[2]]" {
			f1 = i + 1
		}
		if str == "[[6]]" {
			s1 = i + 1
		}
	}

	fmt.Println(f1, s1)

	return fmt.Sprint(s1 * f1)
}
