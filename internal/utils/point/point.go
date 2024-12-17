package point

type Point struct {
	I int
	J int
}

func (p *Point) InsideBounds(iMin int, jMin int, iMax int, jMax int) bool {
	return p.I >= iMin && p.I <= iMax && p.J >= jMin && p.J <= jMax
}

func (p *Point) Sum(p2 Point) {
	p.I = p.I + p2.I
	p.J = p.J + p2.J
}

func (p *Point) SumNew(p2 Point) Point {
	return Point{I: p.I + p2.I, J: p.J + p2.J}
}

func (p *Point) Clone() Point {
	return Point{I: p.I, J: p.J}
}

func (p *Point) Equal(p2 Point) bool {
    return p.I == p2.I && p.J == p2.J
}

func (p *Point) Symmetric() Point {
	return Point{I: -p.I, J: -p.J}
}

type Direction Point

var (
	UP    = Direction{-1, 0}
	RIGHT = Direction{0, 1}
	DOWN  = Direction{1, 0}
	LEFT  = Direction{0, -1}
	TL    = Direction{-1, -1}
	TR    = Direction{-1, 1}
	BL    = Direction{1, -1}
	BR    = Direction{1, 1}
)

var DIRECTIONS = []Direction{UP, RIGHT, DOWN, LEFT}
var DIRECTIONS9 = []Direction{UP, RIGHT, DOWN, LEFT, TL, TR, BL, BR}

func (p *Point) Distance(p2 Point) Point {
	i := p2.I - p.I
	j := p2.J - p.J
	return Point{I: i, J: j}
}
