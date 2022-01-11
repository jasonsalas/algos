package bubblesort

import (
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {
	assertEquality := func(t testing.TB, got, want interface{}) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v | want: %v", got, want)
		}
	}

	t.Run("slice of integers", func(t *testing.T) {
		numbers := []int{4, 7, 0, 9, 1, 2}
		got := BubbleSort(numbers)
		want := []int{0, 1, 2, 4, 7, 9}

		assertEquality(t, got, want)
	})

	t.Run("slice of strings", func(t *testing.T) {
		names := []string{"Jason", "Sabrina", "Marie", "Ken", "Thomas", "Asha", "Brittany"}
		got := BubbleStringSort(names)
		want := []string{"Asha", "Brittany", "Jason", "Ken", "Marie", "Sabrina", "Thomas"}

		assertEquality(t, got, want)
	})
}
