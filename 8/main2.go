package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("8/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	text := fileScanner.Text()

	d := len(text)

	treeRows := make([][]int, d)
	treeCols := make([][]int, d)

	scenic := map[int]map[int]int{}

	for i := range text {
		scenic[i] = map[int]int{}
		for j := range text {
			scenic[i][j] = 1
		}
	}

	i_ := 0

	for fileScanner.Scan() {
		for j, e := range text {
			number, _ := strconv.Atoi(string(e))
			treeRows[i_] = append(treeRows[i_], number)
			treeCols[j] = append(treeCols[j], number)
		}

		i_++
		text = fileScanner.Text()
	}

	for j, e := range text {
		number, _ := strconv.Atoi(string(e))
		treeRows[i_] = append(treeRows[i_], number)
		treeCols[j] = append(treeCols[j], number)
	}

	for i, row := range treeRows {
		for j, e := range row {
			scenicSum := 1

			for _, r_ := range row[j+1:] {
				if r_ >= e {
					break
				}
				scenicSum++
			}
			if scenicSum > len(row[j+1:]) {
				scenicSum--
			}
			scenic[i][j] *= scenicSum

			scenicSum = 1
			for i__ := range row[:j] {
				r_ := row[j-i__-1]
				if r_ >= e {
					break
				}
				scenicSum++
			}
			if scenicSum > len(row[:j]) {
				scenicSum--
			}
			scenic[i][j] *= scenicSum

			scenicSum = 1
			for _, r_ := range treeCols[j][i+1:] {
				if r_ >= e {
					break
				}
				scenicSum++
			}
			if scenicSum > len(treeCols[j][i+1:]) {
				scenicSum--
			}
			scenic[i][j] *= scenicSum

			scenicSum = 1
			for i__ := range treeCols[j][:i] {
				r_ := treeCols[j][i-i__-1]
				if r_ >= e {
					break
				}
				scenicSum++
			}
			if scenicSum > len(treeCols[j][:i]) {
				scenicSum--
			}
			scenic[i][j] *= scenicSum

			if i*j == 0 || i == d-1 || j == d-1 {
				scenic[i][j] = 0
			}
		}
	}

	scenicMax := 0
	for _, row := range scenic {
		for _, s := range row {
			if s > scenicMax {
				scenicMax = s
			}
		}
	}

	fmt.Print(scenicMax)
}
