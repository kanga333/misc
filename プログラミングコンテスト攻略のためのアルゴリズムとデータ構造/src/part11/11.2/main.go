package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var dp []int

func init() {
	dp = make([]int, 45, 45)
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	if dp[n] != 0 {
		return dp[n]
	}

	val := fib(n-2) + fib(n-1)
	dp[n] = val
	return val
}

// ALDS1_10_A: フィボナッチ数列
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := scanInt(stdin)
	fmt.Fprintf(w, "%d\n", fib(n))
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
