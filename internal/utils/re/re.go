package re

import (
	"regexp"
	"strconv"
)

var IntPattern = regexp.MustCompile(`\d+`)

func ParseInts(s string) []int {
	matches := IntPattern.FindAllString(s, -1)
	result := make([]int, len(matches))
	for i, m := range matches {
		v, err := strconv.Atoi(m)
		if err != nil {
			panic(err)
		}
		result[i] = v
	}
	return result
}
