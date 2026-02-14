package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Add(2, 2) should be 4", func(t *testing.T) {
		// given
		a, b := 2, 2
		expected := 4

		// when
		sum := Add(a, b)

		// then
		assertCorrectAdd(t, sum, expected)
	})
	t.Run("Add(1, 1) should be 2", func(t *testing.T) {
		// given
		a, b := 1, 1
		expected := 2

		// when
		sum := Add(a, b)

		// then
		assertCorrectAdd(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)

	fmt.Println(sum)
	// Output: 6
}

func assertCorrectAdd(t testing.TB, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
