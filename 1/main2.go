package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("1/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	currentMax := make([]int, 3)
	counter := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			if counter > currentMax[0] {
				currentMax[0] = counter
			}
			sort.Ints(currentMax)
			counter = 0
		} else {
			calories, _ := strconv.Atoi(line)
			counter += calories
		}
	}
	result := 0
	for _, e := range currentMax {
		result += e
	}
	fmt.Println(currentMax)
	fmt.Println(result)
}
