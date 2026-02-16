package arrays

import "testing"

func TestSum(t *testing.T) {
	// given
	numbers := [5]int{1, 2, 3, 4, 5}
	want := 15

	// when
	got := Sum(numbers)

	// then
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
