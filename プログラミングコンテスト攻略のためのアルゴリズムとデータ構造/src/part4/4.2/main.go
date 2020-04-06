package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	a []int
	i int
}

func newStack() stack {
	return stack{
		a: make([]int, 200),
		i: 0,
	}
}

func (s *stack) push(num int) {
	s.a[s.i] = num
	s.i++
}

func (s *stack) pop() int {
	s.i--
	return s.a[s.i]
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	in := stdin.Text()
	ss := strings.Split(in, " ")
	stack := newStack()

	for _, s := range ss {
		switch s {
		case "+":
			l := stack.pop()
			r := stack.pop()
			stack.push(r + l)
		case "-":
			l := stack.pop()
			r := stack.pop()
			stack.push(r - l)
		case "*":
			l := stack.pop()
			r := stack.pop()
			stack.push(r * l)
		default:
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			stack.push(n)
		}
	}
	fmt.Println(stack.pop())
}
