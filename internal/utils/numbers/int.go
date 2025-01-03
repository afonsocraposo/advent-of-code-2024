package numbers

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IntAbs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
