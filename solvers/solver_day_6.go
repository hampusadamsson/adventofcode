package solvers

import (
	"fmt"
)

func Solve6_1(s string) string {
	for i := 4; i < len(s); i++ {
		sym := s[i-4 : i]
		set := make(map[rune]bool)
		for _, v := range sym {
			set[v] = true
		}
		fmt.Println(sym, i, len(set))
		if len(set) == 4 {
			return fmt.Sprintf("%d", i)
		}
	}
	panic("Err")
}

func Solve6_2(s string) string {
	for i := 14; i < len(s); i++ {
		sym := s[i-14 : i]
		set := make(map[rune]bool)
		for _, v := range sym {
			set[v] = true
		}
		fmt.Println(sym, i, len(set))
		if len(set) == 14 {
			return fmt.Sprintf("%d", i)
		}
	}
	panic("Err")
}
