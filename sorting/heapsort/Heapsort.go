package heapsort

func Sort(unsortedArray []int) {

	//Create a Max Heap
	buildMaxHeap(unsortedArray)

	originalSize := len(unsortedArray);

	//Swap root element with last child of unsorted port
	//and re-heapify to fix the problem caused by the swap.
	for i := 0; i < originalSize; i++ {
		lastChildPosition := len(unsortedArray) - 1
		unsortedArray[0], unsortedArray[lastChildPosition] = unsortedArray[lastChildPosition], unsortedArray[0]
		maxHeapify(unsortedArray[:lastChildPosition], 0)

		// the slice of [start: len(originalArray -1] is the new heap that we need for our next iteration.
		unsortedArray = unsortedArray[:lastChildPosition]
	}
}

// For  size(OriginalArray)/2 to 0
// Call heapify, that will make the sub-tree under i a heap.
func buildMaxHeap(unsortedArray []int) {
	for i := len(unsortedArray) / 2; i >= 0; i-- {
		maxHeapify(unsortedArray, i)
	}
}

// 		Find the maxChild Of of the currentParent
//		If you need to swap elements
//			swap and call again the method for the maxChild in order to fix
//			any problems from the previous step.
func maxHeapify(unsortedArray []int, i int) {
	maxChildPosition := returnMaxChildPosition(unsortedArray, i)

	if maxChildPosition != -1 && unsortedArray[i] < unsortedArray[maxChildPosition] {
		unsortedArray[i], unsortedArray[maxChildPosition] = unsortedArray[maxChildPosition], unsortedArray[i]
		maxHeapify(unsortedArray, maxChildPosition)
	}
}

func returnMaxChildPosition(array []int, parentPosition int) int {

	length := len(array)
	leftChildPosition := parentPosition*2 + 1
	rightChildPosition := leftChildPosition + 1

	//No children, if there is no left child it can not have a right one.
	if leftChildPosition >= length {
		return -1
	}

	//Only left children
	if rightChildPosition >= length {
		return leftChildPosition
	}

	//2 Children
	if array[leftChildPosition] > array[rightChildPosition] {
		return leftChildPosition
	} else {
		return rightChildPosition
	}

}
