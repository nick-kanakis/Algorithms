package heapsort

import (
	"fmt"
	"testing"
)

func TestPrivateReturnMaxChildPosition(t *testing.T) {

	unsortedArray := []int{10, 4, 2, 8, 15, 23}
	returnedValue := returnMaxChildPosition(unsortedArray, 0)
	if 1 != returnedValue {
		fmt.Println("Position: ", returnedValue)
		t.Error()
	}

	returnedValue = returnMaxChildPosition(unsortedArray, 1)
	if 4 != returnedValue {
		fmt.Println("Position: ", returnedValue)
		t.Error()
	}

	returnedValue = returnMaxChildPosition(unsortedArray, 2)
	if 5 != returnedValue {
		fmt.Println("Position: ", returnedValue)
		t.Error()
	}

	returnedValue = returnMaxChildPosition(unsortedArray, 3)
	if -1 != returnedValue {
		fmt.Println("Position: ", returnedValue)
		t.Error()
	}

	returnedValue = returnMaxChildPosition(unsortedArray, 4)
	if -1 != returnedValue {
		fmt.Println("Position: ", returnedValue)
		t.Error()
	}
}

func TestSort(t *testing.T) {

	heap := []int{4, 10, 3, 5, 1}
	fmt.Println(heap)
	Sort(heap)
	fmt.Println(heap)

	for i := 0; i < len(heap)-2; i++ {
		if heap[i] > heap[i+1] {
			t.Error()
		}
	}
}

func TestPrivateBuildMaxHeap(t *testing.T) {

	heap := []int{4, 10, 3, 5, 1}
	fmt.Println(heap)
	buildMaxHeap(heap)
	fmt.Println(heap)

	for i := len(heap) - 1; i >= 0; i-- {
		parentPosition := (i - 1) / 2
		if heap[i] > heap[parentPosition] {
			t.Error()
		}
	}
}
