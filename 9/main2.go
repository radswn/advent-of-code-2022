package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func abs_(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, _ := os.Open("9/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	hX := 0
	hY := 0

	tX := make([]int, 9)
	tY := make([]int, 9)

	visited := map[string]bool{
		"0 0": true,
	}

	for fileScanner.Scan() {
		lineSplit := strings.Split(fileScanner.Text(), " ")
		direction := lineSplit[0]
		length, _ := strconv.Atoi(lineSplit[1])

		for i := 0; i < length; i++ {
			switch direction {
			case "R":
				hX++
			case "L":
				hX--
			case "U":
				hY++
			case "D":
				hY--
			}

			prevTX := hX
			prevTY := hY

			for k := 0; k < 9; k++ {
				if abs_(tX[k]-prevTX) > 1 || abs_(tY[k]-prevTY) > 1 {
					if tX[k] == prevTX {
						tY[k] = (tY[k] + prevTY) / 2
					} else if tY[k] == prevTY {
						tX[k] = (tX[k] + prevTX) / 2
					} else {
						if prevTX > tX[k] {
							tX[k]++
						} else {
							tX[k]--
						}
						if prevTY > tY[k] {
							tY[k]++
						} else {
							tY[k]--
						}
					}
				}
				prevTX = tX[k]
				prevTY = tY[k]
			}

			posString := strconv.Itoa(tX[8]) + " " + strconv.Itoa(tY[8])
			visited[posString] = true
		}
	}
	print(len(visited))
}
