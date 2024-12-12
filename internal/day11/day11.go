package day11

import (
	"fmt"
	"log"
	"strconv"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
)

const DEBUG = false

func Main() {
	log.Println("DAY 11")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

type stone struct {
	Next   *stone
	Number int
}

func (s *stone) Blink() bool {
	// rule 1
	if s.Number == 0 {
		s.Number = 1
		return false
	}

	// rule 2
	digits := fmt.Sprintf("%d", s.Number)
	n := len(digits)
	if n%2 == 0 {
		ld := digits[:n/2]
		rd := digits[n/2:]

		leftHalf, err := strconv.Atoi(ld)
		if err != nil {
			log.Fatalln(err)
		}
		rightHalf, err := strconv.Atoi(rd)
		if err != nil {
			log.Fatalln(err)
		}

		newStone := stone{Next: s.Next, Number: rightHalf}
		s.Next = &newStone
		s.Number = leftHalf

		return true
	}

	// rule 3
	s.Number *= 2024
	return false
}

func count(s *stone) int {
	c := 0
	for s != nil {
		c++
		s = s.Next
	}
	return c
}

func print(s *stone) {
	for s != nil {
		fmt.Printf("%d ", s.Number)
		s = s.Next
	}
	fmt.Print("\n")
}

func part1() {
	f := filereader.NewFromDayInput(11, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}
	v, err := matrix.ParseVector(lines[0], " ")
	if err != nil {
		log.Fatalln(err)
	}

	var firstStone *stone
	var prevStone *stone
	for i, value := range v.Values {
		stone := stone{Next: nil, Number: value}
		if i == 0 {
			firstStone = &stone
		}
		if prevStone != nil {
			prevStone.Next = &stone
		}
		prevStone = &stone
	}

	if DEBUG {
		fmt.Println("Initial arrangement:")
		print(firstStone)
		fmt.Println("")
	}

	BLINKS := 25
	for b := range BLINKS {
		s := firstStone
		for s != nil {
			skipNext := s.Blink()
			if skipNext {
				s = s.Next
			}
			s = s.Next
		}
		if DEBUG {
			fmt.Printf("After %d blink:\n", b+1)
			print(firstStone)
			fmt.Println()
		}
	}

	solution := count(firstStone)
	log.Println("The solution is:", solution)
}

func cloneMap(m map[int]int) map[int]int {
	clone := make(map[int]int)
	for k, v := range m {
		clone[k] = v
	}
	return clone
}

func addToMap(m *map[int]int, k int, v int) {
	if _, ok := (*m)[k]; !ok {
		(*m)[k] = 0
	}
	(*m)[k] += v
}

func part2() {
	f := filereader.NewFromDayInput(11, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}
	v, err := matrix.ParseVector(lines[0], " ")
	if err != nil {
		log.Fatalln(err)
	}

	stones := map[int]int{}
	for _, value := range v.Values {
		addToMap(&stones, value, 1)
	}

	BLINKS := 75
	for range BLINKS {
		clone := cloneMap(stones)
		for stone, count := range clone {
			addToMap(&stones, stone, -count)
			if stones[stone] == 0 {
				delete(stones, stone)
			}

			if stone == 0 {
				addToMap(&stones, 1, count)
				continue
			}

			// rule 2
			digits := fmt.Sprintf("%d", stone)
			n := len(digits)
			if n%2 == 0 {
				ld := digits[:n/2]
				rd := digits[n/2:]

				leftHalf, err := strconv.Atoi(ld)
				if err != nil {
					log.Fatalln(err)
				}
				rightHalf, err := strconv.Atoi(rd)
				if err != nil {
					log.Fatalln(err)
				}

				addToMap(&stones, leftHalf, count)
				addToMap(&stones, rightHalf, count)
				continue
			}

			// rule 3
			addToMap(&stones, stone*2024, count)
		}
	}

	solution := 0
	for _, count := range stones {
		solution += count
	}
	log.Println("The solution is:", solution)
}
