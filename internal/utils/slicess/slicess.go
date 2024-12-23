package slicess

func Unique[T comparable](input []T) []T {
	elementSet := make(map[T]struct{})
	uniqueSlice := []T{}

	for _, element := range input {
		if _, exists := elementSet[element]; !exists {
			elementSet[element] = struct{}{}
			uniqueSlice = append(uniqueSlice, element)
		}
	}

	return uniqueSlice
}
