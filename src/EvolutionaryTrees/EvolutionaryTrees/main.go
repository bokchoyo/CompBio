package main

import (
	"fmt"
	"strconv"
)

func main() {
	HemoglobinUPGMA()
}

func HemoglobinUPGMA() {
	fmt.Println("Reading in hemoglobin files.")
	stringMap := ReadDNAStringsFromFile("Data/HBA1/hemoglobin_protein.fasta")
	speciesNames, proteinStrings := GetKeyValues(stringMap)
	fmt.Println("Making distance matrix.")
	mtx := CalculateDistanceMatrix(proteinStrings)
	fmt.Println("Distance matrix made.")
	fmt.Println("Starting UPGMA")
	t := UPGMA(mtx, speciesNames)
	fmt.Println("Tree built! Writing to file.")
	WriteNewickToFile(t, "Output/HBA1", "hba1.tre")
	fmt.Println("Tree written to file.")
}

func SARS2UPGMA() {
	fmt.Println("Reading in patient virus genomes.")
	directory := "Data/UK-Genomes"
	genomeDatabase := ReadGenomesFromDirectory(directory)
	fmt.Println("Database read! Let's grab some spike proteins.")
	numberOfGenomesPerDate := 3
	numberOfDates := len(genomeDatabase)
	spikeProteins := make([]string, 0, numberOfGenomesPerDate*numberOfDates)
	labels := make([]string, 0, numberOfGenomesPerDate*numberOfDates)

	for date := range genomeDatabase {
		count := 0

		for i := 0; i < len(genomeDatabase[date]) && count < numberOfGenomesPerDate; i++ {
			currentGenome := genomeDatabase[date][i]
			spikeDNA := ExciseSpikeProtein(currentGenome)
			if spikeDNA != "" && ValidDNA(spikeDNA) {
				rnaStrand := DNAToRNA(spikeDNA)
				proteinSequence := Translate(rnaStrand, 0)
				spikeProteins = append(spikeProteins, proteinSequence)
				currentLabel := date + "_" + strconv.Itoa(i)
				labels = append(labels, currentLabel)
				count++
			}
		}
	}

	fmt.Println("Spike proteins now excised.")
	fmt.Println("Making distance matrix.")
	mtx := CalculateDistanceMatrix(spikeProteins)
	fmt.Println("Distance matrix made.")
	fmt.Println("Starting UPGMA.")
	t := UPGMA(mtx, labels)
	fmt.Println("Tree built! Writing to file.")
	WriteNewickToFile(t, "Output/UK-Genomes", "sars-cov-2.tre")
	fmt.Println("Tree written to file.")
}

func Process16SUPGMA(year int) {
	
}
