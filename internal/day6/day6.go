package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/point"
)

var pattern = regexp.MustCompile(`XMAS`)

const RENDER = false
const INPLACE = true
const INTERACTIVE = false
const FRAMERATE = 10000

func Main() {
	log.Println("DAY 6")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

type rule struct {
	First int
	Last  int
}

func printRow(n int) {
	for j := 0; j < n; j++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
}

func handleAnimation() {
	if INTERACTIVE {
		waitForKeyPress()
	} else {
		time.Sleep(1000 / FRAMERATE * time.Millisecond)
	}
}
func printMatrix(mat matrix.Matrix, redraw bool) {
	m, n := mat.Size()
	if redraw && RENDER && INPLACE {
		fmt.Printf("\033[%dA", m+2)
	}
	for i, vector := range mat.Rows {
		if i == 0 {
			printRow(n + 2)
		}
		fmt.Printf("|%s|\n", vector.ToTextString())
		if i == n-1 {
			printRow(n + 2)
		}
	}
	if RENDER {
		handleAnimation()
	}
}

// waitForKeyPress waits until the user presses Enter or Spacebar
func waitForKeyPress() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press Enter or Spacebar to continue...")
	for {
		input, _ := reader.ReadByte()
		if input == '\n' || input == ' ' { // Check for Enter or Space
			break
		}
	}
	fmt.Print("\033[F\033[K") // Clear the "Press Enter" message
}

func update(mat *matrix.Matrix, guard *point.Point) {
	value, err := mat.Get(guard.I, guard.J)
	if err != nil {
		log.Fatalln(err)
	}
	r := rune(value)
	switch r {
	case '^':
		moveGuard(mat, guard, point.UP)
	case 'V':
		moveGuard(mat, guard, point.DOWN)
	case '>':
		moveGuard(mat, guard, point.RIGHT)
	case '<':
		moveGuard(mat, guard, point.LEFT)
	}
}

func moveGuard(mat *matrix.Matrix, guard *point.Point, direction point.Direction) {
	// check if can move
	value, err := mat.Get(guard.I+direction.I, guard.J+direction.J)
	if value == int('#') || value == int('O') {
		switch direction {
		case point.UP:
			mat.Set(guard.I, guard.J, int('>'))
		case point.DOWN:
			mat.Set(guard.I, guard.J, int('<'))
		case point.RIGHT:
			mat.Set(guard.I, guard.J, int('V'))
		case point.LEFT:
			mat.Set(guard.I, guard.J, int('^'))
		}
		return
	}
	g, _ := mat.Get(guard.I, guard.J)
	mat.Set(guard.I, guard.J, int('·'))
	// move guard
	guard.Sum(point.Point(direction))
	// outside of bounds
	if err != nil {
		return
	}
	mat.Set(guard.I, guard.J, g)
}

func countDots(a, b int) int {
	if rune(b) == '·' {
		return a + 1
	} else {
		return a
	}
}

func part1() {
	f := filereader.NewFromDayInput(6, 1)

	lines, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}

	gi, gj := 0, 0
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, ".", " ")
		index := strings.Index(line, "^")
		if index > -1 {
			gi = i
			gj = index
		}
	}

	mat := matrix.ParseRuneMatrix(lines)
	m, n := mat.Size()
	guard := point.Point{I: gi, J: gj}

	if RENDER {
		printMatrix(mat, false)
	}
	for guard.InsideBounds(0, 0, n-1, m-1) {
		update(&mat, &guard)
		if RENDER {
			printMatrix(mat, true)
		}
	}

	solution := mat.Reduce(countDots, 0)
	log.Println("The solution is:", solution)
}

func updateHistory(mat *matrix.Matrix, guard *point.Point, history *matrix.Matrix) int {
	h, err := history.Get(guard.I, guard.J)
	if err != nil {
		log.Fatalln(err)
	}

	value, err := mat.Get(guard.I, guard.J)
	if err != nil {
		log.Fatalln(err)
	}
	r := rune(value)
	switch r {
	case '^':
		h = h ^ 0b1000
	case 'V':
		h = h ^ 0b0100
	case '>':
		h = h ^ 0b0010
	case '<':
		h = h ^ 0b0001
	}
	history.Set(guard.I, guard.J, h)
	return h
}

func testInfiniteLoop(mat matrix.Matrix, guard point.Point) bool {
	m, n := mat.Size()
	if RENDER {
		printMatrix(mat, false)
	}
	history := matrix.NewEmptyMatrix(m, n)
	for guard.InsideBounds(0, 0, n-1, m-1) {
		update(&mat, &guard)

		if RENDER {
			printMatrix(mat, true)
		}

		h, err := history.Get(guard.I, guard.J)
		// outside bounds
		if err != nil {
			break
		}
		newh := updateHistory(&mat, &guard, &history)
		if newh < h {
			return true
		}
	}
	return false
}

func part2() {
	f := filereader.NewFromDayInput(6, 1)

	lines, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}

	gi, gj := 0, 0
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, ".", " ")
		index := strings.Index(line, "^")
		if index > -1 {
			gi = i
			gj = index
		}
	}

	mat := matrix.ParseRuneMatrix(lines)
	m, n := mat.Size()
	guard := point.Point{I: gi, J: gj}

	for guard.InsideBounds(0, 0, n-1, m-1) {
		update(&mat, &guard)
	}

	blockPos := []point.Point{}
	for i, row := range mat.Rows {
		for j, value := range row.Values {
			if value == int('·') && !(i == gi && j == gj) {
				blockPos = append(blockPos, point.Point{I: i, J: j})
			}
		}
	}

	solution := 0
	for _, block := range blockPos {
		puzzle := matrix.ParseRuneMatrix(lines)
		guard := point.Point{I: gi, J: gj}
		puzzle.Set(block.I, block.J, int('O'))
		if testInfiniteLoop(puzzle, guard) {
			fmt.Printf("Loop found, block position: i=%d j=%d\n", block.I, block.J)
			solution++
		}
	}

	log.Println("The solution is:", solution)
}
