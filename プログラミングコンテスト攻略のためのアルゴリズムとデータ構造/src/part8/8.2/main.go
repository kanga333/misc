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
	parent *tree
	child  []int
}

func (t *tree) updateParent(p *tree) {
	t.parent = p
}

func (t *tree) updateChild(c []int, register []*tree) {
	t.child = c
	for _, cid := range c {
		cp := register[cid]
		cp.updateParent(t)
	}
}

func (t *tree) pID() int {
	if t.parent == nil {
		return -1
	}
	return t.parent.id
}

func (t *tree) depth() int {
	if t.parent == nil {
		return 0
	}
	return t.parent.depth() + 1
}

func (t *tree) childStr() string {
	s := []string{}
	for _, c := range t.child {
		s = append(s, strconv.Itoa(c))
	}
	return fmt.Sprintf("[%s]", strings.Join(s, ", "))
}

func (t *tree) nodeType() string {
	if t.parent == nil {
		return "root"
	}
	if len(t.child) == 0 {
		return "leaf"
	}
	return "internal node"
}

func initTrees(n int) []*tree {
	trees := make([]*tree, n, n)
	for i := 0; i < n; i++ {
		trees[i] = &tree{
			id:     i,
			parent: nil,
			child:  []int{},
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
		child := []int{}
		for i := 0; i < degree; i++ {
			child = append(child, params[i+2])
		}
		t.updateChild(child, trees)
	}

	for _, t := range trees {
		fmt.Printf("node %d: parent = %d, depth = %d, %s, %v\n",
			t.id, t.pID(), t.depth(), t.nodeType(), t.childStr())
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
