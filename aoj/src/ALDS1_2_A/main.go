package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var w = bufio.NewWriter(os.Stdout)
var c int

func bubbleSort(a []int) {
	n := len(a)
	flag := true
	c = 0

	for flag {
		flag = false
		for j := n - 1; j >= 1; j-- {
			if a[j] < a[j-1] {
				tmp := a[j-1]
				a[j-1] = a[j]
				a[j] = tmp
				flag = true
				c++
			}
		}
	}
}

func printArray(a []int) {
	n := len(a)

	fmt.Fprintf(w, "%d", a[0])
	for i := 1; i < n; i++ {
		fmt.Fprintf(w, " %d", a[i])
	}
	fmt.Fprintln(w, "")
}

func main() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	defer w.Flush()

	n := nextInt()
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	bubbleSort(a)
	printArray(a)
	fmt.Fprintln(w, c)
}

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	s := next()
	i, _ := strconv.Atoi(s)
	return i
}
