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
	pivot := initialArray[len(initialArray)-1]
	leftPointer := 0
	rightPointer := len(initialArray) - 2

	for {
		if len(initialArray) == 1 {
			return 0
		}

		for initialArray[leftPointer] < pivot {
			leftPointer++
		}

		for initialArray[rightPointer] > pivot {
			rightPointer--
		}

		if leftPointer >= rightPointer {
			break
		} else {
			initialArray[leftPointer], initialArray[rightPointer] = initialArray[rightPointer], initialArray[leftPointer]
		}
	}

	initialArray[leftPointer], initialArray[len(initialArray)-1] = initialArray[len(initialArray)-1], initialArray[leftPointer]
	return leftPointer;
}
