package insertionsort

func Sort(initialArray []int) {

	for i := 1; i < len(initialArray); i++ {
		ptr :=i
		for ptr>0{
			if initialArray[ptr] > initialArray[ptr-1]{
				initialArray[ptr], initialArray[ptr-1] = initialArray[ptr-1], initialArray[ptr]
				ptr--
			} else {
				break
			}
		}
	}
}
