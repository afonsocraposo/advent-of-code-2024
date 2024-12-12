package algorithms

import (
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/point"
)

type CornerType int

const (
	ConvexTL CornerType = iota
	ConvexTR
	ConvexBL
	ConvexBR
	ConcaveTL
	ConcaveTR
	ConcaveBL
	ConcaveBR
)

type Corner struct {
	point.Point
	CornerType CornerType
}

func GetCorners(mat matrix.Matrix, region []point.Point) []Corner {
	corners := []Corner{}
	for _, p := range region {
		tl := p.SumNew(point.Point(point.TL))
		t := p.SumNew(point.Point(point.UP))
		tr := p.SumNew(point.Point(point.TR))
		l := p.SumNew(point.Point(point.LEFT))
		r := p.SumNew(point.Point(point.RIGHT))
		bl := p.SumNew(point.Point(point.BL))
		b := p.SumNew(point.Point(point.DOWN))
		br := p.SumNew(point.Point(point.BR))

		vtl, _ := mat.Get(tl.I, tl.J)
		vt, _ := mat.Get(t.I, t.J)
		vtr, _ := mat.Get(tr.I, tr.J)
		vl, _ := mat.Get(l.I, l.J)
		v, _ := mat.Get(p.I, p.J)
		vr, _ := mat.Get(r.I, r.J)
		vbl, _ := mat.Get(bl.I, bl.J)
		vb, _ := mat.Get(b.I, b.J)
		vbr, _ := mat.Get(br.I, br.J)

		if vt != v && vl != v {
			corners = append(corners, Corner{Point: p, CornerType: ConvexTL})
		}
		if vt == v && vl == v && vtl != v {
			corners = append(corners, Corner{Point: p, CornerType: ConcaveTL})
		}

		if vt != v && vr != v {
			corners = append(corners, Corner{Point: p, CornerType: ConvexTR})
		}
		if vt == v && vr == v && vtr != v {
			corners = append(corners, Corner{Point: p, CornerType: ConcaveTR})
		}

		if vb != v && vl != v {
			corners = append(corners, Corner{Point: p, CornerType: ConvexBL})
		}
		if vb == v && vl == v && vbl != v {
			corners = append(corners, Corner{Point: p, CornerType: ConcaveBL})
		}

		if vb != v && vr != v {
			corners = append(corners, Corner{Point: p, CornerType: ConvexBR})
		}
		if vb == v && vr == v && vbr != v {
			corners = append(corners, Corner{Point: p, CornerType: ConcaveBR})
		}
	}
    return corners
}
