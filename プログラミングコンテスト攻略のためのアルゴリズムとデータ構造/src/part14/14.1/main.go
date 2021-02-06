package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type node struct {
	id     int
	rank   int
	parent *node
}

func (n *node) findSet() *node {
	if n.id == n.parent.id {
		return n
	}
	root := n.parent.findSet()
	n.parent = root
	return root
}

type forest []node

func (f forest) same(x, y int) bool {
	xr := f[x].findSet()
	yr := f[y].findSet()
	return xr.id == yr.id
}

func (f forest) unite(x, y int) {
	xr := f[x].findSet()
	yr := f[y].findSet()
	if xr.rank < yr.rank {
		xr.parent = yr
		return
	}
	yr.parent = xr
	if xr.rank == yr.rank {
		xr.rank++
	}
}

// DSL_1_A: 互いに素な集合
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	a := scanIntArray(stdin)
	n := a[0]
	f := make(forest, n, n)
	for i := 0; i < n; i++ {
		v := node{
			id:   i,
			rank: 0,
		}
		v.parent = &v
		f[i] = v
	}
	q := a[1]
	for i := 0; i < q; i++ {
		a := scanIntArray(stdin)
		com := a[0]
		x := a[1]
		y := a[2]
		if com == 0 {
			f.unite(x, y)
		} else {
			if f.same(x, y) {
				fmt.Fprintln(w, 1)
			} else {
				fmt.Fprintln(w, 0)
			}
		}
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
