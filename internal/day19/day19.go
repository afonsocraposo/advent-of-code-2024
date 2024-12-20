package day19

import (
	"log"
	"strings"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/algorithms"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
)

func Main() {
	log.Println("DAY 19")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}


func part1() {
	f := filereader.NewFromDayInput(19, 1)
	var towels []string
	designs := []string{}
	i := 0
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		if i == 0 {
			parts := strings.Split(line, ", ")
			towels = make([]string, len(parts))
			for i, p := range parts {
				towels[i] = p
			}
		} else if i > 1 {
			designs = append(designs, line)
		}
		i++
	}

	solution := 0
	for _, design := range designs {
		if algorithms.DpSequenceCheck(design, towels) {
			solution++
		}
	}

	log.Println("The solution is:", solution)
}

func part2() {
	f := filereader.NewFromDayExample(19, 1)
	var towels []string
	designs := []string{}
	i := 0
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		if i == 0 {
			parts := strings.Split(line, ", ")
			towels = make([]string, len(parts))
			for i, p := range parts {
				towels[i] = p
			}
		} else if i > 1 {
			designs = append(designs, line)
		}
		i++
	}

	solution := 0
	for _, design := range designs {
        solution +=  algorithms.DpSequenceArrangementsCount(design, towels)
        break
	}

	log.Println("The solution is:", solution)
}
