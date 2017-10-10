package mergesort

import (
	"testing"
	"fmt"
)

func TestSortWhenArrayIsEven(t *testing.T) {

	unsortedArray := []int{1, 5, 10, 4, 2, 6, 7, 3, 9, 8}

	 sortedArray := Sort(unsortedArray)

	for i := 0; i < len(sortedArray)-2; i++ {
		if sortedArray[i] < sortedArray[i+1] {
			fmt.Println(sortedArray)
			t.Error()
		}
	}
}


