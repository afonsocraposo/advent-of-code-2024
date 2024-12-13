package algorithms

import (
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/point"
)

func FloodFill(mat matrix.Matrix, start point.Point) []point.Point {
	region := map[point.Point]bool{start: true}
	queue := []point.Point{start}
	sv, err := mat.Get(start.I, start.J)
	if err != nil {
		panic(err)
	}
	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		for _, direction := range point.DIRECTIONS {
			p := point.Point{I: start.I, J: start.J}
			p.Sum(point.Point(direction))

			_, alreadyVisited := region[p]
			if alreadyVisited {
				continue
			}

			v, err := mat.Get(p.I, p.J)
			if err != nil || v != sv {
				region[p] = false
				continue
			}

			region[p] = true
			queue = append(queue, p)
		}
	}
	result := []point.Point{}
	for k, v := range region {
		if v {
			result = append(result, k)
		}
	}
	return result
}
