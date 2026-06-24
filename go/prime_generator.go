package main

import (
	"fmt"
	"math"
	"time"
)

func isPrime(number int) bool {
	if number < 2 {
		return false
	}

	limit := int(math.Sqrt(float64(number)))

	for divisor := 2; divisor <= limit; divisor++ {
		if number%divisor == 0 {
			return false
		}
	}

	return true
}

func main() {
	startTime := time.Now()

	primes := []int{}
	candidate := 2

	for len(primes) < 10000 {
		if isPrime(candidate) {
			primes = append(primes, candidate)
		}
		candidate++
	}

	elapsed := time.Since(startTime)

	fmt.Println("Generated 10,000 primes")
	fmt.Println("Last prime:", primes[len(primes)-1])
	fmt.Printf("Elapsed time: %.4f seconds\n", elapsed.Seconds())
}
