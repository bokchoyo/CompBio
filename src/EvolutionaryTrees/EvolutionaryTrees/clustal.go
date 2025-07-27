package main

//BuildClustalTree takes a guide tree along with a collection of DNA strings
//labeling the leaves of the tree, where the order of these strings
//is preserved in the tree. The function also takes scoring parameters.
//The function runs the Clustal heuristic to label all of the internal
//nodes of the tree with multiple alignments.
func BuildClustalTree(guideTree Tree, patterns []string, match float64, mismatch float64, gap float64, supergap float64) Tree {

	// initialize final tree according to guide
	initializeGuideTree(guideTree, patterns)

	return guideTree
}

//initializeGuideTree takes a Tree object and a collection of strings as input.
//It sets the alignment of each leaf equal to a blank alignment
//consisting of the string labeling that leaf.
func initializeGuideTree(guide Tree, strs []string) {
	for idx, str := range strs {
		guide[idx].Alignment = Alignment{str}
	}
}
// BuildClustalTree(guideTree, leafStrings, match, mismatch, gap,
// supergap)
// for i ß 0 to length(leafStrings) – 1
// guideTree[i].Alignment = leafStrings[i]
// numNodes ß NumberOfNodes(guideTree)
// for every integer p from numLeaves to numNodes – 1
// alignment1 ß guideTree[p].Child1.Alignment
// alignment2 ß guideTree[p].Child2.Alignment
// guideTree[p].Alignment ß ProgAlign(alignment1,
// alignment2, supergap)
// return guideTree
