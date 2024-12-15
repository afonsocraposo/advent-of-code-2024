package day14

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/animation"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/point"
)

var numPattern = regexp.MustCompile(`[-]{0,1}[0-9]+`)

func Main() {
	log.Println("DAY 14")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

type robot struct {
	position point.Point
	velocity point.Point
}

func (r *robot) update(m int, n int) {
	r.position.Sum(r.velocity)
	if r.position.I >= m {
		r.position.I = r.position.I % m
	} else if r.position.I < 0 {
		r.position.I = m + (r.position.I % m)
	}
	if r.position.J >= n {
		r.position.J = r.position.J % n
	} else if r.position.J < 0 {
		r.position.J = n + (r.position.J % n)
	}
}

func countQuadrants(robots []*robot, m int, n int) (int, int, int, int) {
	qtl := 0
	qtr := 0
	qbl := 0
	qbr := 0
	for _, r := range robots {
		if r.position.I < m/2 && r.position.J < n/2 {
			qtl++
		} else if r.position.I < m/2 && r.position.J >= n-n/2 {
			qtr++
		} else if r.position.I >= m-m/2 && r.position.J < n/2 {
			qbl++
		} else if r.position.I >= m-m/2 && r.position.J >= n-n/2 {
			qbr++
		}
	}
	return qtl, qtr, qbl, qbr
}

func robotsToMat(robots []*robot, m, n int) matrix.Matrix {
	mat := matrix.NewEmptyMatrix(m, n)
	for _, r := range robots {
		v, _ := mat.Get(r.position.I, r.position.J)
		mat.Set(r.position.I, r.position.J, v+1)
	}
	return mat
}

func robotsToMat2(robots []*robot, m, n int) matrix.Matrix {
	mat := matrix.NewMatrixWithValue(m, n, int(' '))
	for _, r := range robots {
		mat.Set(r.position.I, r.position.J, int('#'))
	}
	return mat
}

func detectLines(mat matrix.Matrix, value int, length int) bool {
	m, n := mat.Size()
	for i := range m {
		c := 0
		for j := range n {
			v, _ := mat.Get(i, j)
			if v == value {
				c++
			} else {
				c = 0
			}
			if c >= length {
				return true
			}
		}
	}
	for j := range n {
		c := 0
		for i := range m {
			v, _ := mat.Get(i, j)
			if v == value {
				c++
			} else {
				c = 0
			}
			if c >= length {
				return true
			}
		}
	}
	return false
}

func printRobots(robots []*robot, m, n int) {
	mat := robotsToMat(robots, m, n)
	mat.PrintValues()
}

func part1() {
	f := filereader.NewFromDayInput(14, 1)
	robots := []*robot{}
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		matches := numPattern.FindAllString(line, -1)
		if len(matches) == 0 {
			continue
		}
		px, err := strconv.Atoi(matches[0])
		if err != nil {
			log.Fatalln(err)
		}
		py, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Fatalln(err)
		}
		vx, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalln(err)
		}
		vy, err := strconv.Atoi(matches[3])
		if err != nil {
			log.Fatalln(err)
		}
		r := robot{
			position: point.Point{I: py, J: px},
			velocity: point.Point{I: vy, J: vx},
		}
		robots = append(robots, &r)
	}

	M, N := 103, 101

	SECONDS := 100
	for range SECONDS {
		for _, r := range robots {
			r.update(M, N)
		}
	}

	qtl, qtr, qbl, qbr := countQuadrants(robots, M, N)

	solution := qtl * qtr * qbl * qbr
	log.Println("The solution is:", solution)
}

func part2() {
	f := filereader.NewFromDayInput(14, 1)
	robots := []*robot{}
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		matches := numPattern.FindAllString(line, -1)
		if len(matches) == 0 {
			continue
		}
		px, err := strconv.Atoi(matches[0])
		if err != nil {
			log.Fatalln(err)
		}
		py, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Fatalln(err)
		}
		vx, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalln(err)
		}
		vy, err := strconv.Atoi(matches[3])
		if err != nil {
			log.Fatalln(err)
		}
		r := robot{
			position: point.Point{I: py, J: px},
			velocity: point.Point{I: vy, J: vx},
		}
		robots = append(robots, &r)
	}

	M, N := 103, 101

	SECONDS := 10000
	for s := range SECONDS {
		for _, r := range robots {
			r.update(M, N)
		}
		mat := robotsToMat2(robots, M, N)
		if detectLines(mat, int('#'), 8) {
			animation.PrintRuneMatrix(mat, fmt.Sprintf("Seconds: %d", s+1), true, 0)
		}
	}

	qtl, qtr, qbl, qbr := countQuadrants(robots, M, N)

	solution := qtl * qtr * qbl * qbr
	log.Println("The solution is:", solution)
}
