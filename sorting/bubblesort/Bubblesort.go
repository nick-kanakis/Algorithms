package bubblesort

func Sort(initialArray []int) {

	swapped := true

	for swapped == true {
		swapped = false
		for i := 1; i < len(initialArray); i++ {
			if initialArray[i-1] < initialArray[i] {
				initialArray[i-1], initialArray[i] = initialArray[i], initialArray[i-1]
				swapped = true
			}
		}
	}
}
