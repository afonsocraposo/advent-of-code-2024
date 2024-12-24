package day24

import (
	"fmt"
	"log"
	"regexp"
	"slices"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
)

const day = 24

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

var wirePattern = regexp.MustCompile(`([a-z0-9]{3}): ([0-1])`)
var instructionPattern = regexp.MustCompile(`([a-z0-9]{3})\s+(AND|OR|XOR)\s+([a-z0-9]{3})\s+->\s+([a-z0-9]{3})`)

func part1(lines []string) string {
	wires := map[string]int{}
	operations := []string{}
	wiresPart := true
	for _, line := range lines {
		if wiresPart {
			m := wirePattern.FindAllStringSubmatch(line, -1)
			if len(m) > 0 {
				wire := m[0][1]
				if m[0][2] == "1" {
					wires[wire] = 1
				} else {
					wires[wire] = 0
				}
				continue
			} else {
				wiresPart = false
				continue
			}
		}
		operations = append(operations, line)
	}

	for len(operations) > 0 {
		operation := operations[0]
		m := instructionPattern.FindAllStringSubmatch(operation, -1)
		if len(m) == 0 {
			log.Fatalln("Something wrong with the instructionPattern")
		}
		a := m[0][1]
		o := m[0][2]
		b := m[0][3]
		c := m[0][4]

		aw, afound := wires[a]
		bw, bfound := wires[b]
		if !afound || !bfound {
			operations = operations[1:]
			operations = append(operations, operation)
			continue
		}

		switch o {
		case "AND":
			wires[c] = aw & bw
		case "OR":
			wires[c] = aw | bw
		case "XOR":
			wires[c] = aw ^ bw
		}
		operations = operations[1:]
	}

	zKeys := []string{}
	for k := range wires {
		if k[0] == 'z' {
			zKeys = append(zKeys, k)
		}
	}
	slices.Sort(zKeys)
	slices.Reverse(zKeys)

	solution := 0
	for i, zk := range zKeys {
		z := wires[zk]
		solution |= z
		if i < len(zKeys)-1 {
			solution <<= 1
		}
	}

	return fmt.Sprintf("%d", solution)
}

func part2(lines []string) string {
	operations := []string{}
	wiresPart := true
	for _, line := range lines {
		if line == "" {
			wiresPart = false
			continue
		}
		if !wiresPart {
			operations = append(operations, line)
		}
	}

	slices.Sort(operations)

	mermaid := "\nflowchart-elk TD\n"

	for i, operation := range operations {
		m := instructionPattern.FindAllStringSubmatch(operation, -1)
		if len(m) == 0 {
			log.Fatalln("Something wrong with the instructionPattern")
		}
		a := m[0][1]
		o := m[0][2]
		b := m[0][3]
		c := m[0][4]
		mermaid += fmt.Sprintf("%s{%s} --> O%d[%s]\n", a, a, i, o)
		mermaid += fmt.Sprintf("%s{%s} --> O%d[%s]\n", b, b, i, o)
		mermaid += fmt.Sprintf("O%d[%s] --> %s{%s}\n", i, o, c, c)
	}
	fmt.Println(mermaid)
	return mermaid
}
