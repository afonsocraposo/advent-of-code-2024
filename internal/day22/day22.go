package day22

import (
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
)

const day = 22

var examples = []int{1, 2}

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
				continue
			}

			exampleSolution := partMethod(exampleLines)
			if exampleSolution != expectedSolution {
				log.Printf("WRONG solution for example %d. Expected: %s, Actual: %s\n", example, expectedSolution, exampleSolution)
			} else {
				log.Printf("The solution is CORRECT for example %d. Expected/actual: %s\n", example, exampleSolution)
			}

		}
		inputLines, err := filereader.ReadDayInput(day, 1)
		if err != nil {
			log.Fatalln(err)
		}
		inputSolution := partMethod(inputLines)
		log.Printf("The solution for the input is: %s\n", inputSolution)
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

func hash(pattern []int) string {
	return fmt.Sprintf("%d,%d,%d,%d", pattern[0], pattern[1], pattern[2], pattern[3])
}

func part2(lines []string) string {
	buyersPatterns := []map[string]int{}
	uniquePatterns := []string{}
	for _, line := range lines {
		start, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln(err)
		}
		patterns := map[string]int{}
		buffer := []int{}
		prev := start
		for range 2000 {
			dPrev := prev % 10
			n := next(prev)
			d := n % 10

			diff := d - dPrev
			buffer = append(buffer, diff)

			if len(buffer) >= 4 {
				h := hash(buffer)
				_, found := patterns[h]
				if !found {
					patterns[h] = d
					uniquePatterns = append(uniquePatterns, h)
				}
				buffer = buffer[1:]
			}

			prev = n
		}
		buyersPatterns = append(buyersPatterns, patterns)
	}

	slices.Sort(uniquePatterns)
	uniquePatterns = slices.Compact(uniquePatterns)

	solution := 0
	for _, pattern := range uniquePatterns {
		score := 0
		for _, buyersPattern := range buyersPatterns {
			score += buyersPattern[pattern]
		}
		if score > solution {
			solution = score
		}
	}
	return fmt.Sprintf("%d", solution)
}
