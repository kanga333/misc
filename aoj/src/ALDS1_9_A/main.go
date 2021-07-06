package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type tree struct {
	key    int
	parent *tree
	left   *tree
	right  *tree
}

func (t *tree) state() string {
	ret := fmt.Sprintf("key = %d, ", t.key)
	if t.parent != nil {
		ret += fmt.Sprintf("parent key = %d, ", t.parent.key)
	}
	if t.left != nil {
		ret += fmt.Sprintf("left key = %d, ", t.left.key)
	}

	if t.right != nil {
		ret += fmt.Sprintf("right key = %d, ", t.right.key)
	}
	return ret
}

type heap struct {
	trees []*tree
	last  int
}

func newHeap(n int) *heap {
	// 1 originの配列
	trees := make([]*tree, n+1, n+1)
	return &heap{
		trees: trees,
		last:  1,
	}
}

func (h *heap) append(key int) {
	p := h.trees[h.last/2]
	t := &tree{
		key:    key,
		parent: p,
	}
	h.trees[h.last] = t
	h.last++

	// tはroot node
	if p == nil {
		return
	}

	if p.left == nil {
		p.left = t
		return
	}
	p.right = t
}

func (h *heap) print(w io.Writer) {
	for i, t := range h.trees {
		if i == 0 {
			continue
		}
		fmt.Fprintf(w, "node %d: %s\n", i, t.state())
	}
}

// ALDS1_9_A: 完全二分木
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := scanInt(stdin)
	h := newHeap(n)
	a := scanIntArray(stdin)

	for _, key := range a {
		h.append(key)
	}
	h.print(w)
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
