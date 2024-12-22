package day21

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/algorithms"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/numbers"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/point"
)

func Main() {
	log.Println("DAY 21")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

var numericPoints = map[string]point.Point{
	"7": point.NewPoint(1, 1),
	"8": point.NewPoint(1, 2),
	"9": point.NewPoint(1, 3),
	"4": point.NewPoint(2, 1),
	"5": point.NewPoint(2, 2),
	"6": point.NewPoint(2, 3),
	"1": point.NewPoint(3, 1),
	"2": point.NewPoint(3, 2),
	"3": point.NewPoint(3, 3),
	"0": point.NewPoint(4, 2),
	"A": point.NewPoint(4, 3),
}

var directionPoints = map[string]point.Point{
	"^": point.NewPoint(1, 2),
	"A": point.NewPoint(1, 3),
	"<": point.NewPoint(2, 1),
	"v": point.NewPoint(2, 2),
	">": point.NewPoint(2, 3),
}

var fn = filereader.NewFromDayExample(21, 2)
var fnl, _ = fn.ReadLines()
var matNumeric = matrix.ParseRuneMatrix(fnl)

var fd = filereader.NewFromDayExample(21, 3)
var fdl, _ = fd.ReadLines()
var matDirectional = matrix.ParseRuneMatrix(fdl)

func pathToSequence(path []point.Point) string {
	sequence := ""
	for i, p := range path[1:] {
		d := p.SumNew(path[i].Symmetric())
		dir := point.Direction(d)
		switch dir {
		case point.UP:
			sequence += "^"
		case point.RIGHT:
			sequence += ">"
		case point.DOWN:
			sequence += "v"
		case point.LEFT:
			sequence += "<"
		}
	}
	return sequence
}

var cache = map[string]int{}

func dp(sequence string, limit int, depth int) int {
	// check cache
	key := fmt.Sprintf("%s:%d:%d", sequence, limit, depth)
	v, found := cache[key]
	if found {
		return v
	}

	var points map[string]point.Point
	var mat matrix.Matrix
	if depth == 0 {
		mat = matNumeric
		points = numericPoints
	} else {
		mat = matDirectional
		points = directionPoints
	}
	start := "A"
	total := 0
	for _, s := range sequence {
		end := string(s)

		startP := points[start]
		endP := points[end]
		_, paths := algorithms.FindMazePath(mat, startP, endP, int('#'))

		seqs := make([]string, len(paths))
		for i, path := range paths {
			seqs[i] = pathToSequence(path)+"A"
		}

		m := int(math.Inf(1))
		if depth == limit {
			for _, seq := range seqs {
				m = numbers.IntMin(m, len(seq))
			}
		} else {
			for _, seq := range seqs {
				m = numbers.IntMin(m, dp(seq, limit, depth+1))
			}
		}
		total += m
		start = end
	}

	cache[key] = total
	return total
}

func part1() {
	f := filereader.NewFromDayInput(21, 1)
	codes, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}

	solution := 0
	for _, code := range codes {
		length := dp(code, 2, 0)
		numeric, _ := strconv.Atoi(code[:3])
		solution += length * numeric
	}
	log.Println("The solution is:", solution)
}

func part2() {
	f := filereader.NewFromDayInput(21, 1)
	codes, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}

	solution := 0
    for _, code := range codes {
		length := dp(code, 25, 0)
        log.Println(length)
		numeric, _ := strconv.Atoi(code[:3])
		solution += length * numeric
	}
	log.Println("The solution is:", solution)
}
