package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	ans   []int
	p     []point
	nodes []node
	np    int
)

var sc = bufio.NewScanner(os.Stdin)

type point struct {
	id int
	x  int
	y  int
}

type sorterX struct {
	p   []point
	len int
}

func (s sorterX) Len() int {
	return s.len
}

func (s sorterX) Swap(i, j int) {
	tmp := s.p[i]
	s.p[i] = s.p[j]
	s.p[j] = tmp
}

func (s sorterX) Less(i, j int) bool {
	return s.p[i].x < s.p[j].x
}

type sorterY struct {
	p   []point
	len int
}

func (s sorterY) Len() int {
	return s.len
}

func (s sorterY) Swap(i, j int) {
	tmp := s.p[i]
	s.p[i] = s.p[j]
	s.p[j] = tmp
}

func (s sorterY) Less(i, j int) bool {
	return s.p[i].y < s.p[j].y
}

type node struct {
	location int
	l        int
	r        int
}

func makeKDTreeX(l, r int) int {
	if l >= r {
		return -1
	}

	mid := (l + r) / 2
	s := sorterX{
		p:   p[l:r],
		len: r - l,
	}
	sort.Sort(s)

	t := np
	np++
	nodes[t].location = mid
	nodes[t].l = makeKDTreeY(l, mid)
	nodes[t].r = makeKDTreeY(mid+1, r)

	return t
}

func makeKDTreeY(l, r int) int {
	if l >= r {
		return -1
	}

	mid := (l + r) / 2

	s := sorterY{
		p:   p[l:r],
		len: r - l,
	}
	sort.Sort(s)

	t := np
	np++
	nodes[t].location = mid
	nodes[t].l = makeKDTreeX(l, mid)
	nodes[t].r = makeKDTreeX(mid+1, r)

	return t
}

func find(t, sx, ex, sy, ey, depth int) {
	if t == -1 {
		return
	}
	x := p[nodes[t].location].x
	y := p[nodes[t].location].y

	if sx <= x && x <= ex && sy <= y && y <= ey {
		ans = append(ans, p[nodes[t].location].id)
	}

	if depth%2 == 0 {
		if sx <= x {
			find(nodes[t].l, sx, ex, sy, ey, depth+1)
		}
		if x <= ex {
			find(nodes[t].r, sx, ex, sy, ey, depth+1)
		}
	} else {
		if sy <= y {
			find(nodes[t].l, sx, ex, sy, ey, depth+1)
		}
		if y <= ey {
			find(nodes[t].r, sx, ex, sy, ey, depth+1)
		}
	}
}

// DSL_2_C: 領域探索
func main() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := nextInt()
	p = make([]point, n, n)
	nodes = make([]node, n, n)
	for i := 0; i < n; i++ {
		p[i].id = i
		p[i].x = nextInt()
		p[i].y = nextInt()
	}
	tree := makeKDTreeX(0, n)

	q := nextInt()
	for i := 0; i < q; i++ {
		ans = []int{}
		find(tree, nextInt(), nextInt(), nextInt(), nextInt(), 0)
		sort.Ints(ans)
		for _, v := range ans {
			fmt.Fprintln(w, v)
		}
		fmt.Fprintln(w)
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
