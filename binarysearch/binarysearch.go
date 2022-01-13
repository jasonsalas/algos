package binarysearch

import "fmt"

func BinarySearch(sortedList []int, searchValue int) int {
	switch len(sortedList) {
	case 0:
		return -1
	case 1:
		if sortedList[0] == searchValue {
			return 0
		}
		return -1
	default:
		// default bounds for the search subset
		lowerIndex, upperIndex := 0, len(sortedList)-1

		for lowerIndex <= upperIndex {
			midIndex := (upperIndex + lowerIndex) / 2
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
}
