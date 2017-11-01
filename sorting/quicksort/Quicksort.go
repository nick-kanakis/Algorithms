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
	ptrFirstElementGreaterThanPivot := 0

	// Pile elements smaller than the pivot on the left
	for i:= range initialArray{
		if initialArray[i] < initialArray[pivot] {
			initialArray[i], initialArray[ptrFirstElementGreaterThanPivot] = initialArray[ptrFirstElementGreaterThanPivot], initialArray[i]
			ptrFirstElementGreaterThanPivot++
		}
	}
	// Place the pivot after the last smaller element
	initialArray[ptrFirstElementGreaterThanPivot], initialArray[pivot] = initialArray[pivot], initialArray[ptrFirstElementGreaterThanPivot]
	return ptrFirstElementGreaterThanPivot;
}
