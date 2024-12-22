package day21

import (
	"fmt"
	"log"
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

// |   | 0    | 1    | 2   | 3    | 4    | 5   | 6   | 7     | 8    | 9    | A     |
// |---|------|------|-----|------|------|-----|-----|-------|------|------|-------|
// | 0 |      | ^<   | ^   | ^>   | ^^<  | ^^  | ^^> | ^^^<  | ^^^  | ^^^> | >     |
// | 1 | >v   |      | >   | >>   | ^    | ^>  | ^>> | ^^    | ^^>  | ^^>> | >>>v  |
// | 2 | v    | <    |     | >    | ^<   | ^   | ^>  | ^^<   | ^^   | ^^>  | v>    |
// | 3 | <v   | <<   | <   |      | ^<<  | ^<  | ^   | ^^<<  | ^^<  | ^^   | v     |
// | 4 | >vv  | v    | v>  | v>>  |      | >   | >>  | ^     | ^>   | ^>>  | >>vv  |
// | 5 | vv   | v<   | v   | v>   | <    |     | >   | ^<    | ^    | ^>   | vv>   |
// | 6 | <vv  | <<v  | <v  | v    | <<   | <   |     | <<^   | <^   | ^    | vv    |
// | 7 | >vvv | vv   | vv> | vv>> | v    | v>  | v>> |       | >    | >>   | >>vvv |
// | 8 | vvv  | vv<  | vv  | vv>  | v<   | v   | v>  | <     |      | >    | vvv>  |
// | 9 | <vvv | <<vv | <vv | vv   | <<v  | <v  | v   | <<    | <    |      | vvv   |
// | A | <    | ^<<  | ^<  | ^    | ^^<< | ^^< | ^^  | ^^^<< | ^^^< | ^^^  |       |
var numericMoves = map[string]string{
	"0:0": "", "0:1": "^<", "0:2": "^", "0:3": "^>", "0:4": "^^<",
	"0:5": "^^", "0:6": "^^>", "0:7": "^^^<", "0:8": "^^^", "0:9": "^^^>",
	"0:A": ">",
	"1:0": ">v", "1:1": "", "1:2": ">", "1:3": ">>", "1:4": "^",
	"1:5": "^>", "1:6": "^>>", "1:7": "^^", "1:8": "^^>", "1:9": "^>>",
	"1:A": ">>>v",
	"2:0": "v", "2:1": "<", "2:2": "", "2:3": ">", "2:4": "^<",
	"2:5": "^", "2:6": "^>", "2:7": "^^<", "2:8": "^^", "2:9": "^^>",
	"2:A": "v>",
	"3:0": "<v", "3:1": "<<", "3:2": "<", "3:3": "", "3:4": "^<<",
	"3:5": "^<", "3:6": "^", "3:7": "^^<<", "3:8": "^^<", "3:9": "^^",
	"3:A": "v",
	"4:0": ">vv", "4:1": "v", "4:2": "v>", "4:3": "v>>", "4:4": "",
	"4:5": ">", "4:6": ">>", "4:7": "^", "4:8": "^>", "4:9": "^>>",
	"4:A": ">>vv",
	"5:0": "vv", "5:1": "v<", "5:2": "v", "5:3": "v>", "5:4": "<",
	"5:5": "", "5:6": ">", "5:7": "^<", "5:8": "^", "5:9": "^>",
	"5:A": "vv>",
	"6:0": "<vv", "6:1": "<<v", "6:2": "<v", "6:3": "v", "6:4": "<<",
	"6:5": "<", "6:6": "", "6:7": "<<^", "6:8": "<^", "6:9": "^",
	"6:A": "vv",
	"7:0": ">vvv", "7:1": "vv", "7:2": "vv>", "7:3": "vv>>", "7:4": "v",
	"7:5": "v>", "7:6": "v>>", "7:7": "", "7:8": ">", "7:9": ">>",
	"7:A": ">>vvv",
	"8:0": "vvv", "8:1": "vv<", "8:2": "vv", "8:3": "vv>", "8:4": "v<",
	"8:5": "v", "8:6": "v>", "8:7": "<", "8:8": "", "8:9": ">",
	"8:A": "vvv>",
	"9:0": "<vvv", "9:1": "<<vv", "9:2": "<vv", "9:3": "vv", "9:4": "<<v",
	"9:5": "<v", "9:6": "v", "9:7": "<<", "9:8": "<", "9:9": "",
	"9:A": "vvv",
	"A:0": "<", "A:1": "^<<", "A:2": "^<", "A:3": "^", "A:4": "^^<<",
	"A:5": "^^<", "A:6": "^^", "A:7": "^^^<<", "A:8": "^^^<", "A:9": "^^^",
	"A:A": "",
}

var numericPositions = map[string]string{
	"0:0": "7", "0:1": "8", "0:2": "9",
	"1:0": "4", "1:1": "5", "1:2": "6",
	"2:0": "1", "2:1": "2", "2:2": "3",
	"3:0": "error", "3:1": "0", "3:2": "A",
}
var numericStart = point.NewPoint(3, 2)

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

// |   | ^  | >  | v  | <   | A   |
// |---|----|----|----|-----|-----|
// | ^ |    | v> | v  | v<  | >   |
// | > | <^ |    | <  | <<  | ^   |
// | v | ^  | >  |    | <   | >^  |
// | < | >^ | >> | >  |     | >>^ |
// | A | <  | v  | <v | v<< |     |
var directionalMoves = map[string]string{
	"^:^": "", "^:>": "v>", "^:v": "v", "^:<": "v<", "^:A": ">",
	">:^": "<^", ">:>": "", ">:v": "<", ">:<": "<<", ">:A": "^",
	"v:^": "^", "v:>": ">", "v:v": "", "v:<": "<", "v:A": ">^",
	"<:^": ">^", "<:>": ">>", "<:v": ">", "<:<": "", "<:A": ">>^",
	"A:^": "<", "A:>": "v", "A:v": "<v", "A:<": "v<<", "A:A": "",
}

var directionalPositions = map[string]string{
	"0:0": "error", "0:1": "^", "0:2": "A",
	"1:0": "<", "1:1": "v", "1:2": ">",
}
var directionalStart = point.NewPoint(0, 2)

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
var fdl, _ = fn.ReadLines()
var matDirectional = matrix.ParseRuneMatrix(fdl)

func reconstruct(sequence string, options map[string]string, start point.Point) string {
	result := ""
	for _, s := range sequence {
		switch s {
		case 'A':
			result += options[start.Hash()]
		case '^':
			start.Sum(point.Point(point.UP))
		case '>':
			start.Sum(point.Point(point.RIGHT))
		case 'v':
			start.Sum(point.Point(point.DOWN))
		case '<':
			start.Sum(point.Point(point.LEFT))
		}
	}
	return result
}

func hash(start string, end string) string {
	return fmt.Sprintf("%s:%s", start, end)
}

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

func seqDirectionalToNumeric(code string, matNumeric matrix.Matrix, matDirectional matrix.Matrix) string {
	moves1 := ""
	start := "A"
	for _, c := range code {
		end := string(c)
		startP := numericPoints[start]
		endP := numericPoints[end]
		_, paths := algorithms.FindMazePath(matNumeric, startP, endP, int('#'))
		costs := make([]int, len(paths))
		for i, path := range paths {
			seq := pathToSequence(path)

			pseq := seqDirectionalToDirectional(seq, matDirectional)
			fmt.Println("pseq", pseq)
			cost := len(pseq)
			costs[i] = cost
		}
		bestSeq := ""
		bestPath := []point.Point{}
		minCost := 1000
		for i, c := range costs {
			if c < minCost {
				bestPath = paths[i]
				minCost = c
			}
		}
		bestSeq = pathToSequence(bestPath)
		fmt.Println(start, end, bestSeq)
		moves1 += bestSeq + "A"
		start = end
	}
	r := reconstruct(moves1, numericPositions, numericStart)
	log.Println(code, moves1, "reconstruct", r)
	return moves1
}

func seqDirectionalToDirectional(code string, matDirectional matrix.Matrix) string {
	moves1 := ""
	start := "A"
	for _, c := range code {
		end := string(c)
		startP := directionPoints[start]
		endP := directionPoints[end]
		_, paths := algorithms.FindMazePath(matDirectional, startP, endP, int('#'))
		costs := make([]int, len(paths))
		for i, path := range paths {
			seq := pathToSequence(path)
			totalCost := 0
			startP := directionPoints["A"]
			for _, c := range seq {
				end := string(c)
				endP := directionPoints[end]
				cost, _ := algorithms.FindMazePath(matDirectional, startP, endP, int('#'))
				totalCost += cost
				startP = endP
			}
			costs[i] = totalCost
		}
		bestSeq := ""
		bestPath := []point.Point{}
		minCost := 1000
		for i, c := range costs {
			if c < minCost {
				bestPath = paths[i]
				minCost = c
			}
		}
		bestSeq = pathToSequence(bestPath)

		moves1 += bestSeq + "A"
		start = end
	}
	r := reconstruct(moves1, directionalPositions, directionalStart)
	log.Println(code, moves1, "reconstruct", r)
	return moves1
}

func generatePermutationsRecursive(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}

	permutations := []string{}

	for i, c := range str {
		remaining := str[:i] + str[i+1:]
		subPermutations := generatePermutationsRecursive(remaining)

		for _, p := range subPermutations {
			permutations = append(permutations, string(c)+p)
		}
	}

	return permutations
}

func dp(sequence string, limit int, depth int) int {
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
			seqs[i] = pathToSequence(path) + "A"
		}
		m := 1000
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
	f := filereader.NewFromDayExample(21, 1)
	lines := []string{}
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		lines = append(lines, line)
	}

	solution := 0
	log.Println("The solution is:", solution)
}
