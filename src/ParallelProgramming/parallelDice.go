package main

// ComputeCrapsHouseEdgeMultiProc estimates the "house edge" of craps over multiple simulations.
// Input: an integer corresponding to the number of simulations, and an integer corresponding to the number of "processors" (workers) that we're dividing the algorithm over.
// Output: house edge of craps (average amount won or lost over the number of simulations per unit of currency)
func ComputeCrapsHouseEdgeMultiProc(numTrials, numProcs int) float64 {
	count := 0          // will keep track of amount won (+) or lost (-)
	c := make(chan int) // this is the phone we're using

	//divide numTrials equally over numProcs
	for i := 0; i < numProcs-1; i++ {
		go TotalWinOneProc(numTrials/numProcs, c)
	}
	// the final goroutine is going to get the remainder as well
	go TotalWinOneProc(numTrials/numProcs+numTrials%numProcs, c)

	// get all numProcs values from channel
	// this is you calling me to tell me your "winnings"
	for i := 0; i < numProcs; i++ {
		count += <-c
	}

	return float64(count) / float64(numTrials)
}

// TotalWinOneProc
// Input: number of trials (integer) to run a craps simulation for,
// and an integer channel.
// Output: (none), but put total winnings into channel after playing craps numTrials simulated times.
func TotalWinOneProc(numTrials int, c chan int) {
	count := 0

	for i := 0; i < numTrials; i++ {
		outcome := PlayCrapsOnce()
		if outcome {
			count++
		} else {
			//lost :(
			count--
		}
	}

	// place total winnings into channel
	c <- count
}
