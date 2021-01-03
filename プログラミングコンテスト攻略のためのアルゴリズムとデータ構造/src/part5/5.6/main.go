package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxWeight = 10000

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	t := stdin.Text()
	//n: 荷物数, k: トラック数
	n, k := parseCondition(t)

	// w: n個の荷物の重さ
	w := make([]int, n, n)
	i := 0
	for stdin.Scan() {
		t := stdin.Text()
		wi, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		w[i] = wi
		i++
	}

	ii := binarySearch(maxWeight*n, func(i int) bool {
		//fmt.Println(w, k, i)
		return canBeAllocated(w, k, i)

	})
	fmt.Println(ii)
}

func binarySearch(right int, ORHigher func(int) bool) int {
	left := 0
	for left < right {

		mid := (right + left) / 2
		//fmt.Println(left, right, mid)
		if ORHigher(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func parseCondition(s string) (int, int) {
	ss := strings.Split(s, " ")
	n, err := strconv.Atoi(ss[0])
	if err != nil {
		panic(err)
	}

	k, err := strconv.Atoi(ss[1])
	if err != nil {
		panic(err)
	}

	return n, k
}

func canBeAllocated(w []int, k, p int) bool {
	load := 0
	allocated := 0
	for _, wi := range w {
		//fmt.Printf("load: %v, allocated: %v, wi: %v, k: %v, p: %v\n", load, allocated, wi, k, p)
		if wi > p {
			return false
		}
		if load+wi > p {
			allocated++
			if allocated > k {
				return false
			}
			load = 0
		}
		load += wi
	}
	if allocated >= k {
		return false
	}
	return true
}
