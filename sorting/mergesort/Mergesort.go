package mergesort

func Sort(initialArray []int) []int {

	if len(initialArray) <= 1 {
		return initialArray
	}

	middle := len(initialArray) / 2
	leftPartSorted := Sort(initialArray[:middle])
	rightPartSorted := Sort(initialArray[middle:])

	return merge2(leftPartSorted, rightPartSorted)
}

/*
	As long as both arrays are not empty
		if left is empty
			append everything remaining from the right to the result and return it
		if right is empty
			append everything remaining from the left to the result and return it

		if( first value of right > first value of left)
			append to the result the first value of right and shift  right by 1
		else
			append to the result the first value of left and shift left by 1
*/
func merge(leftPartSorted []int, rightPartSorted []int) []int {

	mergedArray := make([]int, 0, len(rightPartSorted)+len(leftPartSorted))

	for len(leftPartSorted) > 0 || len(rightPartSorted) > 0 {

		if len(leftPartSorted) == 0 {
			return append(mergedArray, rightPartSorted...)
		}

		if len(rightPartSorted) == 0 {
			return append(mergedArray, leftPartSorted...)
		}

		if leftPartSorted[0] >= rightPartSorted[0] {
			mergedArray = append(mergedArray, leftPartSorted[0])
			leftPartSorted = leftPartSorted[1:]
		} else {
			mergedArray = append(mergedArray, rightPartSorted[0])
			rightPartSorted = rightPartSorted[1:]
		}
	}
	return mergedArray
}

func merge2(right []int, left []int) []int {

	mergedResult := make([]int, 0, len(right)+len(left))

	rightPointer := 0
	leftPointer := 0

	for leftPointer < len(left) && rightPointer < len(right) {

		if right[rightPointer] > left[leftPointer] {
			mergedResult = append(mergedResult, right[rightPointer])
			leftPointer++
		} else {
			mergedResult = append(mergedResult, left[leftPointer])
			rightPointer++
		}
	}
	if rightPointer == len(right)-1 && leftPointer < len(left) {
		mergedResult = append(mergedResult, left[leftPointer:]...)

	} else if leftPointer == len(left)-1 && rightPointer < len(right) {
		mergedResult = append(mergedResult, right[rightPointer:]...)
	}
	return mergedResult
}
