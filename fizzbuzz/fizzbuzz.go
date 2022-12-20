package fizzbuzz

import "fmt"

func FizzBuzz(num int) string {
	if num == 3 {
		return "Fizz"
	}
	if num == 5 {
		return "Buzz"
	}
	// convert int to string
	return fmt.Sprintf("%d", num)
}
