package day16

import (
	"fmt"
	"log"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/numbers"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/point"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/queue"
)

func Main() {
	log.Println("DAY 16")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

type reindeer struct {
	Position  point.Point
	Direction point.Direction
}

type priorityQueueElement struct {
	value     int
	Position  point.Point
	Direction point.Direction
}

func (el priorityQueueElement) Value() int {
	return el.value
}

func costsHash(position point.Point, direction point.Direction) string {
	return fmt.Sprintf("%d:%d:%d:%d", position.I, position.J, direction.I, direction.J)
}

func parentHash(position point.Point) string {
	return fmt.Sprintf("%d:%d", position.I, position.J)
}

func getRotateCount(oldDir point.Direction, newDir point.Direction) int {
	dirs := map[point.Direction]int{
		point.UP:    0,
		point.RIGHT: 1,
		point.DOWN:  2,
		point.LEFT:  3,
	}

	n := dirs[newDir]
	o := dirs[oldDir]

    diff := (n - o + 4) % 4       // Clockwise difference
    return numbers.IntMin(diff, 4-diff)      // Return the minimum of clockwise and counter-clockwise rotations
}

func reconstructPath(parent map[string]*point.Point, goal point.Point) []point.Point {
	path := []point.Point{}
	current := &goal
	for current != nil {
		path = append(path, *current)
		current = parent[parentHash(*current)]
	}
	return path
}

// dijkstra algorithm. Find the shortest path in a maze
func part1() {
	f := filereader.NewFromDayInput(16, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}
	mat := matrix.ParseRuneMatrix(lines)

	m, _ := mat.Size()
	r := reindeer{Position: point.Point{I: m - 2, J: 1}, Direction: point.RIGHT}
	start := r.Position

	minHeap := queue.NewPriorityQueue()
	minHeap.Push(priorityQueueElement{
		0, start, r.Direction,
	})
	costs := map[string]int{
		costsHash(start, r.Direction): 0,
	}
	parent := map[string]*point.Point{
		parentHash(start): nil,
	}

	solution := 0
	for minHeap.Len() > 0 {
		current := (*minHeap.Pop()).(priorityQueueElement)
		currentCost := current.Value()
		currentPos := current.Position
		currentDirection := current.Direction

		v, _ := mat.Get(currentPos.I, currentPos.J)
		if v == int('E') {
			solution = currentCost
			break
		}

		for _, direction := range point.DIRECTIONS {
			newPos := currentPos.SumNew(point.Point(direction))
			newDirection := direction

			pd := point.Point(newDirection)
			pd = pd.Symmetric()
			if pd.Equal(point.Point(currentDirection)) {
				continue
			}

			v, err := mat.Get(newPos.I, newPos.J)
			if err != nil {
				continue
			}
			if v != int('#') {
				moveCost := 1 + 1000*getRotateCount(currentDirection, newDirection)
				newCost := currentCost + moveCost

				cost, found := costs[parentHash(newPos)]
				if !found || newCost < cost {
					costs[parentHash(newPos)] = newCost
					parent[parentHash(newPos)] = &currentPos
					minHeap.Push(priorityQueueElement{newCost, newPos, newDirection})
				}
			}
		}
	}
	log.Println("The solution is:", solution)
}

func part2() {
	f := filereader.NewFromDayExample(16, 1)
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
