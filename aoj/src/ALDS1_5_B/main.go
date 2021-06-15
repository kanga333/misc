package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const max = 1000000000

var count = 0

// ALDS1_5_B:マージソート
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	n := scanInt(stdin)
	a := scanIntArray(stdin)
	mergesort(a, 0, n)

	fmt.Println(strings.Trim(fmt.Sprint(a), "[]"))
	fmt.Println(count)
}

func mergesort(a []int, left, right int) {
	if left+1 >= right {
		return
	}
	mid := (left + right) / 2
	mergesort(a, left, mid)
	mergesort(a, mid, right)
	merge(a, left, mid, right)
	return
}

func merge(a []int, left, mid, right int) {
	ls := a[left:mid]
	l := make([]int, len(ls))
	copy(l, ls)
	rs := a[mid:right]
	r := make([]int, len(rs))
	copy(r, rs)

	l = append(l, max+1)
	r = append(r, max+1)

	li, ri := 0, 0
	for i := left; i < right; i++ {
		count++
		if l[li] <= r[ri] {
			a[i] = l[li]
			li++
		} else {
			a[i] = r[ri]
			ri++
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
