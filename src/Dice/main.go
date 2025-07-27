package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	// fmt.Println("Parallel and concurrent ")
	// n := 100000000
	// start := time.Now()
	// Factorial(n)
	// elapsed := time.Since(start)
	// log.Printf("Multi processors took %s", elapsed)
	// runtime.GOMAXPROCS(1)
	// fmt.Println("Rolling dice.")

	// start2 := time.Now()
	// Factorial(n)
	// elapsed2 := time.Since(start2)
	// log.Printf("Multi processors took %s", elapsed2)
	// runtime.GOMAXPROCS(2)
	// fmt.Println("Rolling dice.")

	CrapsHouseEdgeTiming()
	fmt.Println("Program finished")
}

func CrapsHouseEdgeTiming() {
	numTrials := 1000000000
	numProcs := runtime.NumCPU()
	start := time.Now()
	ComputeCrapsHouseEdge(numTrials)
	elapsed := time.Since(start)
	start2 := time.Now()
	ComputeCrapsHouseEdgeMultiProc(numTrials, numProcs)
	elapsed2 := time.Since(start2)
	fmt.Println(elapsed, elapsed2)
}

func ParallelFactorial() {
	n := 10
	c := make(chan int)

	go Perm(1, n/2, c)
	go Perm(n/2, n, c)

	p1 := <-c
	p2 := <-c

	fmt.Println("n1 is", p1*p2)
}

func ChannelBasics() {
	c := make(chan string)

	SayHi(c)

	msg := <-c

	fmt.Println(msg)
}

func SayHi(c chan string) {
	c <- "Hello"
}

func Perm(k, n int, c chan int) {
	p := 1

	for i := k; i < n; i++ {
		p *= i
	}

	c <- p
}

func Factorial(n int) {
	p := 1
	for i := 1; i <= n; i++ {
		fmt.Println(p)
		p *= i
	}
}

// PlayCrapsOnce simulates one game of craps.
// Input: none
// Output: true if the game is a win and false if it's a loss
func PlayCrapsOnce() bool {
	firstRoll := SumDice(2)
	if firstRoll == 7 || firstRoll == 11 {
		// winner!
		return true
	} else if firstRoll == 2 || firstRoll == 3 || firstRoll == 12 {
		return false //loser! :(
	} else {
		// keep rolling until you hit a 7 or your original roll
		for { // while forever
			newRoll := SumDice(2)
			if newRoll == firstRoll {
				return true // winner!
			} else if newRoll == 7 {
				return false //loser! :(
			}
		}
	}
}

// ComputeCrapsHouseEdge estimates the "house edge" of craps over multiple simulations.
// Input: an integer corresponding to the number of simulations.
// Output: house edge of craps (average amount won or lost over the number of simulations per unit of currency)
func ComputeCrapsHouseEdgeMultiProc(numTrials int, numProcs int) float64 {
	count := 0 // will keep track of amount won (+) or lost (-)
	c := make(chan int)

	for i := 0; i < numProcs-1; i++ {
		go TotalWinOneProc(numTrials/numProcs, c)
	}

	go TotalWinOneProc(numTrials/numProcs+numTrials%numProcs, c)

	for i := 0; i < numProcs; i++ {
		count += <-c
	}

	// //run n trials and update count accordingly
	// for i := 0; i < numTrials; i++ {
	// 	//play the game
	// 	outcome := PlayCrapsOnce()
	// 	if outcome {
	// 		//winner!
	// 		count++
	// 	} else {
	// 		//loser :(
	// 		count--
	// 	}
	// }

	return float64(count) / float64(numTrials)
}

func ComputeCrapsHouseEdge(numTrials int) float64 {
	count := 0 // will keep track of amount won (+) or lost (-)}

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

func TotalWinOneProc(numTrials int, c chan int) {
	count := 0

	for i := 0; i <= numTrials; i++ {
		outcome := PlayCrapsOnce()
		if outcome {
			count++
		} else {
			count--
		}
	}

	c <- count
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
