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

// ALDS1_6_B:パーティション
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	n := scanInt(stdin)
	a := scanIntArray(stdin)
	q := partition(a, 0, n-1)
	str := make([]string, n, n)
	for i, val := range a {
		if i == q {
			str[i] = fmt.Sprintf("[%d]", val)
			continue
		}
		str[i] = fmt.Sprintf("%d", val)
	}

	fmt.Println(strings.Join(str, " "))
}

func partition(a []int, p, r int) int {
	x := a[r]
	i := p
	for j := 0; j < r; j++ {
		if a[j] <= x {
			swap(a, i, j)
			i++
		}
	}
	swap(a, i, r)
	return i
}

func swap(a []int, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp

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
