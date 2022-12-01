package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("1/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	currentMax := 0
	counter := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			if counter > currentMax {
				currentMax = counter
			}
			counter = 0
		} else {
			calories, _ := strconv.Atoi(line)
			counter += calories
		}
	}
	fmt.Println(currentMax)
}
