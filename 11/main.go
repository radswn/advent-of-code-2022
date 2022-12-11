package main

import (
	"fmt"
)

type Operation func(old int64) int64

type ThrowTo func(item int64) int64

type Monkey struct {
	items           []int64
	inspectionCount int
	operation       Operation
	throwTo         ThrowTo
}

func main() {
	monkeys := []Monkey{
		{
			items:           []int64{98, 89, 52},
			inspectionCount: 0,
			operation: func(old int64) int64 {
				return old * 2
			},
			throwTo: func(item int64) int64 {
				if item%5 == 0 {
					return 6
				} else {
					return 1
				}
			},
		},
		{
			items:           []int64{57, 95, 80, 92, 57, 78},
			inspectionCount: 0,
			operation: func(old int64) int64 {
				return old * 13
			},
			throwTo: func(item int64) int64 {
				if item%2 == 0 {
					return 2
				} else {
					return 6
				}
			},
		},
		{
			items:           []int64{82, 74, 97, 75, 51, 92, 83},
			inspectionCount: 0,
			operation: func(old int64) int64 {
				return old + 5
			},
			throwTo: func(item int64) int64 {
				if item%19 == 0 {
					return 7
				} else {
					return 5
				}
			},
		},
		{
			items:           []int64{97, 88, 51, 68, 76},
			inspectionCount: 0,
			operation: func(old int64) int64 {
				return old + 6
			},
			throwTo: func(item int64) int64 {
				if item%7 == 0 {
					return 0
				} else {
					return 4
				}
			},
		},
		{
			items:           []int64{63},
			inspectionCount: 0,
			operation: func(old int64) int64 {
				return old + 1
			},
			throwTo: func(item int64) int64 {
				if item%17 == 0 {
					return 0
				} else {
					return 1
				}
			},
		},
		{
			items:           []int64{94, 91, 51, 63},
			inspectionCount: 0,
			operation: func(old int64) int64 {
				return old + 4
			},
			throwTo: func(item int64) int64 {
				if item%13 == 0 {
					return 4
				} else {
					return 3
				}
			},
		},
		{
			items:           []int64{61, 54, 94, 71, 74, 68, 98, 83},
			inspectionCount: 0,
			operation: func(old int64) int64 {
				return old + 2
			},
			throwTo: func(item int64) int64 {
				if item%3 == 0 {
					return 2
				} else {
					return 7
				}
			},
		},
		{
			items:           []int64{90, 56},
			inspectionCount: 0,
			operation: func(old int64) int64 {
				return old * old
			},
			throwTo: func(item int64) int64 {
				if item%11 == 0 {
					return 3
				} else {
					return 5
				}
			},
		},
	}

	for round := 0; round < 20; round++ {
		for i, monkey := range monkeys {
			fmt.Println("Monkey ", i)
			for _, item := range monkey.items {
				fmt.Println("	Inspecting item ", item)
				newItem := monkey.operation(item)
				fmt.Println("		Worry level after operation ", newItem)
				newItem /= 3
				fmt.Println("		Boring, now ", newItem)
				monkeys[i].inspectionCount++
				newMonkey := monkey.throwTo(newItem)
				fmt.Println("		Throwing to monkey ", newMonkey)
				monkeys[newMonkey].items = append(monkeys[newMonkey].items, newItem)
			}
			monkeys[i].items = []int64{}
		}
		fmt.Println("End of round ", round)
		for i, m := range monkeys {
			fmt.Println("Monkey ", i, ": ", m.items)
		}
	}

	mostActive := 0
	secondMostActive := 0

	for i, monkey := range monkeys {
		fmt.Println("Monkey ", i, " inspections: ", monkey.inspectionCount)
		activity := monkey.inspectionCount
		if activity > mostActive && activity > secondMostActive {
			secondMostActive = mostActive
			mostActive = activity
		} else if activity > secondMostActive {
			secondMostActive = activity
		}
	}

	fmt.Println(mostActive * secondMostActive)
}
