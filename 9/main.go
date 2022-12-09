package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
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

	tX := 0
	tY := 0

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

			if abs(tX-hX) > 1 || abs(tY-hY) > 1 {
				if tX == hX {
					tY = (tY + hY) / 2
				} else if tY == hY {
					tX = (tX + hX) / 2
				} else {
					if hX > tX {
						tX++
					} else {
						tX--
					}
					if hY > tY {
						tY++
					} else {
						tY--
					}
				}
			}

			posString := strconv.Itoa(tX) + " " + strconv.Itoa(tY)
			visited[posString] = true
		}
	}
	print(len(visited))
}
