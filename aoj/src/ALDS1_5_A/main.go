package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxWeight = 10000

// 配列Aの要素を組み合わせて配列mの要素の値を作れるか全探索で調べる
// n個の配列Aとq個の配列M
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan() // skip n
	a := scanIntArray(stdin)
	stdin.Scan() // skip q
	m := scanIntArray(stdin)
	for _, mi := range m {
		if exhaustiveSearch(a, mi) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func exhaustiveSearch(a []int, mi int) bool {
	if len(a) == 0 {
		return false
	}
	if a[0] == mi {
		return true
	}

	if a[0] < mi {
		return exhaustiveSearch(a[1:], mi-a[0]) || exhaustiveSearch(a[1:], mi)
	}

	return exhaustiveSearch(a[1:], mi)
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
