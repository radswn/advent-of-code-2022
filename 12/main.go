package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	y int
	x int
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func canBeWalked(s Position, e Position, h []string, v [][]bool) bool {
	rowLen := len(h[0])
	colLen := len(h)

	if e.x < 0 || e.x >= rowLen || e.y < 0 || e.y >= colLen {
		return false
	}

	if abs(e.y, s.y) > 1 || abs(e.x, s.x) > 1 {
		return false
	}

	if v[e.y][e.x] {
		return false
	}

	if int(h[e.y][e.x])-int(h[s.y][s.x]) > 1 {
		return false
	}

	return true
}

func main() {
	file, _ := os.Open("12/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var e Position
	eY := 0

	var s Position
	sY := 0

	positions := make([]Position, 0)
	stepMap := make([][]int, 0)
	visited := make([][]bool, 0)
	h := make([]string, 0)

	for fileScanner.Scan() {
		text := fileScanner.Text()

		if strings.ContainsRune(text, 'S') {
			s = Position{
				x: strings.IndexRune(text, 'S'),
				y: sY,
			}
			text = strings.ReplaceAll(text, "S", "a")
			positions = append(positions, s)
		}

		if strings.ContainsRune(text, 'E') {
			e = Position{
				x: strings.IndexRune(text, 'E'),
				y: eY,
			}
			text = strings.ReplaceAll(text, "E", "z")
		}

		stepRow := make([]int, len(text))
		visitedRow := make([]bool, len(text))

		h = append(h, text)
		stepMap = append(stepMap, stepRow)
		visited = append(visited, visitedRow)

		eY++
		sY++
	}

	for step := 0; !visited[e.y][e.x]; step++ {
		newPositions := map[string]bool{}

		for _, p := range positions {
			visited[p.y][p.x] = true
			stepMap[p.y][p.x] = step

			left := Position{
				y: p.y,
				x: p.x - 1,
			}
			if canBeWalked(p, left, h, visited) {
				posString := strconv.Itoa(left.y) + " " + strconv.Itoa(left.x)
				newPositions[posString] = true
			} // left

			right := Position{
				y: p.y,
				x: p.x + 1,
			}
			if canBeWalked(p, right, h, visited) {
				posString := strconv.Itoa(right.y) + " " + strconv.Itoa(right.x)
				newPositions[posString] = true
			} // right

			up := Position{
				y: p.y - 1,
				x: p.x,
			}
			if canBeWalked(p, up, h, visited) {
				posString := strconv.Itoa(up.y) + " " + strconv.Itoa(up.x)
				newPositions[posString] = true
			} // up

			down := Position{
				y: p.y + 1,
				x: p.x,
			}
			if canBeWalked(p, down, h, visited) {
				posString := strconv.Itoa(down.y) + " " + strconv.Itoa(down.x)
				newPositions[posString] = true
			} // down
		}

		positions = make([]Position, 0)
		for np, _ := range newPositions {
			spl := strings.Split(np, " ")
			y, _ := strconv.Atoi(spl[0])
			x, _ := strconv.Atoi(spl[1])

			positions = append(positions, Position{
				y: y,
				x: x,
			})
		}
	}

	fmt.Println(stepMap[e.y][e.x])
}
