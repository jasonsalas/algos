package binarysearch

import "fmt"

func BinarySearch(sortedList []int, searchValue int) int {
	// default bounds for the search subset
	var lowerIndex = 0
	var upperIndex = len(sortedList) - 1

	for lowerIndex <= upperIndex {
		midIndex := lowerIndex + (upperIndex-lowerIndex)/2
		midValue := sortedList[midIndex]
		fmt.Println("middle value:", midValue)

		if midValue == searchValue {
			return midIndex
		} else if midValue > searchValue {
			upperIndex = midIndex - 1 // use left split of the list
		} else if midValue < searchValue {
			lowerIndex = midIndex + 1 // use right split of the list
		}
	}
	return -1 // the search value doesn't exist in the list
}
