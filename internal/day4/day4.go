package day4

import (
	"log"
	"regexp"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
)

var pattern = regexp.MustCompile(`XMAS`)

func Main() {
	log.Println("DAY 4")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func part1() {
	f := filereader.NewFromDayInput(4, 1)

	lines, err := f.ReadLines()
	if err != nil {
		log.Fatal(err)
	}

	matrix := matrix.ParseRuneMatrix(lines)
	m, n := matrix.Size()

    solution := 0

	// search horizontally
	for i := 0; i < m; i++ {
		row, err := matrix.Row(i)
        if err != nil {
            log.Fatal(err)
        }

        solution = solution + len(pattern.FindAllString(row.ToTextString(), -1))

        reversed := row.Reverse()
        solution = solution + len(pattern.FindAllString(reversed.ToTextString(), -1))
	}

	// search vertically
	for j := 0; j < n; j++ {
		col, err := matrix.Column(j)
        if err != nil {
            log.Fatal(err)
        }

        solution = solution + len(pattern.FindAllString(col.ToTextString(), -1))

        reversed := col.Reverse()
        solution = solution + len(pattern.FindAllString(reversed.ToTextString(), -1))
	}

    // search diagonals
    diagonals := matrix.Diagonals()
    for _, diagonal := range diagonals {
        solution = solution + len(pattern.FindAllString(diagonal.ToTextString(), -1))

        reversed := diagonal.Reverse()
        solution = solution + len(pattern.FindAllString(reversed.ToTextString(), -1))
    }

    mirrored, _ := matrix.Mirror()
    diagonalsMirrored := mirrored.Diagonals()
    for _, diagonal := range diagonalsMirrored {
        solution = solution + len(pattern.FindAllString(diagonal.ToTextString(), -1))

        reversed := diagonal.Reverse()
        solution = solution + len(pattern.FindAllString(reversed.ToTextString(), -1))
    }


	log.Println("The solution is:", solution)
}

func part2() {
}
