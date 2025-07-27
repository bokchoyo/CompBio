package main

import (
	"fmt"
	"runtime"
	"sync"
)

func MakeOverlapNetworkMinimizers(reads []string, minDic StringIndex, match, mismatch, gap, threshold float64) map[string][]string {
	numReads := len(reads)
	bigNegative := threshold - 1e7
	overlapMatrix := InitializeFloatMatrix(numReads, numReads)
	for i := range overlapMatrix {
		for j := range overlapMatrix[i] {
			overlapMatrix[i][j] = bigNegative
		}
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	type pair struct {
		i, j int
	}
	pairChan := make(chan pair, 1000)
	numWorkers := runtime.NumCPU()
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range pairChan {
				read1 := reads[p.i]
				read2 := reads[p.j]
				score1 := ScoreOverlapAlignment(read1, read2, match, mismatch, gap)
				score2 := ScoreOverlapAlignment(read2, read1, match, mismatch, gap)

				mu.Lock()
				overlapMatrix[p.i][p.j] = score1
				overlapMatrix[p.j][p.i] = score2
				mu.Unlock()
			}
		}()
	}

	counter := 0
	for _, readIndices := range minDic {
		if counter%100 == 0 {
			fmt.Println("Considering element", counter, "of minimizer map")
		}
		for i := 0; i < len(readIndices); i++ {
			for j := i + 1; j < len(readIndices); j++ {
				index1 := readIndices[i]
				index2 := readIndices[j]

				mu.Lock()
				alreadyDone := overlapMatrix[index1][index2] != bigNegative || overlapMatrix[index2][index1] != bigNegative
				mu.Unlock()

				if !alreadyDone {
					pairChan <- pair{index1, index2}
				}
			}
		}
		counter++
	}

	close(pairChan)
	wg.Wait()

	b := BinarizeMatrix(overlapMatrix, threshold)
	return ConvertAdjacencyMatrixToList(reads, b)
}

func ConvertAdjacencyMatrixToList(reads []string, matrix [][]int) map[string][]string {
	adjList := make(map[string][]string)

	for i := 0; i < len(matrix); i++ {
		read := reads[i]
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				adjList[read] = append(adjList[read], reads[j])
			}
		}
	}

	return adjList
}
