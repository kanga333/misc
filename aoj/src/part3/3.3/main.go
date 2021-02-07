package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	n := scanInt(stdin)
	stdin.Scan()
	sort := scanIntArray(stdin)

	flag := true
	c := 0

	for flag {
		flag = false
		for j := n - 1; j >= 1; j-- {
			if sort[j] < sort[j-1] {
				tmp := sort[j-1]
				sort[j-1] = sort[j]
				sort[j] = tmp
				flag = true
				c++
			}
		}
	}
	fmt.Println(strings.Trim(fmt.Sprint(sort), "[]"))
	fmt.Println(c)
}

func scanInt(sc *bufio.Scanner) int {
	t := sc.Text()
	i, err := strconv.Atoi(t)
	if err != nil {
		panic(err)
	}
	return i
}

func scanIntArray(sc *bufio.Scanner) []int {
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
