package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	minMatches = 11
	maxMatches = 30
	computer   = "Computer"
)

func main() {
	var players []string
	ioHandler := newIOHandler()

	fmt.Print("Enter your name: ")
	players = append(players, ioHandler.getUserName(), computer)

	fmt.Println("Shuffle...")
	rand.Shuffle(2, func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})
	time.Sleep(time.Second)
	initMatches := max(rand.Intn(maxMatches), minMatches)
	fmt.Printf("%d matches for this game. First move goes to %s\n", initMatches, players[0])

	var moves []int
	bot := newBot(initMatches)

	matches := initMatches
	playerIdx := 0
	prevMove := matches
	for {
		canTake := min(prevMove*2, min(matches, initMatches-1))
		if canTake == matches {
			moves = append(moves, canTake)
			break
		}
		var move int
		if players[playerIdx] == computer {
			move = min(bot.getMove(matches), canTake)
			fmt.Printf("Available %d matches (%d left). ...Computer grabs %d\n", canTake, matches, move)
			time.Sleep(time.Second)
		} else {
			fmt.Printf("Available %d matches (%d left). Your move: ", canTake, matches)
			move = ioHandler.getMove()
		}
		for (move <= 0) || (move > canTake) {
			fmt.Printf("\nInput correct number of matches, you can take from %d to %d: ", 1, canTake)
			move = ioHandler.getMove()
		}
		matches -= move
		moves = append(moves, move)
		if matches == 0 {
			break
		}

		prevMove = move
		playerIdx = (playerIdx + 1) % len(players)
	}

	fmt.Printf("\nGame over\n%s has won!\n", players[playerIdx])
	fmt.Printf("Initial amount of matches: %d\n", initMatches)
	printMoves(players, initMatches, moves)
}
