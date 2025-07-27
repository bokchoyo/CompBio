package main

// UPGMA takes a distance matrix and a collection of species names as input.
// It returns a Tree (an array of nodes) resulting from applying
// UPGMA to this dataset.
func UPGMA(mtx DistanceMatrix, speciesNames []string) Tree {
	numLeaves := len(mtx)
	t := InitializeTree(speciesNames)
	clusters := InitializeClusters(t)

	for p := numLeaves; p < 2*numLeaves-1; p++ {
		row, col, val := FindMinElement(mtx)
		t[p].Age = val / 2
		t[p].Child1 = clusters[row]
		t[p].Child2 = clusters[col]
		clusterSize1 := CountLeaves(t[p].Child1)
		clusterSize2 := CountLeaves(t[p].Child2)
		mtx = AddRowCol(row, col, clusterSize1, clusterSize2, mtx)
		mtx = DeleteRowCol(mtx, row, col)
		clusters = append(clusters, t[p])
		clusters = DeleteClusters(clusters, row, col)
	}

	return t
}
