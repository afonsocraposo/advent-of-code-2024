package algorithms

import (
	"fmt"
	"slices"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/point"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/queue"
)

type rpath struct {
	Path []point.Point
	Next point.Point
}

func reconstructPaths(parent map[string][]point.Point, goal point.Point, cost int) [][]point.Point {
	rpaths := []rpath{{[]point.Point{}, goal}}
	paths := [][]point.Point{}
	for len(rpaths) > 0 {
		rp := rpaths[0]
		for {
			rp.Path = append(rp.Path, rp.Next)
			parents := parent[rp.Next.Hash()]
			if len(parents) > 0 {
				rp.Next = parents[0]
				for _, p := range parents[1:] {
					nrp := rpath{append([]point.Point{}, rp.Path...), p}
					rpaths = append(rpaths, nrp)
				}
			} else {
				break
			}
		}
		if len(rp.Path)-1 <= cost {
			slices.Reverse(rp.Path)
			paths = append(paths, rp.Path)
		}
		rpaths = rpaths[1:]
	}
	return paths
}

func FindMazePath(mat matrix.Matrix, start point.Point, end point.Point, wallValue int) (int, [][]point.Point) {
	minHeap := queue.NewPriorityQueue()
	minHeap.Push(queue.NewPositionPriorityQueueElement(0, start))
	costs := map[string]int{
		start.Hash(): 0,
	}
	parent := map[string][]point.Point{
		start.Hash(): {},
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
					minHeap.Push(queue.NewPositionPriorityQueueElement(newCost, newPos))
					parent[h] = []point.Point{currentPos}
				} else if newCost == cost {
					parent[h] = append(parent[h], currentPos)
				}
			}
		}
	}

	// solution not found
	if cost == -1 {
        fmt.Println("No solution found")
		return cost, [][]point.Point{}
	}

	paths := reconstructPaths(parent, end, cost)

	return cost, paths
}
