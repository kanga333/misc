package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	d, f []int
)

func visit(u, t int, adj [][]int) int {
	if d[u] == 0 {
		d[u] = t
	}
	f[u] = t
	for i, v := range adj[u] {
		if v == 1 && d[i] == 0 {
			t++
			t = visit(i, t, adj)
		}

	}
	t++
	f[u] = t
	return t
}

// ALDS1_11_B: 深さ優先探索
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := scanInt(stdin)
	adj := make([][]int, n, n)
	for i := 0; i < n; i++ {
		list := make([]int, n, n)
		a := scanIntArray(stdin)
		// 頂点
		u := a[0] - 1
		// 隣接する頂点
		vs := a[2:]

		for _, v := range vs {
			list[v-1] = 1
		}
		adj[u] = list
	}
	d = make([]int, n, n)
	f = make([]int, n, n)
	t := 0
	for i := 0; i < n; i++ {
		if d[i] == 0 {
			t++
			t = visit(i, t, adj)
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%d %d %d\n", i+1, d[i], f[i])
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
