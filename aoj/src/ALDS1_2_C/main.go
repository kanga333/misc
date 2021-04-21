package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	n := scanInt(stdin)
	stdin.Scan()
	bs := make([]string, n)
	ss := make([]string, n)
	bs = scanArray(stdin)
	copy(ss, bs)

	bubbleSort(n, bs)
	selectSort(n, ss)
	if reflect.DeepEqual(bs, ss) {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}

func bubbleSort(n int, a []string) {
	flag := true
	for flag {
		flag = false
		for j := n - 1; j >= 1; j-- {
			if intCard(a[j]) < intCard(a[j-1]) {
				tmp := a[j-1]
				a[j-1] = a[j]
				a[j] = tmp
				flag = true
			}
		}
	}
	fmt.Println(strings.Trim(fmt.Sprint(a), "[]"))
	fmt.Println("Stable")
}

func selectSort(n int, a []string) {
	for i := 0; i < n; i++ {
		minj := i
		for j := i; j < n; j++ {
			if intCard(a[j]) < intCard(a[minj]) {
				minj = j
			}
		}

		tmp := a[minj]
		a[minj] = a[i]
		a[i] = tmp
	}
	fmt.Println(strings.Trim(fmt.Sprint(a), "[]"))
}

func intCard(s string) int {
	return int([]rune(s)[1] - '0')
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
