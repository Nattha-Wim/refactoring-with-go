package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	want := "1"

	got := FizzBuzz(1)

	if got != want {
		t.Errorf("FizzBuzz(1) = %q, want %q", got, want)
	}
}

func TestINput2ShouldBe2(t *testing.T) {
	want := "2"

	got := FizzBuzz(2)

	if got != want {
		t.Errorf("FizzBuzz(2) = %q, want %q", got, want)
	}
}
