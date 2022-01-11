package main

import (
	"fmt"

	"github.com/jasonsalas/algos/bubblesort"
)

func main() {
	// var numbers []int = []int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}
	// var result []int
	var names []string = []string{
		"dog", "cat", "alligator", "cheetah", "rat",
		"moose", "cow", "mink", "porcupine", "dung beetle",
		"camel", "steer", "bat", "hamster", "horse", "colt",
		"bald eagle", "frog", "rooster",
	}
	var result []string
	fmt.Println("unsorted:", names)
	result = bubblesort.BubbleStringSort(names)
	fmt.Println("sorted:", result)
}
