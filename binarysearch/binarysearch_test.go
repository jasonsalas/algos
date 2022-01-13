package binarysearch

import (
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := BinarySearch(nums, 4)
	want := 3

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v | want: %v", got, want)
	}
}
