package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const maxPrice = 1000000000

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ALDS1_1_D: Maximum Profit
func main() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := nextInt()

	bottom := nextInt()
	profit := maxPrice * -1
	for i := 1; i < n; i++ {
		price := nextInt()
		profit = max(profit, price-bottom)
		bottom = min(bottom, price)
	}
	fmt.Fprintln(w, profit)
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
