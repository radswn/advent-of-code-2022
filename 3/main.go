package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func priority(item rune) int {
	if unicode.IsLower(item){
		return int(item) - int('a') + 1
	} else {
		return int(item) - int('A') + 27
	}
}

func main() {
	file, _ := os.Open("3/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	itemSum := 0

	for fileScanner.Scan() {
		rucksack := fileScanner.Text()

		com1 := rucksack[:len(rucksack)/2]
		com2 := rucksack[len(rucksack)/2:]

		a := []rune(com1)

		for _, r := range a {
			if strings.ContainsRune(com2, r) {
				itemSum += priority(r)
				break
			}
		}
	}

	fmt.Println(itemSum)
}
