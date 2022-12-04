package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("4/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	overlapping := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		ranges := strings.Split(line, ",")

		a := strings.Split(ranges[0], "-")
		b := strings.Split(ranges[1], "-")

		aStart, _ := strconv.Atoi(a[0])
		aEnd, _ := strconv.Atoi(a[1])
		bStart, _ := strconv.Atoi(b[0])
		bEnd, _ := strconv.Atoi(b[1])

		if !(aStart > bEnd || aEnd < bStart) {
			overlapping++
		}
	}

	fmt.Println(overlapping)
}
