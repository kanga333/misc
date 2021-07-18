package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const max = 2001

var adj [][]int

func prim() int {
	t := map[int]struct{}{0: {}}
	w := 0
	for len(adj) != len(t) {
		w += minEdge(t)
	}

	return w
}

func minEdge(t map[int]struct{}) int {
	min := max
	next := 0
	for i := range t {
		for j, w := range adj[i] {
			_, visited := t[j]
			if w < min && !visited {
				min = w
				next = j
			}
		}
	}
	t[next] = struct{}{}
	return min
}

// ALDS1_12_A: 最小全域木
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := scanInt(stdin)
	adj = make([][]int, n, n)
	for i := 0; i < n; i++ {
		adj[i] = scanIntArrayPrim(stdin)
	}
	fmt.Fprintln(w, prim())
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

func scanIntArrayPrim(sc *bufio.Scanner) []int {
	sc.Scan()
	t := strings.TrimSpace(sc.Text())

	ss := strings.Split(t, " ")
	is := []int{}
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		if i == -1 {
			i = max
		}
		is = append(is, i)
	}
	return is
}
