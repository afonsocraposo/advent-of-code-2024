package day13

import (
	"log"
	"math"
	"regexp"
	"strconv"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/numbers"
)

func Main() {
	log.Println("DAY 13")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func computeSolution(ax, ay, bx, by, x, y float64) (float64, float64) {
	b := (y - ay*x/ax) / (by - ay*bx/ax)
	a := (x - bx*b) / ax
	return a, b
}

func part1() {
	f := filereader.NewFromDayInput(13, 1)
	i := 0
    solution := 0
	var ax, ay, bx, by, x, y float64
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		n := i % 4
		switch n {
		case 0:
			buttonPattern := regexp.MustCompile(`\d+`)
			matches := buttonPattern.FindAllString(line, -1)
			ax, err = strconv.ParseFloat(matches[0], 64)
			if err != nil {
				log.Fatalln(err)
			}
			ay, err = strconv.ParseFloat(matches[1], 64)
			if err != nil {
				log.Fatalln(err)
			}
		case 1:
			buttonPattern := regexp.MustCompile(`\d+`)
			matches := buttonPattern.FindAllString(line, -1)
			bx, err = strconv.ParseFloat(matches[0], 64)
			if err != nil {
				log.Fatalln(err)
			}
			by, err = strconv.ParseFloat(matches[1], 64)
			if err != nil {
				log.Fatalln(err)
			}
		case 2:
			prizePattern := regexp.MustCompile(`\d+`)
			matches := prizePattern.FindAllString(line, -1)
			x, err = strconv.ParseFloat(matches[0], 64)
			if err != nil {
				log.Fatalln(err)
			}
			y, err = strconv.ParseFloat(matches[1], 64)
			if err != nil {
				log.Fatalln(err)
			}

			a, b := computeSolution(ax, ay, bx, by, x, y)
            if numbers.IsInt(a) && numbers.IsInt(b) {
                solution += int(math.Round(a))*3 + int(math.Round(b))
            }
		}
		i++
	}

	log.Println("The solution is:", solution)
}

func part2() {
	f := filereader.NewFromDayExample(13, 1)
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
