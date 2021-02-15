package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var adj [][]int
var degree, out, visited []int

func tsort(s int) {
	queue := []int{s}
	for len(queue) != 0 {
		qv := queue[0]
		queue = queue[1:]
		out = append(out, qv)
		for _, v := range adj[qv] {
			degree[v]--
			if degree[v] == 0 {
				visited[v] = 1
				queue = append(queue, v)
			}
		}

	}

}

// GRL_1_C: 全点対間最短経路
func main() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	v := nextInt()
	e := nextInt()

	adj = make([][]int, v, v)
	degree = make([]int, v, v)
	visited = make([]int, v, v)
	out = []int{}
	for i := 0; i < v; i++ {
		adj[i] = make([]int, 0, 0)
	}

	for i := 0; i < e; i++ {
		src := nextInt()
		dst := nextInt()
		adj[src] = append(adj[src], dst)
		degree[dst]++
	}

	for i := 0; i < v; i++ {
		if degree[i] == 0 && visited[i] == 0 {
			tsort(i)
		}
	}

	for _, v := range out {
		fmt.Fprintln(w, v)
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
