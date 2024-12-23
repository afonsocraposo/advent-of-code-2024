package set

type Set map[string]struct{}

func NewSet(args ...string) Set {
    set := Set{}
    for _, a := range args {
        set[a] = struct{}{}
    }
    return set
}

func (s Set) Values() []string {
    values := make([]string, len(s))
    i := 0
    for k := range s {
        values[i] = k
        i++
    }
    return values
}

func (s Set) Add(value string) {
    s[value] = struct{}{}
}

func (s Set) Contains(value string) bool {
    _, exists := s[value]
    return exists
}

func (s Set) Remove(value string) {
    delete(s, value)
}

func (s Set) Size() int {
    return len(s)
}

func (s Set) Union(value Set) Set {
    result := make(Set)
    for k := range s {
        result[k] = struct{}{}
    }
    for k := range value {
        result[k] = struct{}{}
    }
    return result
}

func (s Set) Intersection(value Set) Set {
    result := make(Set)
    for k := range s {
        if _, found := value[k]; found {
            result[k] = struct{}{}
        }
    }
    return result
}
