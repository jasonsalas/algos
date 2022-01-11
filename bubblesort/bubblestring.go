// bubble sort algorithm
package bubblesort

func BubbleStringSort(names []string) {
	var i int
	for i = 0; i < len(names); i++ {
		if !StringSweep(names, i) {
			return
		}
	}
}

func StringSweep(names []string, previousPass int) bool {
	var n int = len(names)
	var firstIndex int = 0
	var secondIndex int = 1
	var swapped = false

	for secondIndex < (n - previousPass) {
		var firstNumber string = names[firstIndex]
		var secondNumber string = names[secondIndex]

		if firstNumber > secondNumber {
			names[firstIndex] = secondNumber
			names[secondIndex] = firstNumber
			swapped = true
		}

		firstIndex++
		secondIndex++
	}
	return swapped
}
