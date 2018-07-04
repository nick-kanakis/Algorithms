package bubblesort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	initialArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(initialArray)
	Sort(initialArray)
	fmt.Println(initialArray)

	for i := 0; i <= len(initialArray)-2; i++ {
		if initialArray[i] < initialArray[i+1] {
			t.Error()
		}
	}

}
