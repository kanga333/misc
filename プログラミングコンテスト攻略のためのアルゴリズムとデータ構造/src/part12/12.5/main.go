package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	users []*user
)

type user struct {
	id      int32
	friends map[int32]struct{}
	group   int32
}

func (u *user) grouping(g int32) {
	if u.group != 0 {
		return
	}
	u.group = g
	for f := range u.friends {
		users[f].grouping(g)
	}
}

func (u *user) reachable(d *user) bool {
	if u.group == 0 {
		u.grouping(u.id + 1)
	}
	return u.group == d.group
}

// ALDS1_11_D: 連結成分
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	a := scanIntArray(stdin)
	userN := a[0]
	relationN := a[1]

	users = make([]*user, userN, userN)
	for i := 0; i < userN; i++ {
		f := make(map[int32]struct{})
		users[i] = &user{
			id:      int32(i),
			friends: f,
		}
	}

	for i := 0; i < relationN; i++ {
		a := scanIntArray(stdin)
		src := users[a[0]]
		dst := users[a[1]]
		src.friends[dst.id] = struct{}{}
		dst.friends[src.id] = struct{}{}
	}

	q := scanInt(stdin)
	for i := 0; i < q; i++ {
		a := scanIntArray(stdin)

		src := users[a[0]]
		dst := users[a[1]]

		ok := src.reachable(dst)
		if ok {
			fmt.Fprintln(w, "yes")
		} else {
			fmt.Fprintln(w, "no")
		}

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
