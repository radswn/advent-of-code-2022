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

	treeRows := make([][]int, len(text))
	treeCols := make([][]int, len(text))

	visible := map[int]map[int]bool{}

	for i := range text {
		visible[i] = map[int]bool{}
	}

	i := 0

	for fileScanner.Scan() {
		for j, e := range text {
			number, _ := strconv.Atoi(string(e))
			treeRows[i] = append(treeRows[i], number)
			treeCols[j] = append(treeCols[j], number)
		}

		i++
		text = fileScanner.Text()
	}

	for j, e := range text {
		number, _ := strconv.Atoi(string(e))
		treeRows[i] = append(treeRows[i], number)
		treeCols[j] = append(treeCols[j], number)
	}

	for i, row := range treeRows {
		currentMax := -1
		for j, e := range row {
			if e > currentMax {
				visible[i][j] = true
				currentMax = e
			}
		}

		currentMax = -1
		for j := range row {
			e := row[len(row)-1-j]
			if e > currentMax {
				visible[i][len(row)-1-j] = true
				currentMax = e
			}
		}
	}

	for i, col := range treeCols {
		currentMax := -1
		for j, e := range col {
			if e > currentMax {
				visible[j][i] = true
				currentMax = e
			}
		}

		currentMax = -1
		for j := range col {
			idx := len(col) - 1 - j
			e := col[idx]
			if e > currentMax {
				visible[idx][i] = true
				currentMax = e
			}
		}
	}

	visibleSum := 0

	for _, m := range visible {
		for range m {
			visibleSum++
		}
	}
	fmt.Println(visibleSum)
}
