package mergesort

// MergeSort is used to sort an array of integer
func MergeSort(input []int) []int {
	lengthInput := len(input)
	if lengthInput <= 1 {
		return input
	}
	firstHalf := MergeSort(input[:lengthInput/2])
	secondHalf := MergeSort(input[lengthInput/2:])
	var firstPointer, secondPointer, outPointer int
	output := make([]int, lengthInput)
	for firstPointer < len(firstHalf) {
		for secondPointer < len(secondHalf) && secondHalf[secondPointer] < firstHalf[firstPointer] {
			output[outPointer] = secondHalf[secondPointer]
			secondPointer++
			outPointer++
		}
		output[outPointer] = firstHalf[firstPointer]
		firstPointer++
		outPointer++
	}
	for secondPointer < len(secondHalf) {
		output[outPointer] = secondHalf[secondPointer]
		secondPointer++
		outPointer++
	}
	return output
}
