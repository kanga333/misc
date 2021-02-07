package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ALDS1_11_A: グラフ
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
	for _, list := range adj {
		for i, val := range list {
			if i == 0 {
				fmt.Fprintf(w, "%d", val)
			} else {
				fmt.Fprintf(w, " %d", val)
			}
		}
		fmt.Fprintf(w, "\n")

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
