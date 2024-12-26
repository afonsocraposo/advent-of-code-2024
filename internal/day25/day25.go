package day25

import (
	"fmt"
	"log"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
)

const day = 25

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
				continue
			}

			exampleSolution := partMethod(exampleLines)
			if exampleSolution != expectedSolution {
				log.Fatalf("WRONG solution for example %d. Expected: %s, Actual: %s\n", example, expectedSolution, exampleSolution)
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

func getHeights(mat matrix.Matrix) []int {
	m, n := mat.Size()
	result := make([]int, n)
	for j := range n {
		c := -1
		for i := range m {
			v, _ := mat.Get(i, j)
			if v == int('#') {
				c++
			}
		}
		result[j] = c
	}
	return result
}

func isLock(mat matrix.Matrix) bool {
	_, n := mat.Size()
	for j := range n {
		v, _ := mat.Get(0, j)
		if v == int('.') {
			return false
		}
	}
	return true
}

func keyFitsLock(key []int, lock []int) bool {
	for i := range key {
		k := key[i]
		l := lock[i]
		if k+l > 5 {
			return false
		}
	}
	return true
}

func part1(lines []string) string {
	buffer := []string{}
	locks := [][]int{}
	keys := [][]int{}
	for i, line := range lines {
		if line == "" || i == len(lines)-1 {
			mat := matrix.ParseRuneMatrix(buffer)
			heights := getHeights(mat)
			if isLock(mat) {
				locks = append(locks, heights)
			} else {
				keys = append(keys, heights)
			}
			buffer = []string{}
		} else {
			buffer = append(buffer, line)
		}
	}

	solution := 0
	for _, key := range keys {
		for _, lock := range locks {
			if keyFitsLock(key, lock) {
				solution++
			}
		}
	}
	return fmt.Sprintf("%d", solution)
}

func part2(lines []string) string {
	return "Merry Christmas"
}
