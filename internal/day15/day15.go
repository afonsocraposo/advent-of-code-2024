package day15

import (
	"fmt"
	"log"
	"strings"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/animation"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/point"
)

const ANIMATE = false

func Main() {
	log.Println("DAY 15")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func hash(i, j int) string {
	return fmt.Sprintf("%d:%d", i, j)
}

func canMove(robot *point.Point, direction point.Direction, boxes map[string]*point.Point, mat matrix.Matrix) bool {
	newPos := robot.SumNew(point.Point(direction))
	v, _ := mat.Get(newPos.I, newPos.J)
	if v == int('#') {
		return false
	}
	_, found := boxes[hash(newPos.I, newPos.J)]
	if !found {
		return true
	} else {
		return canMove(&newPos, direction, boxes, mat)
	}
}

func move(robot *point.Point, direction point.Direction, boxes map[string]*point.Point) {
	robot.Sum(point.Point(direction))
	box, found := boxes[hash(robot.I, robot.J)]
	if found {
		delete(boxes, hash(box.I, box.J))
		for found {
			box.Sum(point.Point(direction))
			boxCopy := box
			h := hash(box.I, boxCopy.J)
			box, found = boxes[h]
			boxes[h] = boxCopy
		}
	}
}

func getFrame(robot *point.Point, boxes map[string]*point.Point, mat matrix.Matrix) matrix.Matrix {
	m, n := mat.Size()
	result := matrix.NewMatrixWithValue(m, n, int('.'))
	for i := range m {
		for j := range n {
			if robot.I == i && robot.J == j {
				result.Set(i, j, int('@'))
				continue
			}
			_, found := boxes[hash(i, j)]
			if found {
				result.Set(i, j, int('O'))
				continue
			}
			v, _ := mat.Get(i, j)
			if v == int('#') {
				result.Set(i, j, int('#'))
				continue
			}
		}
	}
	return result
}

func part1() {
	f := filereader.NewFromDayInput(15, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}

	var puzzle []string
	var instructionsLines []string
	for i, line := range lines {
		if line == "" {
			puzzle = lines[:i]
			instructionsLines = lines[i+1:]
			break
		}
	}
	instructions := strings.Join(instructionsLines, "")
	mat := matrix.ParseRuneMatrix(puzzle)
	var robot *point.Point
	boxes := map[string]*point.Point{}
	m, n := mat.Size()
	for i := range m {
		for j := range n {
			v, _ := mat.Get(i, j)
			switch v {
			case '@':
				robot = &point.Point{I: i, J: j}
			case 'O':
				boxes[hash(i, j)] = &point.Point{I: i, J: j}
			}
		}
	}

	if ANIMATE {
		frame := getFrame(robot, boxes, mat)
		animation.PrintRuneMatrix(frame, "Initial state", false, 60)
	}
	for _, i := range instructions {
		switch i {
		case '>':
			if canMove(robot, point.RIGHT, boxes, mat) {
				move(robot, point.RIGHT, boxes)
			}
		case '<':
			if canMove(robot, point.LEFT, boxes, mat) {
				move(robot, point.LEFT, boxes)
			}
		case '^':
			if canMove(robot, point.UP, boxes, mat) {
				move(robot, point.UP, boxes)
			}
		case 'v':
			if canMove(robot, point.DOWN, boxes, mat) {
				move(robot, point.DOWN, boxes)
			}
		}
		if ANIMATE {
			frame := getFrame(robot, boxes, mat)
			animation.PrintRuneMatrix(frame, fmt.Sprintf("Move %s:", string(i)), true, 100)
		}
	}

	solution := 0
	for _, box := range boxes {
		solution += 100*box.I + box.J
	}
	log.Println("The solution is:", solution)
}

func scalePuzzle(lines []string) []string {
	scaled := make([]string, len(lines))
	for i, line := range lines {
		scaledLine := ""
		for _, r := range line {
			switch r {
			case '#':
				scaledLine += "##"
			case 'O':
				scaledLine += "[]"
			case '.':
				scaledLine += ".."
			case '@':
				scaledLine += "@."
			}
		}
		scaled[i] = scaledLine
	}
	return scaled
}

func canMoveBox(box *point.Point, direction point.Direction, boxes map[string]*point.Point, mat matrix.Matrix) bool {
	newPos := box.SumNew(point.Point(direction))
	vl, _ := mat.Get(newPos.I, newPos.J)
	vr, _ := mat.Get(newPos.I, newPos.J+1)
	if vl == int('#') || vr == int('#') {
		return false
	}

	switch direction {
	case point.LEFT:
		b, found := boxes[hash(newPos.I, newPos.J-1)]
		if !found {
			return true
		} else {
			return canMoveBox(b, direction, boxes, mat)
		}
	case point.RIGHT:
		b, found := boxes[hash(newPos.I, newPos.J+1)]
		if !found {
			return true
		} else {
			return canMoveBox(b, direction, boxes, mat)
		}
	case point.UP, point.DOWN:
		bb := []*point.Point{}
		bc, foundc := boxes[hash(newPos.I, newPos.J)]
		if foundc {
			bb = append(bb, bc)
		}
		bl, foundl := boxes[hash(newPos.I, newPos.J-1)]
		if foundl {
			bb = append(bb, bl)
		}
		br, foundr := boxes[hash(newPos.I, newPos.J+1)]
		if foundr {
			bb = append(bb, br)
		}
		for _, box := range bb {
			if !canMoveBox(box, direction, boxes, mat) {
				return false
			}
		}
		return true
	}
	return true
}

func canMoveRobot(robot *point.Point, direction point.Direction, boxes map[string]*point.Point, mat matrix.Matrix) bool {
	newPos := robot.SumNew(point.Point(direction))
	v, _ := mat.Get(newPos.I, newPos.J)
	if v == int('#') {
		return false
	}

	switch direction {
	case point.LEFT:
		b, found := boxes[hash(newPos.I, newPos.J-1)]
		if !found {
			return true
		} else {
			return canMoveBox(b, direction, boxes, mat)
		}
	case point.RIGHT:
		b, found := boxes[hash(newPos.I, newPos.J)]
		if !found {
			return true
		} else {
			return canMoveBox(b, direction, boxes, mat)
		}
	case point.UP, point.DOWN:
		b, found := boxes[hash(newPos.I, newPos.J)]
		if !found {
			b, found = boxes[hash(newPos.I, newPos.J-1)]
		}
		if found {
			return canMoveBox(b, direction, boxes, mat)
		}
	}
	return true
}

func moveBox(box *point.Point, direction point.Direction, boxes map[string]*point.Point) {
	delete(boxes, hash(box.I, box.J))
	box.Sum(point.Point(direction))
	switch direction {
	case point.LEFT:
		b, found := boxes[hash(box.I, box.J-1)]
		if found {
			moveBox(b, direction, boxes)
		}
	case point.RIGHT:
		b, found := boxes[hash(box.I, box.J+1)]
		if found {
			moveBox(b, direction, boxes)
		}
	case point.UP, point.DOWN:
		b, found := boxes[hash(box.I, box.J)]
		if found {
			moveBox(b, direction, boxes)
		}
		b, found = boxes[hash(box.I, box.J-1)]
		if found {
			moveBox(b, direction, boxes)
		}
		b, found = boxes[hash(box.I, box.J+1)]
		if found {
			moveBox(b, direction, boxes)
		}
	}
	boxes[hash(box.I, box.J)] = box
}

func moveRobot(robot *point.Point, direction point.Direction, boxes map[string]*point.Point) {
	robot.Sum(point.Point(direction))
	switch direction {
	case point.LEFT:
		b, found := boxes[hash(robot.I, robot.J-1)]
		if found {
			moveBox(b, direction, boxes)
		}
	case point.RIGHT:
		b, found := boxes[hash(robot.I, robot.J)]
		if found {
			moveBox(b, direction, boxes)
		}
	case point.UP, point.DOWN:
		b, found := boxes[hash(robot.I, robot.J)]
		if !found {
			b, found = boxes[hash(robot.I, robot.J-1)]
		}
		if found {
			moveBox(b, direction, boxes)
		}
	}
}

func getFrame2(robot *point.Point, boxes map[string]*point.Point, mat matrix.Matrix) matrix.Matrix {
	m, n := mat.Size()
	result := matrix.NewMatrixWithValue(m, n, int('.'))
	for i := range m {
		for j := range n {
			if robot.I == i && robot.J == j {
				result.Set(i, j, int('@'))
				continue
			}
			_, found := boxes[hash(i, j)]
			if found {
				result.Set(i, j, int('['))
				result.Set(i, j+1, int(']'))
				continue
			}
			v, _ := mat.Get(i, j)
			if v == int('#') {
				result.Set(i, j, int('#'))
				continue
			}
		}
	}
	return result
}

func part2() {
	f := filereader.NewFromDayInput(15, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}

	var puzzle []string
	var instructionsLines []string
	for i, line := range lines {
		if line == "" {
			puzzle = lines[:i]
			instructionsLines = lines[i+1:]
			break
		}
	}

	puzzle = scalePuzzle(puzzle)

	instructions := strings.Join(instructionsLines, "")
	mat := matrix.ParseRuneMatrix(puzzle)
	var robot *point.Point
	boxes := map[string]*point.Point{}
	m, n := mat.Size()
	for i := range m {
		for j := range n {
			v, _ := mat.Get(i, j)
			switch v {
			case '@':
				robot = &point.Point{I: i, J: j}
			case '[':
				boxes[hash(i, j)] = &point.Point{I: i, J: j}
			}
		}
	}

	if ANIMATE {
		frame := getFrame2(robot, boxes, mat)
		animation.PrintRuneMatrix(frame, "Initial state", false, 60)
	}
	for _, i := range instructions {
		switch i {
		case '>':
			if canMoveRobot(robot, point.RIGHT, boxes, mat) {
				moveRobot(robot, point.RIGHT, boxes)
			}
		case '<':
			if canMoveRobot(robot, point.LEFT, boxes, mat) {
				moveRobot(robot, point.LEFT, boxes)
			}
		case '^':
			if canMoveRobot(robot, point.UP, boxes, mat) {
				moveRobot(robot, point.UP, boxes)
			}
		case 'v':
			if canMoveRobot(robot, point.DOWN, boxes, mat) {
				moveRobot(robot, point.DOWN, boxes)
			}
		}
		if ANIMATE {
			frame := getFrame2(robot, boxes, mat)
			animation.PrintRuneMatrix(frame, fmt.Sprintf("Move %s:", string(i)), true, 100)
		}
	}

	solution := 0
	for _, box := range boxes {
		solution += 100*box.I + box.J
	}
	log.Println("The solution is:", solution)
}
