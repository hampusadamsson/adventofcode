package solvers

import (
	"fmt"
	"strconv"
	"strings"
)

type fs struct {
	name   string
	f      []*fs
	size   int
	parent *fs
}

func (f *fs) getSize() int {
	sum := f.size
	for i, _ := range f.f {
		sum += f.f[i].getSize()
	}
	return sum
}

func (f *fs) cd(s string) *fs {
	if s == "/" {
		if f.parent == nil {
			return f
		} else {
			return f.parent.cd("/")
		}
	} else if s == ".." {
		return f.parent
	}
	for i, _ := range f.f {
		if f.f[i].name == s {
			return f.f[i]
		}
	}
	fmt.Println(f)
	panic("CD: " + s)
}

func (f *fs) find(s string) *fs {
	if f.name == s {
		return f
	}
	for i, _ := range f.f {
		found := f.f[i].find(s)
		if found != nil {
			return found
		}
	}
	return nil
}

func (f *fs) findMoreThan() int {
	sum := 0
	found := f.getSize()
	if f.size == 0 && found <= 100000 {
		fmt.Println(f.name, f.size, found)
		sum += found
	}
	for i, _ := range f.f {
		sum += f.f[i].findMoreThan()
	}
	return sum
}

func (f *fs) remove(small, req int) int {
	fmt.Println("------", f.name, f.size, small)
	found := f.getSize()
	if f.size == 0 && found >= req {
		if found < small {
			fmt.Println("SMALLER", f.name, f.size, found)
			small = found
		}
		small = found
	}
	for i, _ := range f.f {
		smaller := f.f[i].remove(small, req)
		if smaller < small {
			small = smaller
		}
	}
	return small
}

func (f *fs) mkdir(n string, size int) {
	f.f = append(f.f, &fs{name: n, size: size, parent: f, f: make([]*fs, 0)})
}

func Solve7_1(s string) string {
	files := &fs{name: "/", size: 0, parent: nil, f: make([]*fs, 0)}
	cur := files

	lines := strings.Split(s, "\n")
	for _, v := range lines {
		if v == "" {
			break
		}
		if v[:2] == "$ " { //command
			if v[2:4] == "cd" {
				cur = cur.cd(v[5:])
			} else if v[2:4] == "ls" {
				//nothing
			}
		}
		if v[:3] == "dir" {
			fmt.Println("Making dir")
			cur.mkdir(v[4:], 0)
		} else {
			meta := strings.Split(v, " ")
			size, _ := strconv.Atoi(meta[0])
			cur.mkdir(meta[1], size)
		}

		fmt.Println(v)
	}
	return fmt.Sprint(cur.cd("/").findMoreThan())
}

func Solve7_2(s string) string {
	files := &fs{name: "/", size: 0, parent: nil, f: make([]*fs, 0)}
	cur := files

	lines := strings.Split(s, "\n")
	for _, v := range lines {
		if v == "" {
			break
		}
		if v[:2] == "$ " { //command
			if v[2:4] == "cd" {
				cur = cur.cd(v[5:])
			} else if v[2:4] == "ls" {
				//nothing
			}
		}
		if v[:3] == "dir" {
			fmt.Println("Making dir")
			cur.mkdir(v[4:], 0)
		} else {
			meta := strings.Split(v, " ")
			size, _ := strconv.Atoi(meta[0])
			cur.mkdir(meta[1], size)
		}

		fmt.Println(v)
	}
	full := cur.cd("/").getSize()
	req := full - 40000000
	fmt.Println("FULL", req)

	return fmt.Sprint(cur.cd("/").remove(999999999999999, req))
}
