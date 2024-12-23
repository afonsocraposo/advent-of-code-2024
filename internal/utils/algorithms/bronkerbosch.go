package algorithms

import (
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/mapp"
	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/set"
)

// Finds all maximal cliques in a graph using the Bron-Kerbosch algorithm without pivoting.
//
// Args:
//
//	R: The current clique being built.
//	P: The set of nodes that can still be added to the clique.
//	X: The set of nodes already excluded from the clique.
//	graph: The adjacency list of the graph.
//	cliques: The list to store all maximal cliques found.
func bronKerbosch(R, P, X set.Set, graph map[string]set.Set, cliques *[]set.Set) {
	if len(P) == 0 && len(X) == 0 {
		(*cliques) = append(*cliques, R)
		return
	}
	// iterate through all possible nodes
	for v := range P {
		bronKerbosch(R.Union(set.NewSet(v)), P.Intersection(graph[v]), X.Intersection(graph[v]), graph, cliques)
        P.Remove(v)
        X.Add(v)
	}
}

func LargestClique(graph map[string]set.Set) set.Set {
    cliques := []set.Set{}
    bronKerbosch(set.Set{}, set.NewSet(mapp.GetMapKeys(graph)...), set.Set{}, graph, &cliques)
    var maxSet set.Set
    maxLen := 0
    for _, s := range cliques {
        l := len(s)
        if l > maxLen {
            maxLen = l
            maxSet = s
        }
    }
    return maxSet
}
