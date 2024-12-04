package matrix

import (
	"errors"
	"fmt"
)

type Matrix struct {
	Rows []Vector
}

func NewMatrix(rows []Vector) Matrix {
	return Matrix{rows}
}

func ParseMatrix(lines []string) Matrix {
	m := len(lines)
	rows := make([]Vector, m)
	for i, line := range lines {
		rows[i] = ParseVector(line)
	}
	return NewMatrix(rows)
}

func ParseRuneMatrix(lines []string) Matrix {
	m := len(lines)
	rows := make([]Vector, m)
	for i, line := range lines {
		rows[i] = ParseRuneVector(line)
	}
	return NewMatrix(rows)
}

func (matrix *Matrix) Size() (int, int) {
	m := len(matrix.Rows)
	if m == 0 {
		return 0, 0
	}

	n := len(matrix.Rows[0].Values)
	return m, n
}

func (matrix *Matrix) Transpose() (*Matrix, error) {
	m, n := matrix.Size()
	if m == 0 {
		return nil, errors.New("Nothing to transpose")
	}

	t := make([]Vector, n)
	for i := 0; i < m; i++ {
		col, err := matrix.Column(i)
		if err != nil {
			return nil, err
		}
		t[i] = *col
	}
	transposed := NewMatrix(t)
	return &transposed, nil
}

func (matrix *Matrix) Mirror() (*Matrix, error) {
	m, n := matrix.Size()
	if m == 0 || n == 0 {
		return nil, errors.New("Empty matrix")
	}
	reversedRows := make([]Vector, m)
	for i, row := range matrix.Rows {
		reversedRows[i] = row.Reverse()
	}
	return &Matrix{reversedRows}, nil
}

func (matrix *Matrix) Row(m int) (*Vector, error) {
	if m < 0 || m >= len(matrix.Rows) {
		return nil, errors.New("Invalid row index")
	}
	vector := NewVector(matrix.Rows[m].Values)
	return &vector, nil
}

func (matrix *Matrix) Column(n int) (*Vector, error) {
	if n < 0 || len(matrix.Rows) == 0 || n >= matrix.Rows[0].Size() {
		return nil, errors.New("Invalid row index")
	}
	m, _ := matrix.Size()
	col := make([]int, m)
	for i, row := range matrix.Rows {
		col[i] = row.Get(n)
	}
	vector := NewVector(col)
	return &vector, nil
}

func (matrix *Matrix) Diagonals() []Vector {
	m, n := matrix.Size()
	if m == 0 || n == 0 {
		return []Vector{}
	}

	diagonals := make([]Vector, m+n-1)
	for k := 0; k < m; k++ {
		diagonal := make([]int, k+1)
		for index := 0; index <= k; index++ {
			i := k - index
			j := index
			value, err := matrix.Get(i, j)
			if err != nil {
				panic(err)
			}
			diagonal[index] = value
		}
		diagonals[k*2] = NewVector(diagonal)

		if k < m-1 {
			diagonalBottom := make([]int, k+1)
			for index := 0; index <= k; index++ {
				i := k - index
				j := index
				value, err := matrix.Get(m-1-i, n-1-j)
				if err != nil {
					panic(err)
				}
				diagonalBottom[index] = value
			}
			diagonals[k*2+1] = NewVector(diagonalBottom)
		}
	}
	return diagonals
}

func (matrix *Matrix) Get(i int, j int) (int, error) {
	m, n := matrix.Size()
	if m == 0 || n == 0 {
		return -1, errors.New("Empty matrix")
	}

	if i < 0 || i >= m || j < 0 || j >= n {
		return -1, errors.New(fmt.Sprint("Invalid position", i, j))
	}

	return matrix.Rows[i].Get(j), nil
}
