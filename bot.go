package main

import (
	"slices"
)

func newBot(matches int) bot {
	b := bot{make([]int, 10)}
	b.calculateFibNumbers(matches)
	return b
}

type bot struct {
	fibs []int
}

func (bo *bot) calculateFibNumbers(x int) {
	a := 1
	b := 1
	for b <= x {
		bo.fibs = append(bo.fibs, b)
		tmp := b
		b = b + a
		a = tmp
	}
}

func (bo bot) getMove(curX int) int {
	y := 0
	for _, it := range slices.Backward(bo.fibs) {
		if curX == it {
			return 1
		}
		if curX > it {
			y = curX % it
			break
		}
	}
	idx := len(bo.fibs) - 1
	for idx >= 0 && y != bo.fibs[idx] {
		if y > bo.fibs[idx] {
			y %= bo.fibs[idx]
		}
		idx--
	}
	return y
}
