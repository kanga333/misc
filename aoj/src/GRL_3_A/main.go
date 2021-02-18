package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var adj [][]int
var prenum, parent, lowest, out []int
var visited []bool
var timer = 1

func dfs(current, prev int) {

	prenum[current] = timer
	lowest[current] = timer
	timer++
	visited[current] = true

	for _, v := range adj[current] {
		if !visited[v] {
			parent[v] = current

			dfs(v, current)
			if lowest[current] > lowest[v] {
				lowest[current] = lowest[v]
			}
		} else if v != prev {
			if lowest[current] > prenum[v] {
				lowest[current] = prenum[v]
			}
		}
	}
}

// GRL_3_A: 関節点
func main() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	v := nextInt()
	e := nextInt()

	adj = make([][]int, v, v)
	prenum = make([]int, v, v)
	parent = make([]int, v, v)
	lowest = make([]int, v, v)
	visited = make([]bool, v, v)
	out = []int{}

	for i := 0; i < v; i++ {
		adj[i] = make([]int, 0, 0)
	}

	for i := 0; i < e; i++ {
		src := nextInt()
		dst := nextInt()
		adj[src] = append(adj[src], dst)
		adj[dst] = append(adj[dst], src)
	}

	dfs(0, 0)

	np := 0
	for i := 1; i < v; i++ {
		p := parent[i]
		if p == 0 {
			np++
		} else {
			if prenum[p] <= lowest[i] {
				out = append(out, p)
			}
		}
	}

	if np > 1 {
		out = append(out, 0)
	}

	sort.Sort(sort.IntSlice(out))

	prev := -1
	for _, v := range out {
		if v == prev {
			continue
		}
		fmt.Fprintln(w, v)
		prev = v
	}
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
