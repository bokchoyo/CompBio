package main

// InitializeClusters takes a tree and returns a slice of
// pointers to the leaves of the tree.
func InitializeClusters(t Tree) []*Node {
	numNodes := len(t)
	numLeaves := (numNodes + 1) / 2
	clusters := make([]*Node, numLeaves)

	for i := range clusters {
		clusters[i] = t[i]
	}

	return clusters
}
