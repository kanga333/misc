package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type process struct {
	name string
	time int
}

type roundRobin struct {
	quantum int
	time    int
	pQueue  *processQueue
}

func (rb *roundRobin) run() {
	for {
		rb.calc()
		if rb.pQueue.isEmpty() {
			break
		}
	}
}

func (rb *roundRobin) calc() {
	p := rb.pQueue.dequeue()
	if p.time > rb.quantum {
		p.time = p.time - rb.quantum
		rb.time = rb.time + rb.quantum
		rb.pQueue.enqueue(p)
	} else {
		rb.time = rb.time + p.time
		fmt.Printf("%s %d\n", p.name, rb.time)
	}
	return
}

type processQueue struct {
	queue    []*process
	firstIdx int
	lastIdx  int
}

func (pq *processQueue) enqueue(p *process) {
	pq.queue[pq.lastIdx] = p
	pq.lastIdx++
	if pq.lastIdx >= len(pq.queue) {
		pq.lastIdx = 0
	}
}

func (pq *processQueue) dequeue() *process {
	p := pq.queue[pq.firstIdx]
	pq.queue[pq.firstIdx] = nil
	pq.firstIdx++
	if pq.firstIdx >= len(pq.queue) {
		pq.firstIdx = 0
	}

	return p
}

func (pq *processQueue) isEmpty() bool {
	return pq.queue[pq.firstIdx] == nil
}

func initProcessQueue(l int) *processQueue {
	return &processQueue{
		queue:    make([]*process, l),
		firstIdx: 0,
		lastIdx:  0,
	}
}

func initRoudRobin(line string) *roundRobin {
	in := strings.Split(line, " ")
	l, err := strconv.Atoi(in[0])
	if err != nil {
		panic(err)
	}
	pq := initProcessQueue(l)

	q, err := strconv.Atoi(in[1])
	if err != nil {
		panic(err)
	}
	return &roundRobin{
		quantum: q,
		time:    0,
		pQueue:  pq,
	}
}

func initProcess(line string) *process {
	in := strings.Split(line, " ")
	t, err := strconv.Atoi(in[1])
	if err != nil {
		panic(err)
	}
	return &process{
		name: in[0],
		time: t,
	}
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	rb := initRoudRobin(stdin.Text())

	for stdin.Scan() {
		p := initProcess(stdin.Text())
		rb.pQueue.enqueue(p)
	}
	rb.run()
}
