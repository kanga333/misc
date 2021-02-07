package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	//n := stdin.Text()
	stdin.Scan()
	s := stdin.Text()
	stdin.Scan()
	//q := stdin.Text()
	stdin.Scan()
	t := stdin.Text()

	cnt := 0
	for _, tv := range strings.Split(t, " ") {
		for _, sv := range strings.Split(s, " ") {
			if tv == sv {
				cnt++
				break
			}
		}
	}
	fmt.Println(cnt)

}
