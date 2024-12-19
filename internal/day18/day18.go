package day18

import (
	"fmt"
	"log"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/algorithms"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/animation"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/point"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/re"
)

func Main() {
	log.Println("DAY 18")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func reconstructPath(parent map[string]*point.Point, goal point.Point) []point.Point {
	path := []point.Point{}
	current := &goal
	for current != nil {
		path = append(path, *current)
		current = parent[current.Hash()]
	}
	return path
}

func part1() {
	f := filereader.NewFromDayInput(18, 1)
	M, N := 71, 71
	FALLEN_BYTES := 1024

	coords := []point.Point{}
	i := 0
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		values := re.ParseInts(line)
		x := values[0]
		y := values[1]
		coords = append(coords, point.NewPoint(y, x))
		i++
	}

	mat := matrix.NewMatrixWithValue(M, N, int('.'))
	for _, c := range coords[:FALLEN_BYTES] {
		mat.Set(c.I, c.J, int('#'))
	}

	start := point.NewPoint(0, 0)
	goal := point.NewPoint(M-1, N-1)

    cost, _ := algorithms.FindMazePath(mat, start, goal, int('#'))

	log.Println("The solution is:", cost)
}

func part2() {
	f := filereader.NewFromDayInput(18, 1)
	M, N := 71, 71
	FALLEN_BYTES := 1024

	coords := []point.Point{}
	i := 0
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		values := re.ParseInts(line)
		x := values[0]
		y := values[1]
		coords = append(coords, point.NewPoint(y, x))
		i++
	}

    start := point.NewPoint(0, 0)
    end := point.NewPoint(M-1, N-1)

    var solution point.Point
	for i := range len(coords) - FALLEN_BYTES {
		animation.PrintString(fmt.Sprintf("Attempt %d out of %d", i+1, len(coords)-FALLEN_BYTES), i > 0)
		mat := matrix.NewMatrixWithValue(M, N, int('.'))
		for _, c := range coords[:FALLEN_BYTES+i] {
			mat.Set(c.I, c.J, int('#'))
		}

        cost, _ := algorithms.FindMazePath(mat, start,end,int('#') )

		if cost == -1 {
			solution = coords[FALLEN_BYTES+i-1]
			break
		}
	}

	log.Println("The solution is:", fmt.Sprintf("%d,%d", solution.J, solution.I))
}
