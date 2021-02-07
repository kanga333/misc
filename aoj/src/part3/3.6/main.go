package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var cnt int

func main() {
	rand.Seed(time.Now().UnixNano())
	cnt = 0
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	n := scanInt(stdin)

	a := scanIntArray(stdin)
	m, g := shellSort(a, n)
	fmt.Println(m)
	fmt.Println(strings.Trim(fmt.Sprint(g), "[]"))
	fmt.Println(cnt)
	for _, v := range a {
		fmt.Println(v)
	}

}

func shellSort(a []int, n int) (int, []int) {

	m, g := calcInterval(n)

	for _, v := range g {
		insertSort(a, n, v)
	}
	return m, g
}

func calcInterval(n int) (int, []int) {
	rg := []int{}
	m := 0
	for v := 1; v <= n; v = 3*v + 1 {
		m++
		rg = append(rg, v)
	}
	g := []int{}
	for i := len(rg) - 1; i >= 0; i-- {
		g = append(g, rg[i])
	}
	return m, g
}

func insertSort(a []int, n, g int) {
	for i := g; i < n; i++ {
		v := a[i]
		j := i - g
		for j >= 0 && a[j] > v {
			a[j+g] = a[j]
			j = j - g
			cnt++
		}
		a[j+g] = v
	}
}

func scanIntArray(stdin *bufio.Scanner) []int {
	is := []int{}
	for stdin.Scan() {
		i := scanInt(stdin)
		is = append(is, i)
	}
	return is
}

func scanInt(sc *bufio.Scanner) int {
	t := sc.Text()
	i, err := strconv.Atoi(t)
	if err != nil {
		panic(err)
	}
	return i
}
