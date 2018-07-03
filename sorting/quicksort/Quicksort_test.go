package quicksort

import (
	"fmt"
	"testing"
)

func TestPrivatePartition(t *testing.T) {

	initialArray := []int{3, 60, 23, 82, 100, 4, 50, 40}

	fmt.Println(initialArray)
	partition(initialArray)
	fmt.Println(initialArray)

	for i := 0; i < 3; i++ {
		if initialArray[i] > 40 {
			t.Error()
		}
	}

	if initialArray[3] != 40 {
		t.Error()
	}

	for i := 3; i < 7; i++ {
		if initialArray[i] < 40 {
			t.Error()
		}
	}
}

func TestSort(t *testing.T) {
	initialArray := []int{3, 60, 23, 82, 100, 4, 50, 40}

	fmt.Println(initialArray)
	Sort(initialArray)
	fmt.Println(initialArray)

	for i := 0; i < len(initialArray)-2; i++ {
		if initialArray[i] > initialArray[i+1] {
			t.Error()
		}
	}
}
