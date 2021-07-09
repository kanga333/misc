package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	min = -1
	k   = 2000001
)

type heap struct {
	a    []int
	size int
}

func newHeap() *heap {
	a := make([]int, k, k)
	return &heap{
		a:    a,
		size: 0,
	}
}

func (h *heap) swap(i, j int) {
	tmp := h.a[i]
	h.a[i] = h.a[j]
	h.a[j] = tmp
}

func (h *heap) maxHeapity(i int) {
	l := left(h.size, i)
	r := right(h.size, i)

	largest := i
	if l != 0 && h.a[l] > h.a[largest] {
		largest = l
	}
	if r != 0 && h.a[r] > h.a[largest] {
		largest = r
	}

	if largest != i {
		h.swap(largest, i)
		h.maxHeapity(largest)
	}
}

func (h *heap) insert(key int) {
	h.size++
	i := h.size
	h.a[i] = key

	for i > 1 && h.a[parent(h.size, i)] < h.a[i] {
		h.swap(i, parent(h.size, i))
		i = parent(h.size, i)
	}
}

func (h *heap) extractMax() int {
	max := h.a[1]
	h.a[1] = h.a[h.size]
	h.size--
	h.maxHeapity(1)
	return max
}

func parent(size, i int) int {
	return i / 2
}

func left(size, i int) int {
	left := 2 * i
	if left <= size {
		return left
	}
	return 0
}

func right(size, i int) int {
	right := (2 * i) + 1
	if right <= size {
		return right
	}
	return 0
}

// ALDS1_9_C: 優先度付きキュー
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	h := newHeap()

	for stdin.Scan() {
		txt := stdin.Text()

		switch {
		case strings.HasPrefix(txt, "insert"):
			val := parseVal(txt)
			h.insert(val)
		case strings.HasPrefix(txt, "extract"):
			max := h.extractMax()
			fmt.Fprintf(w, "%d\n", max)
		default:
			break
		}
	}
}

func parseVal(s string) int {
	a := strings.Split(s, " ")
	val, err := strconv.Atoi(a[1])
	if err != nil {
		panic(err)
	}
	return val
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
