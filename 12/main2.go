package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position_ struct {
	y int
	x int
}

func abs_(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func canBeWalked_(s Position_, e Position_, h []string, v [][]bool) bool {
	rowLen := len(h[0])
	colLen := len(h)

	if e.x < 0 || e.x >= rowLen || e.y < 0 || e.y >= colLen {
		return false
	}

	if abs_(e.y, s.y) > 1 || abs_(e.x, s.x) > 1 {
		return false
	}

	if v[e.y][e.x] {
		return false
	}

	if int(h[s.y][s.x])-int(h[e.y][e.x]) > 1 {
		return false
	}

	return true
}

func main() {
	file, _ := os.Open("12/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var s Position_
	sY := 0

	positions := make([]Position_, 0)
	stepMap := make([][]int, 0)
	visited := make([][]bool, 0)
	h := make([]string, 0)

	for fileScanner.Scan() {
		text := fileScanner.Text()

		if strings.ContainsRune(text, 'S') {
			text = strings.ReplaceAll(text, "S", "a")
		}

		if strings.ContainsRune(text, 'E') {
			s = Position_{
				x: strings.IndexRune(text, 'E'),
				y: sY,
			}
			text = strings.ReplaceAll(text, "E", "z")
			positions = append(positions, s)
		}

		stepRow := make([]int, len(text))
		visitedRow := make([]bool, len(text))

		h = append(h, text)
		stepMap = append(stepMap, stepRow)
		visited = append(visited, visitedRow)
		sY++
	}

	for step := 0; ; step++ {
		newPositions := map[string]bool{}

		for _, p := range positions {
			visited[p.y][p.x] = true
			if string(h[p.y][p.x]) == "a" {
				fmt.Println(step)
				return
			}

			left := Position_{
				y: p.y,
				x: p.x - 1,
			}
			if canBeWalked_(p, left, h, visited) {
				posString := strconv.Itoa(left.y) + " " + strconv.Itoa(left.x)
				newPositions[posString] = true
			} // left

			right := Position_{
				y: p.y,
				x: p.x + 1,
			}
			if canBeWalked_(p, right, h, visited) {
				posString := strconv.Itoa(right.y) + " " + strconv.Itoa(right.x)
				newPositions[posString] = true
			} // right

			up := Position_{
				y: p.y - 1,
				x: p.x,
			}
			if canBeWalked_(p, up, h, visited) {
				posString := strconv.Itoa(up.y) + " " + strconv.Itoa(up.x)
				newPositions[posString] = true
			} // up

			down := Position_{
				y: p.y + 1,
				x: p.x,
			}
			if canBeWalked_(p, down, h, visited) {
				posString := strconv.Itoa(down.y) + " " + strconv.Itoa(down.x)
				newPositions[posString] = true
			} // down
		}

		positions = make([]Position_, 0)
		for np, _ := range newPositions {
			spl := strings.Split(np, " ")
			y, _ := strconv.Atoi(spl[0])
			x, _ := strconv.Atoi(spl[1])

			positions = append(positions, Position_{
				y: y,
				x: x,
			})
		}
	}
}
