package day10

import (
	"log"
	"slices"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/point"
)

func Main() {
	log.Println("DAY 10")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func checkPaths(mat matrix.Matrix, p point.Point) []point.Point {
	v, err := mat.Get(p.I, p.J)
	if err != nil {
		return []point.Point{}
	}
	if v == 9 {
		return []point.Point{p}
	}
	result := []point.Point{}
	for _, dir := range point.DIRECTIONS {
		move := point.Point{I: p.I, J: p.J}
		move.Sum(point.Point(dir))
		next, err := mat.Get(move.I, move.J)
		if err != nil {
			continue
		}
		if next != v+1 {
			continue
		}
		result = append(result, checkPaths(mat, move)...)
	}
	return result
}

func part1() {
	f := filereader.NewFromDayInput(10, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}
	mat, err := matrix.ParseMatrix(lines, "")
	if err != nil {
		log.Fatalln(err)
	}

	initial := []point.Point{}
	m, n := mat.Size()
	for i := range m {
		for j := range n {
			v, err := mat.Get(i, j)
			if err != nil {
				continue
			}
			if v == 0 {
				initial = append(initial, point.Point{I: i, J: j})
			}
		}
	}

    solution := 0
	for _, p := range initial {
		nines := checkPaths(mat, p)
		unique := map[point.Point]bool{}
		for _, p := range nines {
			unique[p] = true
		}
		solution += len(unique)
	}

	log.Println("The solution is:", solution)
}

func part2() {
	f := filereader.NewFromDayInput(10, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}
	mat, err := matrix.ParseMatrix(lines, "")
	if err != nil {
		log.Fatalln(err)
	}

	initial := []point.Point{}
	m, n := mat.Size()
	for i := range m {
		for j := range n {
			v, err := mat.Get(i, j)
			if err != nil {
				continue
			}
			if v == 0 {
				initial = append(initial, point.Point{I: i, J: j})
			}
		}
	}

    solution := 0
	for _, p := range initial {
		nines := checkPaths(mat, p)
		solution += len(nines)
	}

	log.Println("The solution is:", solution)
}
