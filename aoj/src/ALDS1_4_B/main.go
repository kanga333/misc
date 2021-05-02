package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer([]byte{}, math.MaxInt64)
	stdin.Scan()
	n := stdin.Text()
	nv, _ := strconv.Atoi(n)
	stdin.Scan()
	s := stdin.Text()
	ss := strings.Split(s, " ")
	stdin.Scan()
	//q := stdin.Text()
	stdin.Scan()
	t := stdin.Text()

	cnt := 0
	for _, tv := range strings.Split(t, " ") {
		if binarySearch(tv, nv, ss) {
			cnt++
		}
	}
	fmt.Println(cnt)

}

func binarySearch(target string, length int, a []string) bool {
	ti, _ := strconv.Atoi(target)
	for {
		i := length / 2
		selected := a[i]
		si, _ := strconv.Atoi(selected)
		if si == ti {
			return true
		} else if si < ti {
			a = a[i:]
			length = length - i
		} else {
			length = i
		}
		if i == 0 {
			break
		}
	}

	return false
}
