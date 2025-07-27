package main

func DeleteClusters(clusters []*Node, row, col int) []*Node {
	clusters = append(clusters[:col], clusters[col+1:]...)
	clusters = append(clusters[:row], clusters[row+1:]...)
	return clusters
}
