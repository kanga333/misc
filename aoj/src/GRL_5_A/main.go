package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var adj [][]edge

func dfs(current int, weight int, visited []bool) (int, int) {
	max := weight
	dst := current
	visited[current] = true

	for _, e := range adj[current] {
		if visited[e.dst] {
			continue
		}
		w, d := dfs(e.dst, weight+int(e.weight), visited)
		if w > max {
			max = w
			dst = d
		}
	}
	return max, dst
}

type edge struct {
	dst    int
	weight int16
}

//  GRL_5_A: 木の直径
func main() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := nextInt()

	adj = make([][]edge, n, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]edge, 0, 0)
	}
	for i := 0; i < n; i++ {
		src := nextInt()
		dst := nextInt()
		weight := int16(nextInt())
		adj[src] = append(adj[src], edge{
			dst:    dst,
			weight: weight,
		})
		adj[dst] = append(adj[dst], edge{
			dst:    src,
			weight: weight,
		})

	}
	_, d := dfs(0, 0, make([]bool, n, n))
	weight, _ := dfs(d, 0, make([]bool, n, n))
	fmt.Fprintln(w, weight)
}

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	s := next()
	i, _ := strconv.Atoi(s)
	return i
}
