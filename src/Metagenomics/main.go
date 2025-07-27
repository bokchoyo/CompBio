package main

import (
	"fmt"
)

func main() {
	fmt.Println("Metagenomics!")

	AnalyzeYear("2019")
}

func AnalyzeYear(year string) {
	path := fmt.Sprintf("Data/%s_Samples", year)
	filename := fmt.Sprintf("Data/%s_Samples/fall_Allegheny_1.txt", year)
	freqMap := ReadFrequencyMapFromFile(filename)

	fmt.Println("File read successfully")
	fmt.Println("File contains: ", len(freqMap), " total keys.")

	allMaps := ReadSamplesFromDirectory(path)

	fmt.Println("We have ", len(allMaps), " total samples.")
	fmt.Println("Now writing richness and evenness to file")

	richness := RichnessMap(allMaps)
	richnessFile := fmt.Sprintf("Matrices/RichnessMatrix_%s.csv", year)
	WriteRichnessMapToFile(richness, richnessFile)

	simpsons := SimpsonsMap(allMaps)
	simpsonsFile := fmt.Sprintf("Matrices/SimpsonsMatrix%s.csv", year)
	WriteSimpsonsMapToFile(simpsons, simpsonsFile)

	fmt.Println("Successfully wrote richness and evenness to file")

	distMetric := []string{"Bray-Curtis", "Jaccard"}

	for _, metric := range distMetric {
		sampleNames, mtx := BetaDiversityMatrix(allMaps, metric)
		file := fmt.Sprintf("Matrices/%sBetaDiversityMatrix%s.csv", metric, year)
		WriteBetaDiversityMatrixToFile(mtx, sampleNames, file)
	}

	fmt.Println("Successfuly wrote matrices to file")
}
