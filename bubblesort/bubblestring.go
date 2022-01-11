// bubble sort algorithm
package bubblesort

func BubbleStringSort(names []string) []string {
	var i int
	for i = 0; i < len(names); i++ {
		if !StringSweep(names, i) {
			break
		}
	}
	return names
}

func StringSweep(names []string, previousPass int) bool {
	var n int = len(names)
	var firstIndex int = 0
	var secondIndex int = 1
	var swapped = false

	for secondIndex < (n - previousPass) {
		var firstName string = names[firstIndex]
		var secondName string = names[secondIndex]

		if greaterThan(firstName, secondName) {
			names[firstIndex] = secondName
			names[secondIndex] = firstName
			swapped = true
		}

		firstIndex++
		secondIndex++
	}
	return swapped
}

func greaterThan(a, b string) bool {
	return a > b
}
