package day8

import (
	"log"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/point"
)

func Main() {
	log.Println("DAY 8")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func countNonZero(a int, b int) int {
	if b == 0 {
		return a
	}
	return a + 1
}

func part1() {
	f := filereader.NewFromDayInput(8, 1)

	i := 0
	j := 0
	antennas := map[int][]point.Point{}
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		vector := matrix.ParseRuneVector(line)
		j = 0
		for _, antenna := range vector.Values {
			if antenna != int('.') {
				pos := point.Point{I: i, J: j}
				v, found := antennas[antenna]
				if !found {
					antennas[antenna] = []point.Point{}
				}
				antennas[antenna] = append(v, pos)
			}
			j++
		}
		i++
	}

	antinodes := matrix.NewEmptyMatrix(i, j)
	for _, positions := range antennas {
		for i, pos1 := range positions {
			for j, pos2 := range positions {
				if i != j {
					distance := pos2.Distance(pos1)
					antinodeI := pos1.I + distance.I
					antinodeJ := pos1.J + distance.J
					v, err := antinodes.Get(antinodeI, antinodeJ)
					if err != nil {
						// outside of bounds
						continue
					}
					antinodes.Set(antinodeI, antinodeJ, v+1)
				}
			}
		}
	}

	solution := antinodes.Reduce(countNonZero, 0)
	log.Println("The solution is:", solution)
}

func part2() {
    f := filereader.NewFromDayInput(8, 1)

	i := 0
	j := 0
	antennas := map[int][]point.Point{}
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		vector := matrix.ParseRuneVector(line)
		j = 0
		for _, antenna := range vector.Values {
			if antenna != int('.') {
				pos := point.Point{I: i, J: j}
				v, found := antennas[antenna]
				if !found {
					antennas[antenna] = []point.Point{}
				}
				antennas[antenna] = append(v, pos)
			}
			j++
		}
		i++
	}

	antinodes := matrix.NewEmptyMatrix(i, j)
	for _, positions := range antennas {
        if len(positions)<2 {
            continue
        }
		for i, pos1 := range positions {
			for j, pos2 := range positions {
				if i != j {
					distance := pos2.Distance(pos1)
                    antinodeI := pos1.I
                    antinodeJ := pos1.J
					for {
						v, err := antinodes.Get(antinodeI, antinodeJ)
						if err != nil {
							// outside of bounds
							break
						}
						antinodes.Set(antinodeI, antinodeJ, v+1)
						antinodeI += distance.I
						antinodeJ += distance.J
					}
				}
			}
		}
	}

	solution := antinodes.Reduce(countNonZero, 0)
	log.Println("The solution is:", solution)
}
