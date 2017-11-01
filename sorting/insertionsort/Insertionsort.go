package insertionsort


func Sort(initialArray []int){

	for i:= 1; i< len(initialArray); i++ {
		if initialArray[i] >= initialArray[i-1] {
			continue
		}
		for j:=i ; j>=1; j-- {
			if initialArray[j] < initialArray[j-1] {
				initialArray[j], initialArray[j-1] = initialArray[j-1], initialArray[j]
			} else {
				break
			}
		}
	}
}