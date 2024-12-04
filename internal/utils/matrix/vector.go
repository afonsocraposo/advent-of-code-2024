package matrix

import (
	"log"
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

func ParseVector(line string) Vector {
	parts := strings.Split(line, " ")
	vector := make([]int, len(parts))
	for i, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			log.Fatalln(err)
		}
		vector[i] = n
	}
	return VectorFromSlice(vector)
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
	return NewVector(append([]int{},v.Values...))
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
