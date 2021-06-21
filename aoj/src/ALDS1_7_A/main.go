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
	id        int
	depth     int
	parent    *tree
	child     *tree
	lastChild *tree
	sibling   *tree
}

func (t *tree) appendChild(c *tree) {
	c.parent = t

	if t.child == nil {
		t.child = c
		t.lastChild = c
		return
	}

	t.lastChild.sibling = c
	t.lastChild = c
}

func (t *tree) pID() int {
	if t.parent == nil {
		return -1
	}
	return t.parent.id
}

func (t *tree) childStr() string {
	c := t.child
	if c == nil {
		return "[]"
	}

	s := []string{strconv.Itoa(c.id)}
	for {
		c = c.sibling
		if c == nil {
			break
		}
		s = append(s, strconv.Itoa(c.id))
	}
	return fmt.Sprintf("[%s]", strings.Join(s, ", "))
}

func (t *tree) nodeType() string {
	if t.parent == nil {
		return "root"
	}
	if t.child == nil {
		return "leaf"
	}
	return "internal node"
}

func (t *tree) calcDepth(depth int) {
	t.depth = depth
	if t.child != nil {
		t.child.calcDepth(depth + 1)
	}
	if t.sibling != nil {
		t.sibling.calcDepth(depth)
	}
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

//  ALDS1_7_A: 根付き木
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	n := scanInt(stdin)
	trees := initTrees(n)
	for i := 0; i < n; i++ {
		params := scanIntArray(stdin)
		id := params[0]
		t := trees[id]
		degree := params[1]
		for i := 0; i < degree; i++ {
			c := trees[params[i+2]]
			t.appendChild(c)
		}
	}

	for _, t := range trees {
		if t.nodeType() == "root" {
			t.calcDepth(0)
			break
		}
	}

	for _, t := range trees {
		fmt.Printf("node %d: parent = %d, depth = %d, %s, %v\n",
			t.id, t.pID(), t.depth, t.nodeType(), t.childStr())
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
