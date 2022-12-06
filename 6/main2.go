package main

import (
	"bufio"
	"fmt"
	"os"
)

func allUniques(r string, ru rune) bool {
	depth := len(r)
	fmt.Println("checking ", r, " with ", string(ru))
	if depth == 0 {
		return true
	} else {
		for _, e := range r {
			if e == ru {
				fmt.Println("failed ", r, " with ", string(e))
				return false
			}
		}
		return allUniques(r[:depth-1], rune(r[depth-1]))
	}
}

func min(i1 int, i2 int) int {
	if i1 < i2 {
		return i1
	} else {
		return i2
	}
}

func main() {
	file, _ := os.Open("6/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()

	text := fileScanner.Text()

	for i, e := range text {
		allU := allUniques(text[i-min(13, i):i], e)
		if allU && i >= 13 {
			fmt.Println(i + 1)
			break
		}
	}
}
