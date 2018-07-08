package selectionsort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	initialArray := []int{9, 7, 6, 15, 16, 5, 10, 11}

	fmt.Println(initialArray)
	Sort(initialArray)
	fmt.Println(initialArray)

	for i := 0; i < len(initialArray)-2; i++ {
		if initialArray[i] > initialArray[i+1] {
			t.Error()
		}
	}
}
