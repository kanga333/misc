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

var adj [][]edge

type edge struct {
	weight int
	node   int
}


func dijkstra(n int) []int {
	d := make([]int, n, n)
	for i := range d {
		d[i] = max
	}
	d[0] = 0

	s := map[int]struct{}{0: {}}
	for len(adj) != len(s) {
		min := max
		u := 0
		for i, w := range d {
			_, visited := s[i]
			if w < min && !visited {
				min = w
				u = i
			}
		}
		s[u] = struct{}{}
		for _, e := range adj[u] {
			if d[u]+e.weight < d[e.node] {
				d[e.node] = d[u] + e.weight
			}
		}
	}
	return d
}

// ALDS1_12_C: 単一始点最短経路の高速化
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := scanInt(stdin)
	adj = make([][]edge, n, n)
	for i := 0; i < n; i++ {
		in := scanIntArray(stdin)
		u := in[0]
		k := in[1]
		line := make([]edge, k, k)

		for j := 0; j < k; j++ {
			index := (j + 1) * 2
			v := in[index]
			c := in[index+1]
			line[j] = edge{
				weight: c,
				node:   v,
			}
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
