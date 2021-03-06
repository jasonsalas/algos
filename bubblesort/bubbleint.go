package bubblesort

func BubbleSort(numbers []int) []int {
	var i int
	for i = 0; i < len(numbers); i++ {
		if !Sweep(numbers, i) {
			break
		}
	}
	return numbers
}

func Sweep(numbers []int, previousPass int) bool {
	var n int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1
	var swapped = false

	for secondIndex < (n - previousPass) {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]

		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			swapped = true
		}

		firstIndex++
		secondIndex++
	}
	return swapped
}
