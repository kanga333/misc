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

	fmt.Println(strings.Trim(fmt.Sprint(sort), "[]"))

	for i := 1; i < n; i++ {
		v := sort[i]
		j := i - 1
		for j >= 0 && sort[j] > v {
			sort[j+1] = sort[j]
			j--
		}
		sort[j+1] = v
		fmt.Println(strings.Trim(fmt.Sprint(sort), "[]"))
	}
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
