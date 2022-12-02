package main

import (
	"bufio"
	"fmt"
	"os"
)

var picks = map[string]string{
	"A X": "A Z",
	"A Y": "A X",
	"A Z": "A Y",
	"B X": "B X",
	"B Y": "B Y",
	"B Z": "B Z",
	"C X": "C Y",
	"C Y": "C Z",
	"C Z": "C X",
}

var points2 = map[string]int{
	"A X": 3 + 1,
	"A Y": 6 + 2,
	"A Z": 0 + 3,
	"B X": 0 + 1,
	"B Y": 3 + 2,
	"B Z": 6 + 3,
	"C X": 6 + 1,
	"C Y": 0 + 2,
	"C Z": 3 + 3,
}

func main() {
	file, _ := os.Open("2/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	score := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		score += points2[picks[line]]
	}
	fmt.Println(score)
}
