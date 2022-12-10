package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSignalInteresting(signal int) bool {
	interestingSignals := []int{20, 60, 100, 140, 180, 220}
	for _, e := range interestingSignals {
		if e == signal {
			return true
		}
	}
	return false
}

func main() {
	file, _ := os.Open("10/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	cycle := 1
	X := 1
	signalSum := 0

	for fileScanner.Scan() {
		text := fileScanner.Text()
		spl := strings.Split(text, " ")

		if len(spl) == 2 {
			add, _ := strconv.Atoi(spl[1])

			if isSignalInteresting(cycle) {
				signalSum += cycle * X
				fmt.Printf("X=%d at cycle=%d\n", X, cycle)
			}

			cycle++

			if isSignalInteresting(cycle) {
				signalSum += cycle * X
				fmt.Printf("X=%d at cycle=%d\n", X, cycle)
			}

			cycle++
			X += add
		} else {
			if isSignalInteresting(cycle) {
				signalSum += cycle * X
				fmt.Printf("X=%d at cycle=%d\n", X, cycle)
			}

			cycle++
		}
	}

	fmt.Println(signalSum)
}
