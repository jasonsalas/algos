package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := BinarySearch(nums, 4)
	want := 3

	if got != want {
		t.Errorf("got: %d | want: %d", got, want)
	}
}
