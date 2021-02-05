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

func swap(h []edge, i, j int) {
	tmp := h[i]
	h[i] = h[j]
	h[j] = tmp
}

func left(i int) int {
	return i * 2
}

func right(i int) int {
	return i*2 + 1
}

func parent(i int) int {
	return i / 2
}

func maxHepify(h []edge, i int) {
	l := left(i)
	r := right(i)
	size := len(h)

	largest := i
	if l < size && h[l].weight > h[largest].weight {
		largest = l
	}
	if r < size && h[r].weight > h[largest].weight {
		largest = r
	}

	if largest != i {
		swap(h, largest, i)
		maxHepify(h, largest)
	}
}

func extractMax(h []edge) (edge, []edge) {
	max := h[1]
	end := len(h) - 1
	h[1] = h[end]
	h = h[:end]
	maxHepify(h, 1)
	return max, h
}

func insert(h []edge, e edge) []edge {
	i := len(h)
	h = append(h, e)
	for ; i > 1 && h[i].weight > h[parent(i)].weight; i = parent(i) {
		swap(h, i, parent(i))
	}
	return h
}

func dijkstra(n int) []int {
	d := make([]int, n, n)
	for i := range d {
		d[i] = max
	}
	d[0] = 0

	h := []edge{{0, 0}, {0, 0}}
	s := map[int]struct{}{}
	for len(h) != 1 {

		e := edge{}
		e, h = extractMax(h)

		u := e.node
		s[u] = struct{}{}
		if e.weight*-1 != d[u] {
			continue
		}

		for _, e := range adj[u] {
			_, visited := s[e.node]
			if visited {
				continue
			}
			if d[u]+e.weight < d[e.node] {
				d[e.node] = d[u] + e.weight
				h = insert(h, edge{d[e.node] * -1, e.node})
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
