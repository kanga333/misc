package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	ans []int
	p   []point
)

type point struct {
	id int
	x  int
	y  int
}

type node struct {
	location int
	l        *node
	r        *node
}

func makeKDTree(l, r, depth int) *node {
	if l >= r {
		return nil
	}

	mid := (l + r) / 2
	pp := p[l:r]
	if depth%2 == 0 {
		sort.Slice(pp, func(i, j int) bool {
			return pp[i].x < pp[j].x
		})
	} else {
		sort.Slice(pp, func(i, j int) bool {
			return pp[i].y < pp[j].y
		})
	}

	n := node{
		location: mid,
	}
	n.l = makeKDTree(l, mid, depth+1)
	n.r = makeKDTree(mid+1, r, depth+1)

	return &n
}

func find(n *node, sx, ex, sy, ey, depth int) {
	if n == nil {
		return
	}
	x := p[n.location].x
	y := p[n.location].y

	if sx <= x && x <= ex && sy <= y && y <= ey {
		ans = append(ans, p[n.location].id)
	}

	if depth%2 == 0 {
		if sx <= x {
			find(n.l, sx, ex, sy, ey, depth+1)
		}
		if x <= ex {
			find(n.r, sx, ex, sy, ey, depth+1)
		}
	} else {
		if sy <= y {
			find(n.l, sx, ex, sy, ey, depth+1)
		}
		if y <= ey {
			find(n.r, sx, ex, sy, ey, depth+1)
		}
	}
}

// DSL_2_C: 領域探索
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := scanInt(stdin)
	p = make([]point, n, n)
	for i := 0; i < n; i++ {
		a := scanIntArray(stdin)
		p[i] = point{
			id: i,
			x:  a[0],
			y:  a[1],
		}
	}
	tree := makeKDTree(0, n, 0)
	q := scanInt(stdin)
	for i := 0; i < q; i++ {
		ans = []int{}
		a := scanIntArray(stdin)
		find(tree, a[0], a[1], a[2], a[3], 0)
		sort.Ints(ans)
		for _, v := range ans {
			fmt.Fprintln(w, v)
		}
		fmt.Fprintln(w)
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
