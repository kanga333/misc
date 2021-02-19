package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var adj [][]node

func dfs(current int, weight int, visited []bool) (int, int) {
	max := weight
	dest := current
	visited[current] = true

	for _, n := range adj[current] {
		if visited[n.dest] {
			continue
		}
		ret, d := dfs(n.dest, weight+int(n.weight), visited)
		if ret > max {
			max = ret
			dest = d
		}
	}
	return max, dest
}

type node struct {
	dest   int
	weight int16
}

//  GRL_5_A: 木の直径
func main() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := nextInt()
	adj = make([][]node, n, n)

	for i := 0; i < n; i++ {
		adj[i] = make([]node, 0, 0)
	}
	for i := 0; i < n; i++ {
		src := nextInt()
		dst := nextInt()
		weight := int16(nextInt())
		adj[src] = append(adj[src], node{
			dest:   dst,
			weight: weight,
		})
		adj[dst] = append(adj[dst], node{
			dest:   src,
			weight: weight,
		})

	}
	visited := make([]bool, n, n)
	_, edge := dfs(0, 0, visited)
	visited = make([]bool, n, n)
	ans, _ := dfs(edge, 0, visited)
	fmt.Fprintln(w, ans)
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
