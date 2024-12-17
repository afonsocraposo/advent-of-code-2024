package day17

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/numbers"
)

var numPattern = regexp.MustCompile(`\d+`)

func Main() {
	log.Println("DAY 17")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func combo(operand int, A, B, C int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return A
	case 5:
		return B
	case 6:
		return C
	case 7:
		panic("Combo operand 7 is reserved and will not appear in valid programs.")
	}
	panic("We shouldn't be here. Combo operand is not between 0-6.")
}

func literal(operand int) int {
	return operand
}

func runProgram(program []int, A, B, C int) []int {
	output := []int{}
	pointer := 0
	for pointer < len(program) {
		instruction := program[pointer]
		operand := program[pointer+1]
		switch instruction {
		case 0: // adv
			c := combo(operand, A, B, C)
			value := A / numbers.IntPow(2, c)
			A = value
		case 1: // bxl
			value := B ^ literal(operand)
			B = value
		case 2: // bst
			c := combo(operand, A, B, C)
			value := c % 8
			B = value
		case 3: // jnz
			if A == 0 {
				break
			}
			pointer = literal(operand)
			continue
		case 4: // bxc
			value := B ^ C
			B = value
		case 5: // out
			c := combo(operand, A, B, C)
			value := c % 8
			output = append(output, value)
		case 6: // bdv
			c := combo(operand, A, B, C)
			value := A / numbers.IntPow(2, c)
			B = value
		case 7: // cdv
			c := combo(operand, A, B, C)
			value := A / numbers.IntPow(2, c)
			C = value
		}
		pointer += 2
	}
	return output
}

func part1() {
	f := filereader.NewFromDayInput(17, 1)
	i := 0
	var A, B, C = 0, 0, 0
	var program []int
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		matches := numPattern.FindAllString(line, -1)
		switch i {
		case 0:
			value, err := strconv.Atoi(matches[0])
			if err != nil {
				log.Fatalln(err)
			}
			A = value
		case 1:
			value, err := strconv.Atoi(matches[0])
			if err != nil {
				log.Fatalln(err)
			}
			B = value
		case 2:
			value, err := strconv.Atoi(matches[0])
			if err != nil {
				log.Fatalln(err)
			}
			C = value
		case 4:
			program = make([]int, len(matches))
			for index, match := range matches {
				value, err := strconv.Atoi(match)
				if err != nil {
					log.Fatalln(err)
				}
				program[index] = value
			}
		}
		i++
	}

	output := runProgram(program, A, B, C)
	outputStr := make([]string, len(output))

	for i, v := range output {
		outputStr[i] = fmt.Sprintf("%d", v)
	}
	solution := strings.Join(outputStr, ",")

	log.Println("The solution is:", solution)
}

func part2() {
	f := filereader.NewFromDayExample(17, 1)
	lines := []string{}
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		lines = append(lines, line)
	}

	solution := 0
	log.Println("The solution is:", solution)
}
