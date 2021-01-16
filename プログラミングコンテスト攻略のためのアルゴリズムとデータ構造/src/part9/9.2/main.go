package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var preIndex = 0

type tree struct {
	val    int
	parent *tree
	left   *tree
	right  *tree
}

func (t *tree) insert(z *tree) {
	var p *tree
	x := t
	for x != nil {
		p = x
		if z.val < x.val {
			x = x.left
			continue
		}
		x = x.right
	}
	z.parent = p

	if z.val < p.val {
		p.left = z
		return
	}
	p.right = z

}

func (t *tree) nodeType() string {
	if t.parent == nil && t.val != 0 {
		return "root"
	}
	if t.left == nil && t.right == nil {
		return "leaf"
	}

	return "internal node"
}

func (t *tree) preOrder() string {
	ret := fmt.Sprintf("%d", t.val)
	if t.left != nil {
		ret = fmt.Sprintf("%s %s", ret, t.left.preOrder())
	}
	if t.right != nil {
		ret = fmt.Sprintf("%s %s", ret, t.right.preOrder())
	}

	return ret
}

func (t *tree) inOrder() string {
	ret := fmt.Sprintf("%d", t.val)
	if t.left != nil {
		ret = fmt.Sprintf("%s %s", t.left.inOrder(), ret)
	}
	if t.right != nil {
		ret = fmt.Sprintf("%s %s", ret, t.right.inOrder())
	}

	return ret
}

func (t *tree) postOrder() string {
	ret := fmt.Sprintf("%d", t.val)
	if t.right != nil {
		ret = fmt.Sprintf("%s %s", t.right.postOrder(), ret)
	}
	if t.left != nil {
		ret = fmt.Sprintf("%s %s", t.left.postOrder(), ret)
	}

	return ret
}

func parseTree(s string) *tree {
	a := strings.Split(s, " ")
	val, err := strconv.Atoi(a[1])
	if err != nil {
		panic(err)
	}
	return &tree{
		val: val,
	}
}

//  ALDS1_8_A: 二分木探索: 挿入
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	stdin.Scan() //skip scan n

	var p *tree
	for stdin.Scan() {
		txt := stdin.Text()

		switch {
		case strings.HasPrefix(txt, "insert"):
			t := parseTree(txt)
			if p == nil {
				p = t
				continue
			}
			p.insert(t)
		case strings.HasPrefix(txt, "print"):
			fmt.Fprintln(w, "", p.inOrder())
			fmt.Fprintln(w, "", p.preOrder())
		default:
			panic(fmt.Sprintf("Unexpected case: %s", txt))
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
