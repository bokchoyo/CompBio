package main

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your AddRowCol() function here, along with any subroutines that you need. Please note the subroutines indicated in the problem description that are provided for you.
func AddRowCol(row, col, clusterSize1, clusterSize2 int, mtx DistanceMatrix) DistanceMatrix {
	numRows := len(mtx)
	newSize := numRows + 1
	newRow := make([]float64, newSize)

	for r := 0; r < numRows; r++ {
		if r == row || r == col {
			continue
		}
		dist := (float64(clusterSize1)*mtx[row][r] + float64(clusterSize2)*mtx[col][r]) / float64(clusterSize1+clusterSize2)
		newRow[r] = dist
	}
	mtx = append(mtx, newRow)
	for r := 0; r < numRows; r++ {
		mtx[r] = append(mtx[r], newRow[r])
	}
	mtx[newSize-1][newSize-1] = 0.0

	return mtx
}
