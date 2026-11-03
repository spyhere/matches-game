package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Matches Game
2 Players

Given N amount of matches, first player is grabbing as much as he wants, but not all of them. The next player
can grab <= then (previous player grabbed * 2). They continue grabbing matches until all matches have been taken.
Player who grabbed the last match has won.

At the end program should show the history of moves and display who has won.

For example:
Given 11 matches.
Moves: [4, 7]
P2 wins!
*/

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

	var moves []int
	initMatches := max(rand.Intn(maxMatches), minMatches)
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
			time.Sleep(time.Second * 1)
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
