package insertionsort

import (
	"testing"
	"fmt"
)

func TestSort(t *testing.T) {
	initialArray := []int{9,8,7,6,5,4,3,2,1}

	fmt.Println(initialArray)
	Sort(initialArray)
	fmt.Println(initialArray)

	for i := 0; i < len(initialArray)-2; i++ {
		if initialArray[i] > initialArray[i+1] {
			t.Error()
		}
	}
}