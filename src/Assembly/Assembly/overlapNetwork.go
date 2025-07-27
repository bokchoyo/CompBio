package main

import (
	"runtime"
	"sync"
)

// MakeOverlapNetwork() takes a slice of reads with match, mismatch, gap and a threshold.
// It returns adjacency lists of reads; edges are only included
// in the overlap graph is the alignment score of the two reads is at least the threshold.
// func MakeOverlapNetwork(reads []string, match, mismatch, gap, threshold float64) map[string][]string {
// 	mtx := OverlapScoringMatrix(reads, match, mismatch, gap) //Initialize an adjacency list to represent the overlap graph.
// 	binMtx := BinarizeMatrix(mtx, threshold)
// 	adjacencyList := make(map[string][]string)

// 	for r := range binMtx {
// 		for c := range binMtx {
// 			if binMtx[r][c] == 1 {
// 				adjacencyList[reads[r]] = append(adjacencyList[reads[r]], reads[c])
// 			}
// 		}
// 	}

// 	return adjacencyList
// }

func MakeOverlapNetwork(reads []string, match, mismatch, gap, threshold float64) map[string][]string {
	n := len(reads)

	scores := make([][]float64, n)
	for i := range scores {
		scores[i] = make([]float64, n)
	}

	type job struct{ i, j int }

	numWorkers := runtime.NumCPU()
	jobs := make(chan job, n*n)
	var wg sync.WaitGroup

	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range jobs {
				i, j := task.i, task.j
				if i == j {
					continue
				}
				scores[i][j] = ScoreOverlapAlignment(reads[i], reads[j], match, mismatch, gap)
			}
		}()
	}

	for i := range n {
		for j := range n {
			jobs <- job{i, j}
		}
	}
	close(jobs)
	wg.Wait()

	adjList := make(map[string][]string, n)
	for i := range n {
		for j := range n {
			if i == j {
				continue
			}
			if scores[i][j] >= threshold {
				adjList[reads[i]] = append(adjList[reads[i]], reads[j])
			}
		}
	}

	return adjList
}
