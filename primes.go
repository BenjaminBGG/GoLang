/*
100Primes in GoLang -- Prints first 100 primes numbers to console
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var num int = 1
	for i := 0; i <= 100; {
		if IsPrime(num) {
			fmt.Println(num)
			i++
		}
		num++
	}

}

// IsPrime - return true if number is prime
func IsPrime(num int) bool {
	if num < 4 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
