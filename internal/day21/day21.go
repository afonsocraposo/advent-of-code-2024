package day21

import (
	"fmt"
	"log"
	"strconv"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
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

var numericPositions = map[string]string {
    "0:0": "7", "0:1": "8", "0:2": "9",
    "1:0": "4", "1:1": "5", "1:2": "6",
    "2:0": "1", "2:1": "2", "2:2": "3",
    "3:0": "error", "3:1": "0", "3:2": "A",
}
var numericStart = point.NewPoint(3,2)

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

var directionalPositions = map[string]string {
    "0:0": "error", "0:1": "^", "0:2": "A",
    "1:0": "<", "1:1": "v", "1:2": ">",
}
var directionalStart = point.NewPoint(0,2)

func reconstruct(sequence string, options map[string]string, start point.Point) string {
    result := ""
    for _, s := range sequence {
        switch s {
        case 'A':
            fmt.Println(start, options[start.Hash()])
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

func part1() {
	f := filereader.NewFromDayExample(21, 1)
	codes, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}

	solution := 0
    for _, code := range codes[4:] {
		// moves for numeric keypad outside the door used by robot 1
		moves1 := ""
		start := "A"
		for _, c := range code {
			end := string(c)
			requiredMoves := numericMoves[hash(start, end)]
			moves1 += requiredMoves
			moves1 += "A"
			start = end
		}
        r := reconstruct(moves1, numericPositions, numericStart)
		log.Println(code, moves1, "reconstruct", r)

		// moves for directional keypad in area with high levels of radiation used by robot 2
		moves2 := ""
		start = "A"
		for _, m := range moves1 {
			end := string(m)
			requiredMoves := directionalMoves[hash(start, end)]
			moves2 += requiredMoves
			moves2 += "A"
			start = end
		}
        r = reconstruct(moves2, directionalPositions, directionalStart)
		log.Println(code, moves1, moves2, "reconstruct", r)

		// moves for directional keypad in area at -40 degrees used by robot 3
		moves3 := ""
		start = "A"
		for _, m := range moves2 {
			end := string(m)
			requiredMoves := directionalMoves[hash(start, end)]
			moves3 += requiredMoves
			moves3 += "A"
			start = end
		}
        r = reconstruct(moves3, directionalPositions, directionalStart)
        r2 := reconstruct(r, directionalPositions, directionalStart)
        r3 := reconstruct(r2, numericPositions, numericStart)
		log.Println(code, moves1, moves2, moves3, "reconstruct", r, r2, r3)

		length := len(moves3)
		numeric, err := strconv.Atoi(code[:3])
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(length, numeric)
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
