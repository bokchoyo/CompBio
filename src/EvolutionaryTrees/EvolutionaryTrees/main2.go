package main

import (
	"fmt"
	"strconv"
)

// these two functions take a really long time to run (around 8 hours each)

func SARSSpikeAlignment() {
	fmt.Println("Reading patient virus genomes.")
	directory := "Data/UK-Genomes"
	genomeDatabase := ReadGenomesFromDirectory(directory)
	fmt.Println("Database read! Now, let's grab some spike proteins.")
	date := "2020_11_16"
	_, ok := genomeDatabase[date]

	if !ok {
		panic("Date not found in genome database.")
	}

	spikeProteins := make([]string, 0, len(genomeDatabase[date]))

	for i := range genomeDatabase[date] {
		currGenome := genomeDatabase[date][i]
		spikeDNA := ExciseSpikeProtein(currGenome)
		if spikeDNA != "" {
			rnaStrand := DNAToRNA(spikeDNA)
			proteinSequence := Translate(rnaStrand, 0)
			spikeProteins = append(spikeProteins, proteinSequence)
		}
	}

	fmt.Println(len(spikeProteins), "spike proteins excised for", date+".", "Making distance matrix.")
	mtx := CalculateDistanceMatrix(spikeProteins)
	fmt.Println("Distance matrix made.")
	labels := make([]string, len(spikeProteins))

	for i := range labels {
		labels[i] = strconv.Itoa(i)
	}

	fmt.Println("Building UPGMA tree.")
	T := UPGMA(mtx, labels)
	fmt.Println("UPGMA tree built!")
	match := 1.0
	mismatch := 1.0
	gap := 5.0
	supergap := 5000.0
	fmt.Println("Now, building the Clustal tree.")
	T = BuildClustalTree(T, spikeProteins, match, mismatch, gap, supergap)
	fmt.Println("Clustal tree built! Writing alignment to file.")
	alignment := GetMultipleAlignment(T)
	alignment = RemoveGaps(alignment) // do this after running pipeline the first time
	WriteAlignmentToFile(alignment, SequenceOrder(T), "Output/UK-Genomes", "spike_"+date+".aln")
	fmt.Println("Pipeline complete.")
}

func RemoveGaps(a Alignment) Alignment {
	for j := len(a[0]) - 1; j >= 0; j-- {
		if AllGaps(a, j) {
			for i := range a {
				a[i] = a[i][:j] + a[i][j+1:]
			}
		}
	}
	return a
}

func AllGaps(a Alignment, j int) bool {
	for i := range a {
		if a[i][j] != '-' {
			return false
		}
	}
	return true
}

func IdentifyVariants() {
	fmt.Println("Variant finding! First, reading database.")
	directory := "Data/UK-Genomes"
	genomeDatabase := ReadGenomesFromDirectory(directory)
	k := 10
	allMaps := KmerMapsFromGenomeDatabase(genomeDatabase, k)
	fmt.Println("Made", len(allMaps), "kmer maps!")
	fmt.Println("Forming distance matrix.")
	sampleNames, mtxJac := BetaDiversityMatrix(allMaps, "Jaccard")
	fmt.Println("Writing distance matrix to file.")
	jaccardFile := "Matrices/JaccardBetaDiversityMatrix.csv"
	WriteBetaDiversityMatrixToFile(mtxJac, sampleNames, jaccardFile)
}

func IdentifyVariantsLabelled() {
	fmt.Println("Variant finding! First, reading database.")
	directory := "Data/UK-Genomes"
	genomeDatabase := ReadGenomesFromDirectory(directory)
	k := 10
	allMaps := KmerMapsFromGenomeDatabaseLabelled(genomeDatabase, k)
	fmt.Println("Made", len(allMaps), "kmer maps!")
	fmt.Println("Forming distance matrix.")
	sampleNames, mtxJac := BetaDiversityMatrix(allMaps, "Jaccard")
	fmt.Println("Writing distance matrix to file.")
	jaccardFile := "Matrices/JaccardBetaDiversityMatrixLabelled.csv"
	WriteBetaDiversityMatrixToFile(mtxJac, sampleNames, jaccardFile)
}
