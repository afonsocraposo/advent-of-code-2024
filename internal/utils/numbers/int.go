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
