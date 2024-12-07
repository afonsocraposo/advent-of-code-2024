package day7

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/numbers"
)

func Main() {
	log.Println("DAY 7")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

type Equation struct {
	testValue int
	numbers   []int
}

func (eq *Equation) isValid2() bool {
	N := numbers.IntPow(2, len(eq.numbers)-1)
	for n := 0; n < N; n++ {
		result := eq.numbers[0]
        operators := n
		for i, v := range eq.numbers[1:] {
			operator := (operators >> i) & 1
			switch operator {
			case 0: // sum
				result += v
			case 1: // multiply
				result *= v
			}
		}
		if result == eq.testValue {
			return true
		}
	}
	return false
}

func (eq *Equation) isValid3() bool {
	N := numbers.IntPow(3, len(eq.numbers)-1)
	for n := 0; n < N; n++ {
		result := eq.numbers[0]
        operators := n
		for _, v := range eq.numbers[1:] {
			operator := operators % 3
			operators /= 3
			switch operator {
			case 0: // sum
				result += v
			case 1: // multiply
				result *= v
			case 2: // concatenation
				concat, err := strconv.Atoi(fmt.Sprintf("%d%d", result, v))
				if err != nil {
					log.Fatalln(err)
				}
				result = concat
			}
		}
		if result == eq.testValue {
			return true
		}
	}
	return false
}

func part1() {
	f := filereader.NewFromDayInput(7, 1)
	solution := 0
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		parts := strings.Split(line, ": ")
		testValue, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalln(err)
		}
		v, err := matrix.ParseVector(parts[1], " ")
		if err != nil {
			log.Fatalln(err)
		}
		eq := Equation{testValue: testValue, numbers: v.Values}
		if eq.isValid2() {
			solution = solution + eq.testValue
		}
	}

	log.Println("The solution is:", solution)
}

func part2() {
	f := filereader.NewFromDayInput(7, 1)
	solution := 0
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		parts := strings.Split(line, ": ")
		testValue, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalln(err)
		}
		v, err := matrix.ParseVector(parts[1], " ")
		if err != nil {
			log.Fatalln(err)
		}
		eq := Equation{testValue: testValue, numbers: v.Values}
		if eq.isValid3() {
			solution = solution + eq.testValue
		}
	}

	log.Println("The solution is:", solution)
}
