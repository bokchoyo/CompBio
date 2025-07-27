package main

func main() {
	LocalAlignment("AGCTGATCGTAGAGCTATA", "AGCTATGATATCGCGCTAGAGAGCTTAGGCGAGCTACGAGCGACGATATCCGATCGATCGGATAGCGATATAGAGGGGCGCGCCCTAAGTCA", 1, 1, 100)
}

func Hemoglobin() {
	match := 1.0
	mismatch := 1.0
	gap := 3.0
	zebrafish := ReadFASTAFile(`C:\Users\bokch\git\CompBio\src\Alignment\Alignment\Data\Hemoglobin\Danio_rerio_hemoglobin.fasta`)
	cow := ReadFASTAFile(`C:\Users\bokch\git\CompBio\src\Alignment\Alignment\Data\Hemoglobin\Bos_taurus_hemoglobin.fasta`)
	human := ReadFASTAFile(`C:\Users\bokch\git\CompBio\src\Alignment\Alignment\Data\Hemoglobin\Homo_sapiens_hemoglobin.fasta`)
	gorilla := ReadFASTAFile(`C:\Users\bokch\git\CompBio\src\Alignment\Alignment\Data\Hemoglobin\Gorilla_gorilla_hemoglobin.fasta`)
	alignment1 := GlobalAlignment(zebrafish, human, match, mismatch, gap)
	WriteAlignmentToFASTA(alignment1, "Output/zebrafish_human_hemoglobin.txt")
	alignment2 := GlobalAlignment(cow, human, match, mismatch, gap)
	WriteAlignmentToFASTA(alignment2, "Output/cow_human_hemoglobin.txt")
	alignment3 := GlobalAlignment(gorilla, human, match, mismatch, gap)
	WriteAlignmentToFASTA(alignment3, "Output/gorilla_human_hemoglobin.txt")

}

func SarsAlignment() {
	match := 1.0
	mismatch := 0.7
	gap := 1.4
	var start1, end1, start2, end2 int

	sars1 := ReadFASTAFile(`C:\Users\bokch\git\CompBio\src\Alignment\Alignment\Data\Coronaviruses\SARS-CoV_genome.fasta`)
	sars2 := ReadFASTAFile(`C:\Users\bokch\git\CompBio\src\Alignment\Alignment\Data\Coronaviruses\SARS-CoV-2_genome.fasta`)

	sarsAlignment := GlobalAlignment(sars1, sars2, match, mismatch, gap)
	WriteAlignmentToFASTA(sarsAlignment, "Output/sars_genome_alignment.txt")

	spikeSequence1 := ReadFASTAFile(`C:\Users\bokch\git\CompBio\src\Alignment\Alignment\Data\Coronaviruses\SARS-CoV_genome_spike_protein.fasta`)
	spikeSequence2 := ReadFASTAFile(`C:\Users\bokch\git\CompBio\src\Alignment\Alignment\Data\Coronaviruses\SARS-CoV-2_genome_spike_protein.fasta`)
	spikeAlignment, start1, end1, start2, end2:= LocalAlignment(spikeSequence1, spikeSequence2, match, mismatch, gap)
	WriteLocalAlignmentToFASTA(spikeAlignment, "Output/SARS_spike_SARS-2_genome_alignment.txt", start1, end1, start2, end2)
}
