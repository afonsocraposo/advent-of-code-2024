package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/afonsocraposo/advent-of-code-2024/internal/day1"
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
	1: day1.Main,
	2: day2.Main,
	3: day3.Main,
	4: day4.Main,
	5: day5.Main,
	6: day6.Main,
	7: day7.Main,
	8: day8.Main,
	9: day9.Main,
}

func main() {
	log.SetFlags(0)

	fmt.Println("Advent of Code 2024")

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
