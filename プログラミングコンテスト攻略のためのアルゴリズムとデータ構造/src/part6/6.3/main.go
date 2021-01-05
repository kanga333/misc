package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type point struct {
	x float64
	y float64
}

func koch(p1 *point, p2 *point, n int) {
	if n == 0 {
		return
	}

	s := divideVector(1.0/3.0, p1, p2)
	t := divideVector(2.0/3.0, p1, p2)
	u := calcTriangle(s, t)

	koch(p1, s, n-1)
	fmt.Printf("%f %f\n", s.x, s.y)
	koch(s, u, n-1)
	fmt.Printf("%f %f\n", u.x, u.y)
	koch(u, t, n-1)
	fmt.Printf("%f %f\n", t.x, t.y)
	koch(t, p2, n-1)

	return
}

const th = math.Pi * 60.0 / 180.0

func calcTriangle(p1 *point, p2 *point) *point {
	return &point{
		x: (p2.x-p1.x)*math.Cos(th) - (p2.y-p1.y)*math.Sin(th) + p1.x,
		y: (p2.x-p1.x)*math.Sin(th) + (p2.y-p1.y)*math.Cos(th) + p1.y,
	}
}
func divideVector(ratio float64, p1 *point, p2 *point) *point {
	return &point{
		x: (1.0-ratio)*p1.x + ratio*p2.x,
		y: (1.0-ratio)*p1.y + ratio*p2.y,
	}
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	n := scanInt(stdin) // skip n
	p1 := point{
		x: 0.0,
		y: 0.0,
	}
	p2 := point{
		x: 100.0,
		y: 0.0,
	}
	fmt.Printf("%f %f\n", p1.x, p1.y)
	koch(&p1, &p2, n)
	fmt.Printf("%f %f\n", p2.x, p2.y)
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
