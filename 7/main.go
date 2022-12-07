package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("7/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	dirStack := make([]string, 0)
	dirSize := map[string]int{}

	for fileScanner.Scan() {
		text := fileScanner.Text()

		if string(text[0]) == "$" {
			if string(text[2:4]) == "cd" {
				dest := strings.Split(text, " ")[2]

				if dest == ".." {
					dirStack = dirStack[:len(dirStack)-1]
				} else {
					fullPath := strings.Join(dirStack, "/") + dest
					dirSize[fullPath] = 0
					dirStack = append(dirStack, dest)
				}
			}
		} else {
			if string(text[0]) != "d" {
				for i, _ := range dirStack {
					dirStackSlice := dirStack[:i + 1]
					fullPath := strings.Join(dirStackSlice, "/")
					fileSize, _ := strconv.Atoi(strings.Split(text, " ")[0])
					dirSize[fullPath] += fileSize
				}
			}
		}
	}

	bigDirSum := 0
	for _, v := range dirSize {
		if v <= 100000 {
			bigDirSum += v
		}
	}
	fmt.Println(bigDirSum)
}
