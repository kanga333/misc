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

func printArray(a []int) {
	n := len(a)

	fmt.Fprintf(w, "%d", a[0])
	for i := 1; i < n; i++ {
		fmt.Fprintf(w, " %d", a[i])
	}
	fmt.Fprintln(w, "")
}

func insertSort(a []int) {
	n := len(a)

	for i := 1; i < n; i++ {
		printArray(a)
		v := a[i]
		j := i - 1
		for ; j >= 0 && a[j] > v; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = v
	}
	printArray(a)
}

// ALDS1_1_A: 挿入ソート
func main() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	defer w.Flush()

	n := nextInt()

	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	insertSort(a)

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
