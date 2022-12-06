package main

import (
	"bufio"
	"fmt"
	"os"
)

func allUnique(r []rune, ru rune) bool {
	for _, v := range r {
		if v == ru {
			return true
		}
	}
	if r[0] == r[1] || r[0] == r[2] || r[1] == r[2] {
		return true
	}
	return false
}

func reorder(r []rune, ru rune) {
	r[2] = r[1]
	r[1] = r[0]
	r[0] = ru
}

func main() {
	file, _ := os.Open("6/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()

	text := fileScanner.Text()
	bufferRunes := make([]rune, 3)

	for i, e := range text {
		if !allUnique(bufferRunes, e) && i > 3 {
			fmt.Println(i + 1)
			break
		}
		reorder(bufferRunes, e)
	}

}
