package day22

import (
	"fmt"
	"log"
	"strconv"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
)

const day = 22

var examples = []int{1}

func Main() {
	log.Printf("DAY %d\n", day)

	for part := 1; part <= 2; part++ {
		var partMethod func([]string) string
		if part == 1 {
			partMethod = part1
		} else {
			partMethod = part2
		}

		log.Printf("Part %d:\n", part)
		for _, example := range examples {
			exampleLines, err := filereader.ReadDayExample(day, example)
			if err != nil {
				log.Fatalln(err)
			}
			expectedSolution, err := filereader.ReadDayExampleSolution(day, example, part)
			if err != nil {
				log.Fatalln(err)
			}

			exampleSolution := partMethod(exampleLines)
			if exampleSolution != expectedSolution {
				log.Fatalf("WRONG solution for example %d. Expected: %s, Actual: %s\n", example, expectedSolution, exampleSolution)
			} else {
				log.Printf("The solution is CORRECT for example %d. Expected/actual: %s\n", example, exampleSolution)
			}

			inputLines, err := filereader.ReadDayInput(day, 1)
			if err != nil {
				log.Fatalln(err)
			}
			inputSolution := partMethod(inputLines)
            log.Printf("The solution for the input is: %s\n", inputSolution)
		}
	}
}

func next(n int) int {
	tmp := n * 64
	n ^= tmp
	n %= 16777216

	tmp = n / 32
	n ^= tmp
	n %= 16777216

	tmp = n * 2048
	n ^= tmp
	n %= 16777216

	return n
}

func secretNumberN(number int, N int) int {
	for range N {
		number = next(number)
	}
	return number
}

func part1(lines []string) string {
	solution := 0
	for _, line := range lines {
		start, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln(err)
		}
		secret := secretNumberN(start, 2000)
		solution += secret
	}

	return fmt.Sprintf("%d", solution)
}

func part2(lines []string) string {
	fmt.Println(lines)

	solution := 0
	return fmt.Sprintf("%d", solution)
}
