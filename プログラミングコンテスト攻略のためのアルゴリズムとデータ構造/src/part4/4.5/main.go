package main

import (
	"bufio"
	"fmt"
	"os"
)

type intStack struct {
	stack []int
	idx   int
}

func (s *intStack) push(i int) {
	s.stack[s.idx] = i
	s.idx++
}

func (s *intStack) pop() (int, bool) {
	if s.idx == 0 {
		return 0, false
	}
	val := s.stack[s.idx-1]
	s.idx--
	return val, true
}

func initStack() *intStack {
	s := make([]int, 20000)
	return &intStack{
		stack: s,
		idx:   0,
	}
}

type intStack2 struct {
	stack [][2]int
	idx   int
}

func (s *intStack2) push(i, v int) {
	s.stack[s.idx][0] = i
	s.stack[s.idx][1] = v
	s.idx++
}

func (s *intStack2) pop() (int, int, bool) {
	if s.idx == 0 {
		return 0, 0, false
	}
	val1 := s.stack[s.idx-1][0]
	val2 := s.stack[s.idx-1][1]
	s.idx--
	return val1, val2, true
}

func initStack2() *intStack2 {
	s := make([][2]int, 10000)
	return &intStack2{
		stack: s,
		idx:   0,
	}
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	slant := initStack()
	pools := initStack2()
	sum := 0

	t := stdin.Text()
	for i, c := range t {
		switch c {
		case rune('\\'):
			slant.push(i)
		case rune('/'):
			s, ok := slant.pop()
			if ok {
				pools.push(s, i-s)
				sum = sum + i - s
			}
		}

	}

	fmt.Println(sum)

	s := 1000000
	pool := 0
	ps := []int{}
	for {
		v1, v2, ok := pools.pop()
		if !ok {
			break
		}
		if s > v1 {
			if pool != 0 {
				ps = append(ps, pool)
				pool = 0
			}
			s = v1
		}
		pool = pool + v2
	}
	if pool != 0 {
		ps = append(ps, pool)
		pool = 0
	}
	fmt.Print(len(ps))
	for i := len(ps) - 1; i >= 0; i-- {
		fmt.Print(" ", ps[i])
	}
	fmt.Println()

}
