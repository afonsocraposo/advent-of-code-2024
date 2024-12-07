package day2

import (
	"log"
	"math"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
)

func Main() {
	log.Println("DAY 2")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func part1() {
	f := filereader.NewFromDayInput(2, 1)
	safeReports := 0
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		report, err := matrix.ParseVector(line, " ")
        if err != nil {
            log.Fatalln(err)
        }

        if !checkReport(report) {
            continue
        }

		safeReports++
	}

	log.Println("The solution is:", safeReports)
}

func part2() {
	f := filereader.NewFromDayInput(2, 1)
	safeReports := 0
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}

		report, err := matrix.ParseVector(line, " ")
        if err != nil {
            log.Fatalln(err)
        }
        if checkReport(report) {
            safeReports++
            continue
        }

        for i := range report.Values {
            newReport := report.Remove(i)
            if checkReport(newReport){
                safeReports++
                break
            }
        }
	}

	log.Println("The solution is:", safeReports)
}

func checkReport(report matrix.Vector) bool {
	incDec := report.All(func(a int, b int) bool { return a < b }) || report.All(func(a int, b int) bool { return a > b })
	if !incDec {
        return false
	}

	diff := report.All(func(a int, b int) bool {
		d := int(math.Abs(float64(a - b)))
		return d >= 1 && d <= 3
	})

	if !diff {
        return false
	}
    return true
}
