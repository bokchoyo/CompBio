package main

//RichnessMap takes a map of frequency maps as input.  It returns a map
//whose values are the richness of each sample.

func RichnessMap(allMaps map[string](map[string]int)) map[string]int {
	sampleRichnesses := map[string]int{}
	for i, sample := range allMaps {
		sampleRichnesses[i] = Richness(sample)
	}
	return sampleRichnesses
}
