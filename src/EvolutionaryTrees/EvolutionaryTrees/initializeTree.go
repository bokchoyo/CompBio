package main

import "strconv"

// InitializeTree takes a slice of n species names as input.
// It returns a rooted binary tree with with 2n-1 total nodes, where
// the leaves are the first n and have the associated species names.
func InitializeTree(speciesNames []string) Tree {
	numLeaves := len(speciesNames)
	//var t Tree // a Tree is []*Node
	var t Tree
	t = make([]*Node, 2*numLeaves-1)
	// all of these pointers have default value of nil; we need to point them at nodes

	// we should create 2n-1 nodes.
	for i := range t {
		var vx Node
		// let's label the first numLeaves nodes with the appropriate species name.
		// by default, vx.age = 0.0, and its children are nil.
		if i < numLeaves {
			//at a leaf ... let's assign its label.
			vx.Label = speciesNames[i]
		} else {
			// let's just give it an unspecific name
			vx.Label = "Ancestor species " + strconv.Itoa(i)
		}
		// indexing the node
		vx.Num = i
		// one thing to do: point t[i] at vx
		t[i] = &vx
	}

	return t
}
