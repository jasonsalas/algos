package main

import (
	"fmt"

	"github.com/jasonsalas/algos/bubblesort"
)

func main() {
	var numbers []int = []int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}
	fmt.Println("unsorted:", numbers)
	bubblesort.BubbleSort(numbers)
	fmt.Println("sorted:", numbers)
}
