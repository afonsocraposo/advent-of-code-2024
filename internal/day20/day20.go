package day20

import (
	"log"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/algorithms"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/numbers"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/point"
)

func Main() {
	log.Println("DAY 20")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func part1() {
	f := filereader.NewFromDayInput(20, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}

	mat := matrix.ParseRuneMatrix(lines)
	m, n := mat.Size()

	var start, end point.Point
	for i := range m {
		for j := range n {
			v, _ := mat.Get(i, j)
			switch v {
			case int('S'):
				start = point.NewPoint(i, j)
			case int('E'):
				end = point.NewPoint(i, j)
			}
		}
	}
	cost, path := algorithms.FindMazePath(mat, start, end, int('#'))

	costFromPoint := map[string]int{}
	for i, p := range path {
		costFromPoint[p.Hash()] = cost - i
	}

	cheats := map[string]int{}
	for i, p := range path {
		for _, dir := range point.DIRECTIONS {
			p1 := p.SumNew(point.Point(dir))
			v, err := mat.GetPoint(p1)
			if err != nil {
				continue
			}
			if v != int('#') {
				continue
			}
			p2 := p1.SumNew(point.Point(dir))
			v, err = mat.GetPoint(p2)
			if err != nil {
				continue
			}
			if v == int('#') {
				continue
			}
			c := costFromPoint[p2.Hash()]
			if c < cost-i {
				cheats[p1.Hash()] = i + c + 2
			}
		}
	}

	savings := map[int]int{}
	for _, v := range cheats {
		s := cost - v
		_, found := savings[s]
		if !found {
			savings[s] = 0
		}
		savings[s] += 1
	}

	solution := 0
	for k, v := range savings {
		if k >= 100 {
			solution += v
		}
	}
	log.Println("The solution is:", solution)
}

func part2() {
	f := filereader.NewFromDayInput(20, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}

	mat := matrix.ParseRuneMatrix(lines)
	m, n := mat.Size()

	var start, end point.Point
	for i := range m {
		for j := range n {
			v, _ := mat.Get(i, j)
			switch v {
			case int('S'):
				start = point.NewPoint(i, j)
			case int('E'):
				end = point.NewPoint(i, j)
			}
		}
	}
	cost, path := algorithms.FindMazePath(mat, start, end, int('#'))

	costFromPoint := map[string]int{}
	for i, p := range path {
		costFromPoint[p.Hash()] = cost - i
	}

	savings := map[int]int{}
	for i, p1 := range path[:len(path)-1] {
		for _, p2 := range path[1+i:] {
			distance := numbers.IntAbs(p1.I-p2.I) + numbers.IntAbs(p1.J-p2.J)
			if distance <= 20 {
				c := i + costFromPoint[p2.Hash()] + distance
				s := cost - c
				if s < 50 {
					continue
				}
				_, found := savings[s]
				if !found {
					savings[s] = 0
				}
				savings[s] += 1
			}
		}
	}

	solution := 0
	for k, v := range savings {
		if k >= 100 {
			solution += v
		}
	}
	log.Println("The solution is:", solution)
}
