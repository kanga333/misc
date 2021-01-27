package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func quickSort(a []int, p, r int) {
	if p < r {
		q := partition(a, p, r)
		quickSort(a, p, q-1)
		quickSort(a, q+1, r)
	}

}

func partition(a []int, p, r int) int {
	x := a[r]
	i := p
	for j := p; j < r; j++ {
		if a[j] <= x {
			swap(a, i, j)
			i++
		}
	}
	swap(a, i, r)
	return i
}

func swap(a []int, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp

}

func toMap(a []int) map[int]int {
	m := make(map[int]int, len(a))
	for i, v := range a {
		m[v] = i
	}
	return m
}

func detectCycle(a, b []int) [][]int {
	cycles := make([][]int, 0)
	am := toMap(a)
	bm := toMap(b)

	for i, bi := range b {
		if bi == a[i] || bm[bi] == -1 {
			continue
		}
		cycle := extractCycle(b, am, bm, bi)
		cycles = append(cycles, cycle)
	}
	return cycles
}

func extractCycle(b []int, am, bm map[int]int, val int) []int {
	if bm[val] == -1 {
		return nil
	}
	bm[val] = -1
	r := []int{val}
	next := b[am[val]]
	ret := extractCycle(b, am, bm, next)
	if ret == nil {
		return r
	}
	return append(r, ret...)
}

func sum(a []int) int {
	s := 0
	for _, val := range a {
		s += val
	}
	return s
}

func solve(cycle []int, min int) int {
	s := sum(cycle)

	left := s + (len(cycle)-2)*cycle[0]
	right := s + cycle[0] + (len(cycle)+1)*min
	if left < right {
		return left
	}
	return right
}

// ALDS1_6_D:最小コストソート
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	n := scanInt(stdin)
	a := scanIntArray(stdin)
	b := make([]int, n)
	copy(b, a)
	quickSort(b, 0, n-1)
	cycles := detectCycle(a, b)
	val := 0
	for _, cycle := range cycles {
		val += solve(cycle, b[0])
	}
	fmt.Printf("%d\n", val)

}

func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	t := sc.Text()
	i, err := strconv.Atoi(t)
	if err != nil {
		panic(err)
	}
	return i
}

func scanIntArray(sc *bufio.Scanner) []int {
	sc.Scan()
	t := sc.Text()
	ss := strings.Split(t, " ")
	is := []int{}
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		is = append(is, i)
	}
	return is
}
