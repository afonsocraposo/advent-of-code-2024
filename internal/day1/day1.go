package day1

import (
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
)

func Main() {
	log.Println("DAY 1")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func part1() {
	f := filereader.NewFromDayInput(1, 1)
	list1 := []int{}
	list2 := []int{}
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		parts := strings.Split(line, "   ")
		number1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalln(err)
		}
		number2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalln(err)
		}

		list1 = append(list1, number1)
		list2 = append(list2, number2)
	}
	sort.Sort(sort.IntSlice(list1))
	sort.Sort(sort.IntSlice(list2))

	sumDiffs := 0
	for i, n1 := range list1 {
		n2 := list2[i]

		diff := math.Abs(float64(n1 - n2))
		sumDiffs = sumDiffs + int(diff)
	}

	log.Println("The solution is:", sumDiffs)
}

func part2() {
	f := filereader.NewFromDayInput(1, 1)

	list1 := []int{}
	hist2 := map[int]int{}

	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		parts := strings.Split(line, "   ")
		number1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalln(err)
		}
		number2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalln(err)
		}

		list1 = append(list1, number1)

		v, found := hist2[number2]
		if !found {
			v = 0
		}
		hist2[number2] = v + 1
	}

	score := 0
	for _, n1 := range list1 {
		v, found := hist2[n1]
		if !found {
			continue
		}

		score = score + n1*v
	}

	log.Println("The solution is:", score)
}
