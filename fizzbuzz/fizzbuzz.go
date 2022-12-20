package fizzbuzz

import "fmt"

func FizzBuzz(num int) string {
	if num%15 == 0 {
		return "FizzBuzz"
	}
	if num%3 == 0 {
		return "Fizz"
	}
	if num%5 == 0 {
		return "Buzz"
	}
	// convert int to string
	return fmt.Sprintf("%d", num)
}
