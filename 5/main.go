package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	file, _ := os.Open("5/input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	cargos := map[int]string{
		1: "FHBVRQDP",
		2: "LDZQWV",
		3: "HLZQGRPC",
		4: "RDHFJVB",
		5: "ZWLC",
		6: "JRPNTGVM",
		7: "JRLVMBS",
		8: "DPJ",
		9: "DCNWV",
	}

	for i := 0; i < 10; i++ {
		fileScanner.Scan()
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.Replace(line, "move ", "", 1)
		line = strings.Replace(line, "from ", "", 1)
		line = strings.Replace(line, "to ", "", 1)

		e := strings.Split(line, " ")

		move, _ := strconv.Atoi(e[0])
		from, _ := strconv.Atoi(e[1])
		to, _ := strconv.Atoi(e[2])

		splitIndex := len(cargos[from]) - move

		toCut := cargos[from][splitIndex:]
		cargos[from] = cargos[from][:splitIndex]
		cargos[to] = cargos[to] + reverse(toCut)
	}

	fmt.Println(cargos)
}
