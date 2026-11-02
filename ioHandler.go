package main

import (
	"bufio"
	"os"
	"strconv"
)

func newIOHandler() IOHandler {
	return IOHandler{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

type IOHandler struct {
	scanner *bufio.Scanner
}

func (ioh IOHandler) getUserName() string {
	if ioh.scanner.Scan() {
		return ioh.scanner.Text()
	}
	return ""
}

func (ioh IOHandler) getMove() int {
	if ioh.scanner.Scan() {
		line := ioh.scanner.Text()
		integer, err := strconv.Atoi(line)
		if err != nil {
			return -1
		}
		return integer
	}
	return 0
}
