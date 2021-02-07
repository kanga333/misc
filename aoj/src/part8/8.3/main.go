package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type tree struct {
	id     int
	depth  int
	height int
	parent *tree
	left   *tree
	right  *tree
}

func (t *tree) calcDepth(depth int) {
	t.depth = depth
	if t.left != nil {
		t.left.calcDepth(depth + 1)
	}
	if t.right != nil {
		t.right.calcDepth(depth + 1)
	}
}

func (t *tree) calcHeight() int {
	l, r := 0, 0
	if t.left != nil {
		l = t.left.calcHeight() + 1
	}
	if t.right != nil {
		r = t.right.calcHeight() + 1
	}
	if l < r {
		t.height = r
		return r
	}
	t.height = l
	return l
}

func (t *tree) nodeType() string {
	if t.parent == nil {
		return "root"
	}
	if t.left == nil && t.right == nil {
		return "leaf"
	}

	return "internal node"
}

func (t *tree) pID() int {
	if t.parent == nil {
		return -1
	}
	return t.parent.id
}

func (t *tree) sID() int {
	if t.parent == nil {
		return -1
	}

	if s := t.parent.left; s != nil && s != t {
		return s.id
	}

	if s := t.parent.right; s != nil && s != t {
		return s.id
	}
	return -1
}

func (t *tree) degree() int {
	degree := 0
	if t.left != nil {
		degree++
	}
	if t.right != nil {
		degree++
	}
	return degree
}

func (t *tree) state() string {
	return fmt.Sprintf(
		"node %d: parent = %d, sibling = %d, degree = %d, depth = %d, height = %d, %s",
		t.id, t.pID(), t.sID(), t.degree(), t.depth, t.height, t.nodeType())
}

func initTrees(n int) []*tree {
	trees := make([]*tree, n, n)
	for i := 0; i < n; i++ {
		trees[i] = &tree{
			id: i,
		}
	}
	return trees
}

func setTree(id, left, right int, trees []*tree) {
	p := trees[id]
	if left != -1 {
		l := trees[left]
		l.parent = p
		p.left = l
	}

	if right != -1 {
		r := trees[right]
		p.right = r
		r.parent = p
	}
}

//  ALDS1_7_B: 二分木
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	n := scanInt(stdin)
	trees := initTrees(n)
	for i := 0; i < n; i++ {
		params := scanIntArray(stdin)
		setTree(params[0], params[1], params[2], trees)
	}

	for _, t := range trees {
		if t.nodeType() == "root" {
			t.calcDepth(0)
			t.calcHeight()
			break
		}
	}

	for _, t := range trees {
		fmt.Println(t.state())
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
