package main

import (
	"fmt"
	"strings"
)

func printMoves(players []string, total int, moves []int) {
	playerIdx := 0
	strBuilder := strings.Builder{}
	remainder := total
	for _, it := range moves {
		remainder -= it
		fmt.Fprintf(&strBuilder, "%s %d(%d)", players[playerIdx], it, remainder)
		if (playerIdx == 0) && (remainder != 0) {
			strBuilder.WriteString(" - ")
		} else {
			strBuilder.WriteString("\n")
		}
		playerIdx = (playerIdx + 1) % len(players)
	}
	fmt.Printf("Moves:\n%s\n", strBuilder.String())
}
