package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const INF = math.MaxInt64

var sc = bufio.NewScanner(os.Stdin)
var adj [][]int64

// GRL_1_C: 全点対間最短経路
func main() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	v := nextInt()
	e := nextInt()

	adj = make([][]int64, v, v)
	for i := 0; i < v; i++ {
		adj[i] = make([]int64, v, v)
		for j := range adj[i] {
			if i != j {
				adj[i][j] = INF
			}
		}
	}

	for i := 0; i < e; i++ {
		src := nextInt()
		dst := nextInt()
		weight := nextInt()
		adj[src][dst] = int64(weight)
	}
	warsdhallFloyd(v)
	if negative(v) {
		fmt.Fprintln(w, "NEGATIVE CYCLE")
		return
	}
	for i := 0; i < v; i++ {
		if adj[i][0] == INF {
			fmt.Fprint(w, "INF")
		} else {
			fmt.Fprintf(w, "%d", adj[i][0])
		}

		for j := 1; j < v; j++ {
			if adj[i][j] == INF {
				fmt.Fprint(w, " INF")
			} else {
				fmt.Fprintf(w, " %d", adj[i][j])
			}
		}
		fmt.Fprint(w, "\n")
	}
}

func warsdhallFloyd(v int) {
	for k := 0; k < v; k++ {
		for i := 0; i < v; i++ {
			if adj[i][k] == INF {
				continue
			}
			for j := 0; j < v; j++ {
				if adj[k][j] == INF {
					continue
				}
				r := adj[i][k] + adj[k][j]
				if r < adj[i][j] {
					adj[i][j] = r
				}
			}
		}
	}
}

func negative(v int) bool {
	for i := 0; i < v; i++ {
		if adj[i][i] < 0 {
			return true
		}
	}
	return false
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
