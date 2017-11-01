package quicksort

func Sort(initialArray []int) {

	if len(initialArray) <= 1 {
		return
	}
	pivotPointer := partition(initialArray)
	Sort(initialArray[:pivotPointer])
	Sort(initialArray[pivotPointer+1:])
}

func partition(initialArray []int) int {
	pivot := len(initialArray)-1
	leftPointer := 0

	// Pile elements smaller than the pivot on the left
	for i:= range initialArray{
		if initialArray[i] < initialArray[pivot] {
			initialArray[i], initialArray[leftPointer] = initialArray[leftPointer], initialArray[i]
			leftPointer++
		}
	}
	// Place the pivot after the last smaller element
	initialArray[leftPointer], initialArray[pivot] = initialArray[pivot], initialArray[leftPointer]
	return leftPointer;
}
