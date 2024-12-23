package mapp

func GetMapKeys[M ~map[K]V, K comparable, V any](m M) []K {
    r := make([]K, 0, len(m))
    for k := range m {
        r = append(r, k)
    }
    return r
}
