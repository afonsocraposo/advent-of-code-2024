package matrix

import (
	"slices"
	"strconv"
	"strings"
)

type Vector struct {
	Values []int
}

func NewVector(values []int) Vector {
	return Vector{values}
}

func NewEmptyVector(n int) Vector {
	return Vector{make([]int, n)}
}

func ParseVector(line string, separator string) (Vector, error) {
	parts := strings.Split(line, separator)
	vector := make([]int, len(parts))
	for i, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			return NewEmptyVector(0), err
		}
		vector[i] = n
	}
	return VectorFromSlice(vector), nil
}

func ParseRuneVector(line string) Vector {
	vector := make([]int, len(line))
	for i, r := range line {
		n := int(r)
		vector[i] = n
	}
	return VectorFromSlice(vector)
}

func (v *Vector) All(f func(a int, b int) bool) bool {
	for i := 0; i < len(v.Values)-1; i++ {
		a := v.Values[i]
		b := v.Values[i+1]
		if !f(a, b) {
			return false
		}
	}
	return true
}

func (v *Vector) Any(f func(a int, b int) bool) bool {
	for i := 0; i < len(v.Values)-1; i++ {
		a := v.Values[i]
		b := v.Values[i+1]
		if f(a, b) {
			return true
		}
	}
	return false
}

func VectorFromSlice(s []int) Vector {
	return Vector{s}
}

func (v *Vector) Remove(s int) Vector {
	newValues := make([]int, 0, len(v.Values)-1)
	newValues = append(newValues, v.Values[:s]...)
	newValues = append(newValues, v.Values[s+1:]...)
	return Vector{newValues}
}

func (v *Vector) Size() int {
	return len(v.Values)
}

func (v *Vector) Get(n int) int {
	return v.Values[n]
}

func (v *Vector) Clone() Vector {
	return NewVector(append([]int{}, v.Values...))
}

func (v *Vector) Reverse() Vector {
	reversed := v.Clone()
	slices.Reverse(reversed.Values)
	return reversed
}

func (v *Vector) ToTextString() string {
	runes := make([]rune, v.Size())
	for i, r := range v.Values {
		runes[i] = rune(r)
	}
	return string(runes)
}

func (v *Vector) Equal(v2 Vector) bool {
	return slices.Equal(v.Values, v2.Values)
}

func (v *Vector) Slice(start int, end int) Vector {
	return NewVector(v.Values[start:end])
}

func (v *Vector) Set(position int, value int) {
	v.Values[position] = value
}

func (v *Vector) Reduce(fn func(a int, b int) int, initial int) int {
	result := initial
	for _, value := range v.Values {
		result = fn(result, value)
	}
	return result
}
