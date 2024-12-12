package day12

import (
	"log"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/algorithms"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/point"
)

func Main() {
	log.Println("DAY 12")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func getNextStart(mat matrix.Matrix) (point.Point, bool) {
	m, n := mat.Size()
	for i := range m {
		for j := range n {
			v, _ := mat.Get(i, j)
			if v != -1 {
				return point.Point{I: i, J: j}, true
			}
		}
	}
	return point.Point{I: 0, J: 0}, false
}

func computePrice(mat matrix.Matrix, region []point.Point) int {
	area := 0
	perimeter := 0
	sv, err := mat.Get(region[0].I, region[0].J)
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range region {
		area++
		for _, direction := range point.DIRECTIONS {
			p := point.Point{I: p.I, J: p.J}
			p.Sum(point.Point(direction))

			v, err := mat.Get(p.I, p.J)
			if err != nil || v != sv {
				perimeter++
			}
		}
	}
	return area * perimeter
}

type edge struct {
	point.Point
	point.Direction
}

func foundAdjacent(ed edge, edges []edge) bool {
	p := ed.Point
	direction := ed.Direction
	switch direction {
	case point.UP:
		p.Sum(point.Point(point.RIGHT))
	case point.RIGHT:
		p.Sum(point.Point(point.DOWN))
	case point.DOWN:
		p.Sum(point.Point(point.RIGHT))
	case point.LEFT:
		p.Sum(point.Point(point.DOWN))
	}

	for _, e := range edges {
		if e.Point == p && e.Direction == direction {
			return true
		}
	}
	return false
}

func computePriceWithSides(mat matrix.Matrix, region []point.Point) int {
	area := len(region)
    // Detect corners
    corners := algorithms.GetCorners(mat, region)
	sides := len(corners)
	return area * sides
}

func part1() {
	f := filereader.NewFromDayInput(12, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatal(err)
	}
	mat := matrix.ParseRuneMatrix(lines)
	price := 0
	for true {
		start, found := getNextStart(mat)
		if !found {
			break
		}
		region := algorithms.FloodFill(mat, start)
		price += computePrice(mat, region)
		for _, p := range region {
			mat.Set(p.I, p.J, -1)
		}
	}

	log.Println("The solution is:", price)
}

func part2() {
	f := filereader.NewFromDayInput(12, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatal(err)
	}
	mat := matrix.ParseRuneMatrix(lines)
	price := 0
	for true {
		start, found := getNextStart(mat)
		if !found {
			break
		}
		region := algorithms.FloodFill(mat, start)
		price += computePriceWithSides(mat, region)
		for _, p := range region {
			mat.Set(p.I, p.J, -1)
		}
	}

	log.Println("The solution is:", price)
}
