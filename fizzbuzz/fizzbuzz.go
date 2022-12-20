package fizzbuzz

import "fmt"

func FizzBuzz(num int) string {
	if isFizzBuzz(num) {
		return "FizzBuzz"
	}
	if isFizz(num) {
		return "Fizz"
	}
	if isBuzz(num) {
		return "Buzz"
	}
	// convert int to string
	return fmt.Sprintf("%d", num)
}

func isFizz(num int) bool {
	return num%3 == 0
}

func isBuzz(num int) bool {
	return num%5 == 0
}

func isFizzBuzz(num int) bool {
	return num%15 == 0
}

// go test -coverprofile c.out
// go tool cover -html=c.out
