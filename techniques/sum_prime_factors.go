package techniques

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Function to find the prime factors of a given number
func primeFactors(n int) []int {
	factors := []int{}
	for n%2 == 0 {
		factors = append(factors, 2)
		n = n / 2
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n = n / i
		}
	}

	if n > 2 {
		factors = append(factors, n)
	}
	return factors
}

// Function to sum the prime factors
func sumPrimeFactors(factors []int) int {
	sum := 0
	for _, factor := range factors {
		sum += factor
	}
	return sum
}

func main() {
	n := 60 // Example input

	for {
		fmt.Printf("Number: %d\n", n)
		factors := primeFactors(n)
		if len(factors) == 0 || (len(factors) == 1 && factors[0] == n) {
			fmt.Println("No more prime factors can be found.")
			break
		}
		fmt.Printf("Prime factors: %v\n", factors)
		n = sumPrimeFactors(factors)
	}
}
