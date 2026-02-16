package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		// given
		numbers := []int{1, 2, 3, 4, 5}
		want := 15

		// when
		got := Sum(numbers)

		// then
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		// given
		numbers := []int{1, 2, 3}
		want := 6

		// when
		got := Sum(numbers)

		// then
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	// given
	param := [][]int{{1, 2}, {0, 9}}
	want := []int{3, 9}

	// when
	got := SumAll(param...)

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
