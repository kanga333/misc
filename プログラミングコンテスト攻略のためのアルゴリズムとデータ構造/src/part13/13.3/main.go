package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const max = math.MaxInt32

var adj [][]int

func dijkstra(n int) []int {
	d := make([]int, n, n)
	for i := range d {
		d[i] = max
	}
	d[0] = 0

	s := map[int]struct{}{0: {}}
	for len(adj) != len(s) {

		u := minEdge(s, d)
		updateNeighbor(d, u)
	}

	return d
}

func minEdge(s map[int]struct{}, d []int) int {
	min := max
	next := 0
	for i := range s {
		for j, w := range adj[i] {
			_, visited := s[j]
			weight := w + d[i]
			if weight < min && !visited {
				min = weight
				next = j
			}
		}
	}
	s[next] = struct{}{}
	d[next] = min
	return next
}

func updateNeighbor(d []int, u int) {
	for i, w := range adj[u] {
		if d[u]+w < d[i] {
			d[i] = d[u] + w
		}
	}
}

// ALDS1_12_B: 単一始点最短経路
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := scanInt(stdin)
	adj = make([][]int, n, n)
	for i := 0; i < n; i++ {
		line := make([]int, n, n)
		for j := range line {
			line[j] = max
		}

		in := scanIntArray(stdin)
		u := in[0]
		k := in[1]

		for j := 0; j < k; j++ {
			index := (j + 1) * 2
			v := in[index]
			c := in[index+1]
			line[v] = c
		}
		adj[u] = line
	}
	d := dijkstra(n)
	for i, v := range d {
		fmt.Fprintf(w, "%d %d\n", i, v)
	}
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
