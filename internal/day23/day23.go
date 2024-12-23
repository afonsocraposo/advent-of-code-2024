package day23

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/algorithms"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/set"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/slicess"
)

const day = 23

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

var connectionPattern = regexp.MustCompile("[a-z]{2}")
var chiefConnectionPattern = regexp.MustCompile("t[a-z]")

func addConnection(m *map[string]set.Set, a, b string) {
	_, found := (*m)[a]
	if !found {
		(*m)[a] = set.Set{}
	}
	_, found = (*m)[b]
	if !found {
		(*m)[b] = set.Set{}
	}

	(*m)[a].Add(b)
	(*m)[b].Add(a)
}

func part1(lines []string) string {
	connections := map[string]set.Set{}
	for _, line := range lines {
		matches := connectionPattern.FindAllString(line, -1)
		a := matches[0]
		b := matches[1]
		addConnection(&connections, a, b)
	}

	triplets := []string{}
	for k, v := range connections {
		if len(v) < 2 {
			continue
		}

		keys := v.Values()
		for i, ka := range keys[:len(keys)-1] {
			for _, kb := range keys[i:] {
				kav := connections[ka]
				_, connectedToB := kav[kb]
				if connectedToB {
                    ttt := []string{k, ka, kb}
                    slices.Sort(ttt)
					triplets = append(triplets, utils.HashStringValues(ttt...))
				}
			}
		}
	}
    uniqTriplets := slicess.Unique(triplets)

    solution := 0
    for _, t := range uniqTriplets {
        m := chiefConnectionPattern.MatchString(t)
        if m {
            solution++
        }
    }

	return fmt.Sprintf("%d", solution)
}

func part2(lines []string) string {
	graph := map[string]set.Set{}
	for _, line := range lines {
		matches := connectionPattern.FindAllString(line, -1)
		a := matches[0]
		b := matches[1]
		addConnection(&graph, a, b)
	}


    clique := algorithms.LargestClique(graph)
    computers := clique.Values()
    slices.Sort(computers)

    solution := strings.Join(computers,",")
	return solution
}
