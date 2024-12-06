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

type Direction Point

var (
	UP    = Direction{-1, 0}
	RIGHT = Direction{0, 1}
	DOWN  = Direction{1, 0}
	LEFT  = Direction{0, -1}
)
