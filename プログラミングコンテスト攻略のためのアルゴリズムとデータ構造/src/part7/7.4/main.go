package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const max = 10000

func countingSort(a, b []int16, k int) {
	c := make([]int, k, k)
	for i := range c {
		c[i] = 0
	}

	for _, v := range a {
		c[v]++
	}

	for i := range c {
		if i == 0 {
			continue
		}
		c[i] += c[i-1]
	}

	for i := len(a) - 1; i >= 0; i-- {
		ci := c[a[i]] - 1
		b[ci] = a[i]
		c[a[i]]--
	}
}

// ALDS1_6_A:計数ソート
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	n := scanInt(stdin)
	a := scanIntArray(stdin)
	b := make([]int16, n, n)

	countingSort(a, b, max+1)
	fmt.Println(strings.Trim(fmt.Sprint(b), "[]"))
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

func scanIntArray(sc *bufio.Scanner) []int16 {
	sc.Scan()
	t := sc.Text()
	ss := strings.Split(t, " ")
	is := []int16{}
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		is = append(is, int16(i))
	}
	return is
}
