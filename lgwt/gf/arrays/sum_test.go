package arrays

import "testing"

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
