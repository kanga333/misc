package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const max = 1000000000

type card struct {
	rawText string
	number  int
}

func newCard(s string) *card {
	splited := strings.Split(s, " ")
	number, err := strconv.Atoi(splited[1])
	if err != nil {
		panic(err)
	}
	return &card{
		rawText: s,
		number:  number,
	}
}

// ALDS1_6_C:クイックソート
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	n := scanInt(stdin)
	cards := make([]*card, n, n)
	for i := 0; i < n; i++ {
		stdin.Scan()
		t := stdin.Text()
		cards[i] = newCard(t)
	}
	cards2 := make([]*card, n)
	copy(cards2, cards)
	quickSort(cards, 0, n-1)
	mergesort(cards2, 0, n)
	if reflect.DeepEqual(cards, cards2) {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
	for _, c := range cards {
		fmt.Println(c.rawText)
	}
}

func quickSort(a []*card, p, r int) {
	if p < r {
		q := partition(a, p, r)
		quickSort(a, p, q-1)
		quickSort(a, q+1, r)
	}

}

func partition(a []*card, p, r int) int {
	x := a[r].number
	i := p
	for j := p; j < r; j++ {
		if a[j].number <= x {
			swap(a, i, j)
			i++

		}
	}
	swap(a, i, r)
	return i

}

func swap(a []*card, i, j int) {
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

func mergesort(a []*card, left, right int) {
	if left+1 >= right {
		return
	}
	mid := (left + right) / 2
	mergesort(a, left, mid)
	mergesort(a, mid, right)
	merge(a, left, mid, right)
	return
}

func merge(a []*card, left, mid, right int) {
	ls := a[left:mid]
	l := make([]*card, len(ls))
	copy(l, ls)
	rs := a[mid:right]
	r := make([]*card, len(rs))
	copy(r, rs)

	l = append(l, &card{number: max + 1})
	r = append(r, &card{number: max + 1})

	li, ri := 0, 0
	for i := left; i < right; i++ {
		if l[li].number <= r[ri].number {
			a[i] = l[li]
			li++
		} else {
			a[i] = r[ri]
			ri++
		}
	}

}
