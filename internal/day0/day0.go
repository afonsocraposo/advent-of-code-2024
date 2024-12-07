package day0

import (
	"log"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
)

func Main() {
	log.Println("DAY 0")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func part1() {
	f := filereader.NewFromDayExample(1, 1)
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

func part2() {
	f := filereader.NewFromDayExample(1, 1)
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

