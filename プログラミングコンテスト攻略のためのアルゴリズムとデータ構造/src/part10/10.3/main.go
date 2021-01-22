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

func swap(a []int, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func maxHeapity(h []int, i int) {
	l := left(len(h), i)
	r := right(len(h), i)

	largest := i
	if l != 0 && h[l] > h[largest] {
		largest = l
	}
	if r != 0 && h[r] > h[largest] {
		largest = r
	}

	if largest != i {
		swap(h, largest, i)
		maxHeapity(h, largest)
	}

}

func left(size, i int) int {
	left := 2 * i
	if left < size {
		return left
	}
	return 0
}

func right(size, i int) int {
	right := (2 * i) + 1
	if right < size {
		return right
	}
	return 0
}

// ALDS1_9_B: 最大ヒープ
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	stdin.Scan() // skip n
	a := scanIntArray(stdin)
	heap := []int{0}
	heap = append(heap, a...)

	for i := len(heap); i >= 1; i-- {
		maxHeapity(heap, i)
	}
	for _, val := range heap[1:] {

		fmt.Fprintf(w, " %d", val)
	}
	fmt.Fprintf(w, "\n")
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
