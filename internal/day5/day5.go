package day5

import (
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
)

var pattern = regexp.MustCompile(`XMAS`)

func Main() {
	log.Println("DAY 5")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

type rule struct {
	First int
	Last  int
}

func part1() {
	f := filereader.NewFromDayInput(5, 1)

	instructions := []rule{}
	readingInstructions := true
	solution := 0
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		if line == "" {
			readingInstructions = false
			continue
		}

		if readingInstructions {
			parts := strings.Split(line, "|")
			number1, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatalln(err)
			}
			number2, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalln(err)
			}
			instructions = append(instructions, rule{number1, number2})
		} else {
			update := matrix.ParseVector(line, ",")

			valid := true
			for _, instruction := range instructions {
				lastIndex := slices.Index(update.Values, instruction.Last)
				if lastIndex != -1 {
					firstIndex := slices.Index(update.Values, instruction.First)
					if firstIndex > lastIndex {
						valid = false
						break
					}
				}

			}
			if valid {
				solution = solution + update.Get(update.Size()/2)
			}
		}

	}

	log.Println("The solution is:", solution)
}

func part2() {
	f := filereader.NewFromDayInput(5, 1)

	instructions := []rule{}
	readingInstructions := true
	solution := 0
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		if line == "" {
			readingInstructions = false
			continue
		}

		if readingInstructions {
			parts := strings.Split(line, "|")
			number1, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatalln(err)
			}
			number2, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalln(err)
			}
			instructions = append(instructions, rule{number1, number2})
		} else {
			update := matrix.ParseVector(line, ",")

            add := false
			for true {
                valid := true
				for _, instruction := range instructions {
					lastIndex := slices.Index(update.Values, instruction.Last)
					if lastIndex != -1 {
						firstIndex := slices.Index(update.Values, instruction.First)
						if firstIndex != -1 && firstIndex > lastIndex {
							valid = false
                            add = true

							first := update.Get(firstIndex)
							last := update.Get(lastIndex)
							update.Set(firstIndex, last)
							update.Set(lastIndex, first)

							break
						}
					}

				}
                if valid {
                    break
                }
			}
			if add {
				solution = solution + update.Get(update.Size()/2)
			}
		}

	}

	log.Println("The solution is:", solution)
}
