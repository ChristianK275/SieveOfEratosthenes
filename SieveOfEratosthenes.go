package main

import (
	"fmt"
	"math"
)

/*
* Return the primes under n
 */
func primesUnder(n int) []int{
	if n <= 2 {
		return make([]int, 0)
	}

	var notPrimes = make([]bool, n)

	notPrimes[0] = true
	notPrimes[1] = true

	var sqrtN = int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if !notPrimes[i] {
			for j := i * i; j < n; j += i {
				notPrimes[j] = true
			}
		}
	}

	var out []int
	for value, notPrime := range notPrimes {
		if !notPrime {
			out = append(out, value)
		}
	}

	return out
}

func main() {
	fmt.Println("GoGo")

	var primes = primesUnder(8)
	fmt.Println(primes)
}