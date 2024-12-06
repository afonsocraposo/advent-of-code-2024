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

func NewEmptyMatrix(m int, n int) Matrix {
    rows := make([]Vector, m)
    for i := range rows {
        rows[i] = NewEmptyVector(n)
    }
    return NewMatrix(rows)
}

func ParseMatrix(lines []string, separator string) Matrix {
	m := len(lines)
	rows := make([]Vector, m)
	for i, line := range lines {
		rows[i] = ParseVector(line, separator)
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

func (matrix *Matrix) Row(i int) (*Vector, error) {
	if i < 0 || i >= len(matrix.Rows) {
		return nil, errors.New("Invalid row index")
	}
	vector := NewVector(matrix.Rows[i].Values)
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

func (matrix1 *Matrix) PatternMatch(matrix2 Matrix, mask Matrix) Matrix {
	m1, n1 := matrix1.Size()
	m2, n2 := matrix2.Size()

    matrix2Masked := matrix2.Dot(mask)

	result := make([]Vector, m1-m2+1)
	for i := range len(result) {
		result[i] = NewEmptyVector(n1 - n2 + 1)
		for j := range result[i].Size() {
			p := matrix1.SubMatrix(i, j, i+m2, j+n2)
            pMasked := p.Dot(mask)

			if pMasked.Equal(matrix2Masked) {
				result[i].Set(j, 1)
			} else {
				result[i].Set(j, 0)
			}

		}
	}
	return NewMatrix(result)
}

func (matrix *Matrix) SubMatrix(iStart int, jStart int, iEnd int, jEnd int) Matrix {
	result := make([]Vector, iEnd-iStart)
	for i := range result {
		row := matrix.Rows[iStart+i].Slice(jStart, jEnd)
		result[i] = row
	}
	return NewMatrix(result)
}

func (matrix *Matrix) Equal(m Matrix) bool {
	for i, row := range matrix.Rows {
		mRow, _ := m.Row(i)
		if !row.Equal(*mRow) {
			return false
		}
	}
	return true
}

func (matrix *Matrix) Dot(matrix2 Matrix) Matrix {
	m, n := matrix.Size()

	result := matrix.Clone()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
            val1, _  := matrix.Get(i, j)
            val2, _ := matrix2.Get(i, j)
			result.Set(i, j, val1 * val2)
		}
	}
    return result
}

func (matrix *Matrix) Reduce(fn func(a int, b int) int, initial int) int {
	result := initial
	for _, row := range matrix.Rows {
		result = row.Reduce(fn, result)
	}
	return result
}

func (matrix *Matrix) Clone() Matrix {
	clone := make([]Vector, len(matrix.Rows))
	for i := range clone {
		clone[i] = matrix.Rows[i].Clone()
	}
	return NewMatrix(clone)
}

func (matrix *Matrix) Set(i int, j int, value int) {
	matrix.Rows[i].Set(j, value)
}

func (matrix *Matrix) PrintText() {
    for _, row := range matrix.Rows{
        fmt.Println(row.ToTextString())
    }
}
