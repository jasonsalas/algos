package main

import (
	"fmt"

	"github.com/jasonsalas/algos/bubblesort"
)

func main() {
	var numbers []int = []int{4, 5, 3, 2, 7, 1, 0}
	fmt.Println("unsorted:", numbers)
	bubblesort.BubbleSort(numbers)
	fmt.Println("sorted:", numbers)
}
