package main

import "math/rand"

func PlayCrapsOnce() bool {
	firstRoll := SumDice(2)
	switch firstRoll {
	case 7, 11:
		return true
	case 2, 3, 12:
		return false
	default:
		for {
			newRoll := SumDice(2)
			switch newRoll {
			case firstRoll:
				return true
			case 7:
				return false
			}
		}
	}
}

// ComputeCrapsHouseEdge estimates the "house edge" of craps over multiple simulations.
// Input: an integer corresponding to the number of simulations.
// Output: house edge of craps (average amount won or lost over the number of simulations per unit of currency)
func ComputeCrapsHouseEdge(numTrials int) float64 {
	count := 0 // will keep track of amount won (+) or lost (-)

	//run n trials and update count accordingly
	for i := 0; i < numTrials; i++ {
		//play the game
		outcome := PlayCrapsOnce()
		if outcome {
			//winner!
			count++
		} else {
			//loser :(
			count--
		}
	}

	return float64(count) / float64(numTrials)
}

// SumDice simulates the process of summing n dice.
// Input: an integer numDice
// Output: the sum of numDice simulated dice
func SumDice(numDice int) int {
	sum := 0

	for i := 0; i < numDice; i++ {
		sum += RollDie()
	}

	return sum
}

// RollDie
// Input: none
// Output: a pseudorandom integer between 1 and 6, inclusively.
func RollDie() int {
	return rand.Intn(6) + 1
}

// SumTwoDice
// Input: none
// Output: the simulated sum of two dice (between 2 and 12).
func SumTwoDice() int {
	return RollDie() + RollDie()
	/*
		roll := rand.Float64()
		if roll < 1.0/36.0 {
			return 2
		} else if roll < 3.0/36.0 { // we know that roll > 1/36
			return 3
		} else if roll < 6.0/36.0 {
			return 4
		} // etc.
	*/
}
