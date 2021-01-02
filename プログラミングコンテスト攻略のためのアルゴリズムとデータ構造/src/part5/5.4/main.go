package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type dict struct {
	size  int
	table []*int
}

func (d *dict) insert(key int) {
	for i := 0; i < d.size; i++ {
		k := d.calcKey(key, i)
		val := d.table[k]
		if val == nil {
			//fmt.Println(k, key)
			d.table[k] = &key
			return
		}
	}
}

func (d *dict) find(key int) bool {
	for i := 0; i < d.size; i++ {
		k := d.calcKey(key, i)
		val := d.table[k]
		if val == nil {
			return false
		}
		if *val == key {
			//fmt.Println(*val, k, key)
			return true
		}

	}
	return false
}

func (d *dict) calcKey(key int, i int) int {
	k1 := key % d.size
	k2 := 1 + (key % (d.size - 1))
	return (k1 + i*k2) % d.size
}

func newDict(n int) *dict {
	table := make([]*int, n, n)
	return &dict{
		size:  n,
		table: table,
	}
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	d := newDict(1000000)

	for stdin.Scan() {
		t := stdin.Text()
		cmd, key := parseCmd(t)
		switch cmd {
		case "insert":
			d.insert(key)
		case "find":
			ok := d.find(key)
			if ok {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		default:
			panic("Unexpected case")
		}
	}
}

func parseCmd(s string) (string, int) {
	ss := strings.Split(s, " ")
	return ss[0], convertKey(ss[1])
}

func convertKey(s string) int {
	var i int
	p := 1
	for _, b := range []byte(s) {
		val := castByte(b)
		i = i + p*val
		p = p * 5
	}
	//fmt.Println(s, i)
	return i
}

func castByte(b byte) int {
	switch b {
	case []byte("A")[0]:
		return 1
	case []byte("C")[0]:
		return 2
	case []byte("G")[0]:
		return 3
	case []byte("T")[0]:
		return 4
	default:
		return 0
	}
}
