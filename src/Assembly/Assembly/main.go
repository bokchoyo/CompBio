package main

import (
	"fmt"
)

func main() {
	fmt.Println("Genome assembly!")
	SARSOverlapNetworkMinimizerTrim()
}

func SARSOverlapNetwork() {
	fmt.Println("Read in the SARS-CoV-2 genome.")
	genome := ReadGenomeFromFASTA(`C:\Users\bokch\git\CompBio\src\Assembly\Assembly\Data\SARS-CoV-2_genome.fasta`)
	fmt.Println("Genome read. Sampling reads.")
	readLength := 150
	probability := 0.01
	reads := SimulateReadsClean(genome, readLength, probability)
	fmt.Println("Reads generated and building overlap network.")
	match := 1.0
	mismatch := 1.0
	gap := 5.0
	threshold := 40.0
	adjList := MakeOverlapNetwork(reads, match, mismatch, gap, threshold)
	fmt.Println("Overlap network made")
	fmt.Println("The network has", len(adjList), "total keys.")
	fmt.Println("The average outdegree is", AverageOutDegree(adjList))
}

func SARSOverlapNetworkMinimizer() {
	fmt.Println("Read in the SARS-CoV-2 genome.")
	genome := ReadGenomeFromFASTA(`C:\Users\bokch\git\CompBio\src\Assembly\Assembly\Data\SARS-CoV-2_genome.fasta`)
	fmt.Println("Genome read. Sampling reads.")
	readLength := 150
	probability := 0.1
	reads := SimulateReadsClean(genome, readLength, probability)
	fmt.Println("Reads generated and building overlap network.")
	fmt.Println("Making minimzer map.")
	windowLength := 50
	k := 25
	minimizerDictionary := BuildMinimizerMap(reads, k, windowLength)
	fmt.Println("Minimizer map made. It contains", len(minimizerDictionary), "total keys.")
	match := 1.0
	mismatch := 1.0
	gap := 5.0
	threshold := 150.0

	adjList := MakeOverlapNetworkMinimizers(reads, minimizerDictionary, match, mismatch, gap, threshold)
	fmt.Println("Overlap network made")
	fmt.Println("The network has", len(adjList), "total keys.")
	fmt.Println("The average outdegree is", AverageOutDegree(adjList))
}

func SARSOverlapNetworkMinimizerTrim() {
	fmt.Println("Read in the SARS-CoV-2 genome.")
	genome := ReadGenomeFromFASTA(`C:\Users\bokch\git\CompBio\src\Assembly\Assembly\Data\SARS-CoV-2_genome.fasta`)
	fmt.Println("Genome read. Sampling reads.")
	readLength := 150
	probability := 0.1
	reads := SimulateReadsClean(genome, readLength, probability)
	fmt.Println("Reads generated and building overlap network.")
	fmt.Println("Making minimzer map.")
	windowLength := 50
	k := 25
	minimizerDictionary := BuildMinimizerMap(reads, k, windowLength)
	fmt.Println("Minimizer map made. It contains", len(minimizerDictionary), "total keys.")
	match := 1.0
	mismatch := 1.0
	gap := 5.0
	threshold := 150.0
	maxK := 3

	adjList := MakeOverlapNetworkMinimizers(reads, minimizerDictionary, match, mismatch, gap, threshold)
	fmt.Println("Avg outdegree is", AverageOutDegree(adjList))
	trimmedAdjList := TrimNetwork(adjList, maxK)
	fmt.Println("Overlap network made")
	fmt.Println("The network has", len(trimmedAdjList), "total keys.")
	fmt.Println("The average outdegree is", AverageOutDegree(adjList))
}
