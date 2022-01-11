// bubble sort algorithm
package bubblesort

func BubbleSort(numbers []int) {
	var i int
	for i = 0; i < len(numbers); i++ {
		Sweep(numbers, i)
	}
}

func Sweep(numbers []int, previousPass int) {
	var n int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1

	for secondIndex < (n - previousPass) {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]

		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}

		firstIndex++
		secondIndex++
	}

}
