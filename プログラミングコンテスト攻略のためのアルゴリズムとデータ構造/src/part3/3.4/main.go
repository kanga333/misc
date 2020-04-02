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
	sort := scanArray(stdin)

	c := 0

	for i := 0; i < n; i++ {
		minj := i
		for j := i; j < n; j++ {
			if sort[j] < sort[minj] {
				minj = j
			}
		}
		if minj != i {
			c++
		}
		tmp := sort[minj]
		sort[minj] = sort[i]
		sort[i] = tmp
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

func scanArray(sc *bufio.Scanner) []string {
	t := sc.Text()
	ss := strings.Split(t, " ")

	return ss
}
