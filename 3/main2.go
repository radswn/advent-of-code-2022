package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func priority_(item rune) int {
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

	i := 0
	itemSum := 0
	rucksackCache := ""
	common := make([]rune, 0)

	for fileScanner.Scan() {
		rucksack := fileScanner.Text()

		switch i {
		case 0:
			rucksackCache = rucksack
		case 1:
			c := []rune(rucksackCache)
			for _, r := range c {
				if strings.ContainsRune(rucksack, r) {
					common = append(common, r)
				}
			}
		case 2:
			for _, r := range common {
				if strings.ContainsRune(rucksack, r) {
					itemSum += priority_(r)
					break
				}
			}
			common = make([]rune, 0)
		}

		i = (i + 1) % 3
	}

	fmt.Println(itemSum)
}
