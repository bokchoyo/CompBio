package main

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your GreedyAssembler() function here, along with any subroutines that you need. Please note the subroutines indicated in the problem description that are provided for you.
func GreedyAssembler(reads []string) string {
	reads2 := make([]string, len(reads))
	copy(reads2, reads)
	genome := reads2[0]
	k := len(reads[0])
	reads2 = reads2[1:]

	for {
		hasSuffix, i := HasSuffix(genome, reads2, k)
		if !hasSuffix {
			break
		}
		genome = string(reads2[i][0]) + genome
		reads2 = append(reads2[:i], reads2[i+1:]...)
	}

	for {
		hasPrefix, i := HasPrefix(genome, reads2, k)
		if !hasPrefix {
			break
		}
		genome = genome + string(reads2[i][k-1])
		reads2 = append(reads2[:i], reads2[i+1:]...)

	}

	return genome
}

// beginning of genome = end of read
func HasSuffix(genome string, reads []string, k int) (bool, int) {
	for i, read := range reads {
		if genome[0:k-1] == read[len(read)-k+1:] {
			return true, i
		}
	}
	return false, -1
}

// end of genome = beginning  of read
func HasPrefix(genome string, reads []string, k int) (bool, int) {
	for i, read := range reads {
		if genome[len(genome)-k+1:] == read[0:k-1] {
			return true, i
		}
	}
	return false, -1
}
