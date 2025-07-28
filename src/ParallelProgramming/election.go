package main

import (
	"math/rand"
)

func SimulateMultipleElections(polls map[string]float64, electoralVotes map[string]int, numTrials int, marginOfError float64) (float64, float64, float64) {
	winCount1 := 0
	winCount2 := 0
	tieCount := 0

	for i := 0; i < numTrials; i++ {
		votes1, votes2 := SimulateOneElection(polls, electoralVotes, marginOfError)

		if votes1 > votes2 {
			winCount1++
		} else if votes2 > votes1 {
			winCount2++
		} else {
			tieCount++
		}
	}

	prob1 := float64(winCount1) / float64(numTrials)
	prob2 := float64(winCount2) / float64(numTrials)
	probTie := float64(tieCount) / float64(numTrials)

	return prob1, prob2, probTie
}

func SimulateOneElection(polls map[string]float64, electoralVotes map[string]int, marginOfError float64) (int, int) {
	collegeVotes1 := 0
	collegeVotes2 := 0

	for state, pollingValue := range polls {
		numVotes := electoralVotes[state]
		adjustedPoll := AddNoise(pollingValue, marginOfError)

		if adjustedPoll >= 0.5 {
			collegeVotes1 += numVotes
		} else {
			collegeVotes2 += numVotes
		}
	}

	return collegeVotes1, collegeVotes2
}

func AddNoise(pollingValue, marginOfError float64) float64 {
	x := rand.NormFloat64()
	x /= 2.0
	x *= marginOfError
	return x + pollingValue
}
