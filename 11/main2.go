package main

import "fmt"

type Operation_ func(old uint64) uint64

type ThrowTo_ func(item uint64) uint64

type Monkey_ struct {
	items           []uint64
	inspectionCount int
	operation       Operation_
	throwTo         ThrowTo_
}

func main() {
	monkeys := []Monkey_{
		{
			items:           []uint64{98, 89, 52},
			inspectionCount: 0,
			operation: func(old uint64) uint64 {
				return old * 2
			},
			throwTo: func(item uint64) uint64 {
				if item%5 == 0 {
					return 6
				} else {
					return 1
				}
			},
		},
		{
			items:           []uint64{57, 95, 80, 92, 57, 78},
			inspectionCount: 0,
			operation: func(old uint64) uint64 {
				return old * 13
			},
			throwTo: func(item uint64) uint64 {
				if item%2 == 0 {
					return 2
				} else {
					return 6
				}
			},
		},
		{
			items:           []uint64{82, 74, 97, 75, 51, 92, 83},
			inspectionCount: 0,
			operation: func(old uint64) uint64 {
				return old + 5
			},
			throwTo: func(item uint64) uint64 {
				if item%19 == 0 {
					return 7
				} else {
					return 5
				}
			},
		},
		{
			items:           []uint64{97, 88, 51, 68, 76},
			inspectionCount: 0,
			operation: func(old uint64) uint64 {
				return old + 6
			},
			throwTo: func(item uint64) uint64 {
				if item%7 == 0 {
					return 0
				} else {
					return 4
				}
			},
		},
		{
			items:           []uint64{63},
			inspectionCount: 0,
			operation: func(old uint64) uint64 {
				return old + 1
			},
			throwTo: func(item uint64) uint64 {
				if item%17 == 0 {
					return 0
				} else {
					return 1
				}
			},
		},
		{
			items:           []uint64{94, 91, 51, 63},
			inspectionCount: 0,
			operation: func(old uint64) uint64 {
				return old + 4
			},
			throwTo: func(item uint64) uint64 {
				if item%13 == 0 {
					return 4
				} else {
					return 3
				}
			},
		},
		{
			items:           []uint64{61, 54, 94, 71, 74, 68, 98, 83},
			inspectionCount: 0,
			operation: func(old uint64) uint64 {
				return old + 2
			},
			throwTo: func(item uint64) uint64 {
				if item%3 == 0 {
					return 2
				} else {
					return 7
				}
			},
		},
		{
			items:           []uint64{90, 56},
			inspectionCount: 0,
			operation: func(old uint64) uint64 {
				return old * old
			},
			throwTo: func(item uint64) uint64 {
				if item%11 == 0 {
					return 3
				} else {
					return 5
				}
			},
		},
	}

	lcm := uint64(5 * 2 * 19 * 7 * 17 * 13 * 3 * 11)

	for round := 0; round < 10000; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				newItem := monkey.operation(item)
				monkeys[i].inspectionCount++
				newMonkey := monkey.throwTo(newItem)
				monkeys[newMonkey].items = append(monkeys[newMonkey].items, newItem%lcm)
			}
			monkeys[i].items = []uint64{}
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
