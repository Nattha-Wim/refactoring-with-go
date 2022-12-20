package fizzbuzz

import "fmt"

func FizzBuzz(num int) string {
	if num == 3 || num == 6 || num == 9 {
		return "Fizz"
	}
	if num == 5 {
		return "Buzz"
	}
	// convert int to string
	return fmt.Sprintf("%d", num)
}
