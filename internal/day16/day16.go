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

func positionHash(position point.Point) string {
	return fmt.Sprintf("%d:%d", position.I, position.J)
}

func positionDirectionHash(position point.Point, direction point.Direction) string {
	return fmt.Sprintf("%d:%d:%d:%d", position.I, position.J, direction.I, direction.J)
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

	diff := (n - o + 4) % 4             // Clockwise difference
	return numbers.IntMin(diff, 4-diff) // Return the minimum of clockwise and counter-clockwise rotations
}

func reconstructPath(parent map[string]*point.Point, goal point.Point) []point.Point {
	path := []point.Point{}
	current := &goal
	for current != nil {
		path = append(path, *current)
		current = parent[positionHash(*current)]
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
		positionHash(start): 0,
	}
	parent := map[string]*point.Point{
		positionHash(start): nil,
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
				// rotating has a cost of 1000. Moving forward costs 1
				moveCost := 1 + 1000*getRotateCount(currentDirection, newDirection)
				newCost := currentCost + moveCost

				path := reconstructPath(parent, currentPos)
				mat2 := mat.Clone()
				for _, p := range path {
					mat2.Set(p.I, p.J, int('O'))
				}
				mat2.Set(newPos.I, newPos.J, int('O'))

				h := positionHash(newPos)
				cost, found := costs[h]
				if !found || newCost < cost {
					costs[h] = newCost
					parent[h] = &currentPos
					minHeap.Push(priorityQueueElement{newCost, newPos, newDirection})
				}
			}
		}
	}
	log.Println("The solution is:", solution)
}

func reconstructPath2(parent map[string][]reindeer, goal point.Point, direction point.Direction, path []point.Point) []point.Point {
	paths := []point.Point{}
	paths = append(paths, goal)
	parents, found := parent[positionDirectionHash(goal, direction)]
	if !found {
		return paths
	}
	for _, p := range parents {
		contains := false
		for _, v := range path {
			if p.Position.Equal(v) {
				contains = true
				break
			}
		}
		if !contains {
			paths = append(paths, reconstructPath2(parent, p.Position, p.Direction, paths)...)
		}
	}
	return paths
}

func isSeat(count int, value int) int {
	if value == 'O' {
		count++
	}
	return count
}

func part2() {
	f := filereader.NewFromDayInput(16, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatalln(err)
	}
	mat := matrix.ParseRuneMatrix(lines)

	m, _ := mat.Size()
	r := reindeer{Position: point.Point{I: m - 2, J: 1}, Direction: point.RIGHT}
	start := r.Position
	goal := start
	goalDirection := point.UP

	minHeap := queue.NewPriorityQueue()
	minHeap.Push(priorityQueueElement{
		0, start, r.Direction,
	})
	costs := map[string]int{
		positionHash(start): 0,
	}
	parent := map[string][]reindeer{
		positionHash(start): {},
	}

	for minHeap.Len() > 0 {
		current := (*minHeap.Pop()).(priorityQueueElement)
		currentCost := current.Value()
		currentPos := current.Position
		currentDirection := current.Direction

		v, _ := mat.Get(currentPos.I, currentPos.J)
		if v == int('E') {
			goal = currentPos
			goalDirection = currentDirection
			break
		}

		for _, direction := range point.DIRECTIONS {
			// if the direction is the same, move forward. if the direction is different, simply turn
			newDirection := direction
			var newPos point.Point
			if newDirection == currentDirection {
				newPos = currentPos.SumNew(point.Point(direction))
				pd := point.Point(newDirection)
				pd = pd.Symmetric()
				if pd.Equal(point.Point(currentDirection)) {
					continue
				}
			} else {
				newPos = currentPos.Clone()
			}

			v, err := mat.Get(newPos.I, newPos.J)
			if err != nil {
				continue
			}
			if v != int('#') {
				var moveCost int
				// if the direction is the same, move forward. if the direction is different, simply turn
				if newDirection == currentDirection {
					moveCost = 1
				} else {
					moveCost = 1000*getRotateCount(currentDirection, newDirection)
				}
				newCost := currentCost + moveCost

				pdh := positionDirectionHash(newPos, newDirection)
				cost, found := costs[pdh]

				if !found || newCost < cost {
					costs[pdh] = newCost
					minHeap.Push(priorityQueueElement{newCost, newPos, newDirection})
					parent[pdh] = []reindeer{{currentPos, currentDirection}}
				} else if newCost == cost {
					parent[pdh] = append(parent[pdh], reindeer{currentPos, currentDirection})
				}
			}
		}
	}

	paths := reconstructPath2(parent, goal, goalDirection, []point.Point{})
	mat2 := mat.Clone()
	for _, p := range paths {
		mat2.Set(p.I, p.J, int('O'))
	}
	solution := mat2.Reduce(isSeat, 0)
	log.Println("The solution is:", solution)
}
