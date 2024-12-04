package day3

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
)

var mulPattern = regexp.MustCompile(`mul\((\d{1,3},\d{1,3})\)`)
var doPattern = regexp.MustCompile(`do\(\)`)
var dontPattern = regexp.MustCompile(`don't\(\)`)

func Main() {
	log.Println("DAY 3")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func part1() {
	f := filereader.NewFromDayInput(3, 1)

	solution := 0

	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		instructions := getMulInstructions(line)
		for _, pair := range instructions {
			solution = solution + pair[0]*pair[1]
		}

	}

	log.Println("The solution is:", solution)
}

func part2() {
	f := filereader.NewFromDayInput(3, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}

	line := strings.Join(lines, "")

	solution := 0
	instructions := getEnabledMulInstructions(line)
	for _, pair := range instructions {
		solution = solution + pair[0]*pair[1]
	}

	log.Println("The solution is:", solution)
}

func getMulInstructions(input string) [][]int {

	matches := mulPattern.FindAllStringSubmatch(input, -1)

	instructions := make([][]int, len(matches))

	for i, match := range matches {
		m := match[1]
		parts := strings.Split(m, ",")
		number1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalln(err)
		}
		number2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalln(err)
		}
		instructions[i] = []int{number1, number2}
	}

	return instructions
}

func getEnabledMulInstructions(input string) [][]int {

	matchesDo := doPattern.FindAllStringIndex(input, -1)
	matchesDont := dontPattern.FindAllStringIndex(input, -1)
	iDo, iDont := 0, 0

	matchesMul := mulPattern.FindAllStringSubmatchIndex(input, -1)

	instructions := [][]int{}

	state := true
	for _, match := range matchesMul {
		start := match[2]
		end := match[3]

		for iDo < len(matchesDo) && start > matchesDo[iDo][1] {
			state = true
			iDo++
		}
		for iDont < len(matchesDont) && start > matchesDont[iDont][1] {
			state = false
			iDont++
		}
		if state {
			instruction := input[start:end]

			parts := strings.Split(instruction, ",")
			number1, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatalln(err)
			}
			number2, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalln(err)
			}
			instructions = append(instructions, []int{number1, number2})
		}
	}

	return instructions
}
