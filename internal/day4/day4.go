package day4

import (
	"log"
	"regexp"
	"strings"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
)

var pattern = regexp.MustCompile(`XMAS`)

var xmas = []string{
	"M.S\n.A.\nM.S",
	"S.S\n.A.\nM.M",
	"S.M\n.A.\nS.M",
	"M.M\n.A.\nS.S",
}
var replacePattern = regexp.MustCompile(`[^AMS]`)
var mask = []string{
	"1 0 1",
	"0 1 0",
	"1 0 1",
}

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

	mx := matrix.ParseRuneMatrix(lines)
	m, n := mx.Size()

	solution := 0

	// search horizontally
	for i := 0; i < m; i++ {
		row, err := mx.Row(i)
		if err != nil {
			log.Fatal(err)
		}

		solution = solution + len(pattern.FindAllString(row.ToTextString(), -1))

		reversed := row.Reverse()
		solution = solution + len(pattern.FindAllString(reversed.ToTextString(), -1))
	}

	// search vertically
	for j := 0; j < n; j++ {
		col, err := mx.Column(j)
		if err != nil {
			log.Fatal(err)
		}

		solution = solution + len(pattern.FindAllString(col.ToTextString(), -1))

		reversed := col.Reverse()
		solution = solution + len(pattern.FindAllString(reversed.ToTextString(), -1))
	}

	// search diagonals
	diagonals := mx.Diagonals()
	for _, diagonal := range diagonals {
		solution = solution + len(pattern.FindAllString(diagonal.ToTextString(), -1))

		reversed := diagonal.Reverse()
		solution = solution + len(pattern.FindAllString(reversed.ToTextString(), -1))
	}

	mirrored, _ := mx.Mirror()
	diagonalsMirrored := mirrored.Diagonals()
	for _, diagonal := range diagonalsMirrored {
		solution = solution + len(pattern.FindAllString(diagonal.ToTextString(), -1))

		reversed := diagonal.Reverse()
		solution = solution + len(pattern.FindAllString(reversed.ToTextString(), -1))
	}

	log.Println("The solution is:", solution)
}

func sum(a int, b int) int {
	return a + b
}

func part2() {
	f := filereader.NewFromDayInput(4, 1)

	lines, err := f.ReadLines()
	if err != nil {
		log.Fatal(err)
	}

	for i, line := range lines {
		lines[i] = replacePattern.ReplaceAllString(line, ".")
	}

	mx := matrix.ParseRuneMatrix(lines)

	maskMatrix := matrix.ParseMatrix(mask)

	solution := 0
	for _, kPattern := range xmas {
		k := matrix.ParseRuneMatrix(strings.Split(kPattern, "\n"))

		matches := mx.PatternMatch(k, maskMatrix)
		solution = matches.Reduce(sum, solution)
	}

	log.Println("The solution is:", solution)
}
