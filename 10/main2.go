package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func drawPixel(x int, signal int) {
	if isSpriteVisible(x, signal) {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if shouldBreakLine(signal) {
		fmt.Printf("\n")
	}
}

func shouldBreakLine(signal int) bool {
	interestingSignals := []int{40, 80, 120, 160, 200, 240}
	for _, e := range interestingSignals {
		if e == signal {
			return true
		}
	}
	return false
}

func isSpriteVisible(x int, signal int) bool {
	return signal%40 >= x && signal%40 <= x+2
}

func main() {
	file, _ := os.Open("10/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	cycle := 1
	X := 1

	for fileScanner.Scan() {
		text := fileScanner.Text()
		spl := strings.Split(text, " ")

		if len(spl) == 2 {
			add, _ := strconv.Atoi(spl[1])

			drawPixel(X, cycle)
			cycle++

			drawPixel(X, cycle)
			cycle++
			X += add
		} else {
			drawPixel(X, cycle)
			cycle++
		}
	}

}
