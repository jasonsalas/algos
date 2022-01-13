package main

import (
	"fmt"

	"github.com/jasonsalas/algos/binarysearch"
	"github.com/jasonsalas/algos/bubblesort"
)

// func main() {
// 	var names []string = []string{
// 		"dog", "cat", "alligator", "cheetah", "rat",
// 		"moose", "cow", "mink", "porcupine", "dung beetle",
// 		"camel", "steer", "bat", "hamster", "horse", "colt",
// 		"bald eagle", "frog", "rooster",
// 	}
// 	var result []string
// 	fmt.Println("unsorted:", names)
// 	result = bubblesort.BubbleStringSort(names)
// 	fmt.Println("sorted:", result)
// }

func main() {
	lookingFor := 13
	numbers := []int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}

	fmt.Println("sorting my list:", numbers)
	sortedList := bubblesort.BubbleSort(numbers)
	fmt.Println("sorted!", sortedList)

	fmt.Println("looking for", lookingFor, "in the sorted list:", sortedList)

	index := binarysearch.BinarySearch(sortedList, lookingFor)
	if index >= 0 {
		fmt.Println("found the number", lookingFor, "at index position", index)
	} else {
		fmt.Println("didn't find the number", lookingFor, ":(")
	}
}
