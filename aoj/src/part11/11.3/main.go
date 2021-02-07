package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const max = 1001

type dpLcs struct {
	c [][]int
}

func newDpLcs() *dpLcs {
	c := make([][]int, max, max)
	for i := range c {
		c[i] = make([]int, max, max)
	}
	return &dpLcs{
		c: c,
	}
}

func (dp *dpLcs) lcs(x, y string) int {
	val := 0
	for i, xi := range x {
		for j, yj := range y {
			if xi == yj {
				dp.c[i+1][j+1] = dp.c[i][j] + 1
			} else if dp.c[i][j+1] >= dp.c[i+1][j] {
				dp.c[i+1][j+1] = dp.c[i][j+1]
			} else {
				dp.c[i+1][j+1] = dp.c[i+1][j]
			}
			if val < dp.c[i+1][j+1] {
				val = dp.c[i+1][j+1]
			}
		}

	}

	return val
}

// ALDS1_10_C: LCS
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := scanInt(stdin)
	for i := 0; i < n; i++ {
		stdin.Scan()
		x := stdin.Text()
		stdin.Scan()
		y := stdin.Text()
		dp := newDpLcs()
		fmt.Fprintf(w, "%d\n", dp.lcs(x, y))
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
