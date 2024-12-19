package algorithms

import (
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/point"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/queue"
)

func reconstructPath(parent map[string]*point.Point, goal point.Point) []point.Point {
	path := []point.Point{}
	current := &goal
	for current != nil {
		path = append(path, *current)
		current = parent[current.Hash()]
	}
	return path
}

func FindMazePath(mat matrix.Matrix,start point.Point, end point.Point, wallValue int) (int, []point.Point) {
	minHeap := queue.NewPriorityQueue()
	minHeap.Push(queue.NewPositionPriorityQueueElement(0, start))
	costs := map[string]int{
		start.Hash(): 0,
	}
	parent := map[string]*point.Point{
		start.Hash(): nil,
	}

	cost := -1
	for minHeap.Len() > 0 {
		current := (*minHeap.Pop()).(queue.PositionPriorityQueueElement)
		currentCost := current.Value()
		currentPos := current.Position

		if currentPos.Equal(end) {
			cost = currentCost
			break
		}

		for _, direction := range point.DIRECTIONS {
			newPos := currentPos.SumNew(point.Point(direction))

			v, err := mat.Get(newPos.I, newPos.J)
			if err != nil {
				continue
			}
			if v != wallValue {
				moveCost := 1
				newCost := currentCost + moveCost

				h := newPos.Hash()
				cost, found := costs[h]
				if !found || newCost < cost {
					costs[h] = newCost
					parent[h] = &currentPos
					minHeap.Push(queue.NewPositionPriorityQueueElement(newCost, newPos))
				}
			}
		}
	}

    // solution not found
    if cost == -1 {
        return cost, []point.Point{}
    }

	path := reconstructPath(parent, end)
	mat2 := mat.Clone()
	for _, p := range path {
		mat2.Set(p.I, p.J, int('O'))
	}

	return cost, path
}
