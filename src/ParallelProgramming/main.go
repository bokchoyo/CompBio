package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Parallel and concurrent programming.")
	CrapsHouseEdgeTiming()
	fmt.Println("Program finished!")
}

func CrapsHouseEdgeTiming() {
	numTrials := 10000000
	start := time.Now()
	ComputeCrapsHouseEdge(numTrials)
	elapsed := time.Since(start)
	log.Printf("Serial house edge took %s", elapsed)
	start2 := time.Now()
	numProcs := runtime.NumCPU()
	ComputeCrapsHouseEdgeMultiProc(numTrials, numProcs)
	elapsed2 := time.Since(start2)
	log.Printf("Parallel house edge took %s", elapsed2)

}

func ParallelFactorial() {
	n := 1000000000
	c := make(chan int)
	go Perm(1, n/2, c)
	go Perm(n/2, n, c)

	p1 := <-c
	p2 := <-c
	fmt.Println("n! is", p1*p2)
}

// Perm
// Input: two integers k and n
// Output: product k*(k+1)*...(n-1)
func Perm(k, n int, c chan int) {
	p := 1

	// range between k and n, multiplying each one by p
	for i := k; i < n; i++ {
		p *= i
	}

	c <- p
}

func ChannelBasics() {
	// a channel is like a telephone wire
	// both parties have to be present for a message to be exchanged
	c := make(chan string)

	// this channel is "synchronous"; when you send or receive a message into our out of the channel, the rest of the function "blocks" or refuses to run until the other party is present

	go SayHi(c) // this function starts running in parallel

	msg := <-c // receive from channel and place it into a variable
	fmt.Println(msg)
}

func SayHi(c chan string) {
	// you could put 10,000 lines of code here
	c <- "Hello" // put this string into the channel
	// this function now BLOCKS
	// but there's nothing else in the function
}

// PrintFactorials takes an integer n as input.
// It prints 0!, 1!, 2!, 3!, ..., n!
func PrintFactorials(n int) {
	p := 1

	for i := 1; i <= n+1; i++ {
		fmt.Println(p)
		p *= i
	}
}

func Intro() {
	fmt.Println("This computer has", runtime.NumCPU(), "cores available.")

	n := 1000000000
	start := time.Now()
	Factorial(n)
	elapsed := time.Since(start)
	log.Printf("Multi processors took %s", elapsed)

	// I could set the number of processors to 1
	runtime.GOMAXPROCS(1)

	start2 := time.Now()
	Factorial(n)
	elapsed2 := time.Since(start2)
	log.Printf("One processor took %s", elapsed2)
}

func Factorial(n int) int {
	p := 1

	for i := 1; i <= n; i++ {
		p *= i
	}

	return p
}
