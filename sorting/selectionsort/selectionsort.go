package selectionsort

func Sort(array []int) {

	for i:= 0; i < len(array); i++{
		minPtr := getMinPtr(array[i:]) + i
		array[i], array[minPtr] = array[minPtr], array[i]
	}
}

func getMinPtr(unsortedArray []int) int {

	minVal := unsortedArray[0]
	minPtr := 0
	for  i:= 1; i < len(unsortedArray); i++{
		if unsortedArray[i] < minVal{
			minVal = unsortedArray[i]
			minPtr = i
		}
	}

	return minPtr
}