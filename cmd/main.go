package main

import (
	"github.com/afonsocraposo/advent-of-code-2024/internal/day11"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day12"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day13"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day14"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day15"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day16"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day17"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day18"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day19"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day20"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day21"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day22"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day23"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day24"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day25"
	"log"
	"os"
	"strconv"

	"github.com/afonsocraposo/advent-of-code-2024/internal/day1"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day10"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day2"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day3"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day4"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day5"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day6"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day7"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day8"
	"github.com/afonsocraposo/advent-of-code-2024/internal/day9"
)

var days = map[int]func(){
	25: day25.Main,
	24: day24.Main,
	23: day23.Main,
	22: day22.Main,
	21: day21.Main,
	20: day20.Main,
	19: day19.Main,
	17: day17.Main,
	18: day18.Main,
	16: day16.Main,
	15: day15.Main,
	14: day14.Main,
	13: day13.Main,
	12: day12.Main,
	11: day11.Main,
	1:  day1.Main,
	2:  day2.Main,
	3:  day3.Main,
	4:  day4.Main,
	5:  day5.Main,
	6:  day6.Main,
	7:  day7.Main,
	8:  day8.Main,
	9:  day9.Main,
	10: day10.Main,
}

func main() {
	log.SetFlags(0)

	log.Println("Advent of Code 2024")

	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("You must specify which day you want to run, e.g.: 1")
	}

	dayStr := args[0]

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		log.Fatalln(err)
	}

	days[day]()
}
