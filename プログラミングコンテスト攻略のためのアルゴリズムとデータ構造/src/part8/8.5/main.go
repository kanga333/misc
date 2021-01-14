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
	left   *tree
	right  *tree
}

func (t *tree) appendChild(c *tree) {
	c.parent = t
	if t.left == nil {
		t.left = c
	} else {
		t.right = c
	}

}

func (t *tree) nodeType() string {
	if t.parent == nil && t.id != 0 {
		return "root"
	}
	if t.left == nil && t.right == nil {
		return "leaf"
	}

	return "internal node"
}

func (t *tree) preOrder() string {
	ret := fmt.Sprintf("%d", t.id)
	if t.left != nil {
		ret = fmt.Sprintf("%s %s", ret, t.left.preOrder())
	}
	if t.right != nil {
		ret = fmt.Sprintf("%s %s", ret, t.right.preOrder())
	}

	return ret
}

func (t *tree) inOrder() string {
	ret := fmt.Sprintf("%d", t.id)
	if t.left != nil {
		ret = fmt.Sprintf("%s %s", t.left.inOrder(), ret)
	}
	if t.right != nil {
		ret = fmt.Sprintf("%s %s", ret, t.right.inOrder())
	}

	return ret
}

func (t *tree) postOrder() string {
	ret := fmt.Sprintf("%d", t.id)
	if t.right != nil {
		ret = fmt.Sprintf("%s %s", t.right.postOrder(), ret)
	}
	if t.left != nil {
		ret = fmt.Sprintf("%s %s", t.left.postOrder(), ret)
	}

	return ret
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

//  ALDS1_7_C: 木の復元
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	n := scanInt(stdin)
	trees := initTrees(n + 1)
	set := make([]bool, n+1, n+1)
	pre := scanIntArray(stdin)
	cuurent := trees[pre[0]] // root
	set[cuurent.id] = true
	in := scanIntArray(stdin)
	j := 1
	for _, val := range in {

		target := trees[val]
		for {
			if set[target.id] {
				cuurent = target
				break
			}
			cuurent.appendChild(trees[pre[j]])
			cuurent = trees[pre[j]]
			set[cuurent.id] = true
			j++
		}
	}

	for _, t := range trees {
		if t.nodeType() == "root" {
			fmt.Printf("%s\n", t.postOrder())
			break
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
