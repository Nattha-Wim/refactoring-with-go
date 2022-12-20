package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {

	got := FizzBuzz(input)

	if got != want {
		t.Errorf("FizzBuzz(%d) = %q, want %q", input, got, want)
	}

func TestInput2ShouldBe2(t *testing.T) {
	want := "2"
	input := 2

	got := FizzBuzz(input)

	if got != want {
		t.Errorf("FizzBuzz(%d) = %q, want %q", input, got, want)
	}
}

func TestInput3ShouldBeFizz(t *testing.T) {
	want := "Fizz"
	input := 3

	got := FizzBuzz(input)

	if got != want {
		t.Errorf("FizzBuzz(%d) = %q, want %q", input, got, want)
	}
}

func TestInput4ShouldBe4(t *testing.T) {
	want := "4"
	input := 4

	got := FizzBuzz(input)

	if got != want {
		t.Errorf("FizzBuzz(%d) = %q, want %q", input, got, want)
	}
}
