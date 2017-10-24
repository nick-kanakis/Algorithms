package heapsort

func Sort(unsortedArray []int) {

	//Create a Max Heap
	buildMaxHeap(unsortedArray)
	//Swap and remove root element
	originalSize := len(unsortedArray);
	for i := 0; i < originalSize; i++ {
		lastChildPosition := len(unsortedArray) - 1
		unsortedArray[0], unsortedArray[lastChildPosition] = unsortedArray[lastChildPosition], unsortedArray[0]
		maxHeapify(unsortedArray[:lastChildPosition], 0)
		unsortedArray = unsortedArray[:lastChildPosition]
	}
}

func buildMaxHeap(unsortedArray []int) {
	for i := len(unsortedArray) / 2; i >= 0; i-- {
		maxHeapify(unsortedArray, i)
	}
}

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

	//No children
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
