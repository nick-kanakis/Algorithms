#  Theory

## Basic concepts

![alt text](https://he-s3.s3.amazonaws.com/media/uploads/2888f5b.png "Selection sort")

Selection sort algorithm sorts an array by repeatedly finding the minimum element (considering ascending order) from unsorted part and putting it at the beginning. The algorithm maintains two subarrays in a given array.

1) The subarray which is already sorted.
2) Remaining subarray which is unsorted.

In every iteration of selection sort, the minimum element (considering ascending order) from the unsorted subarray is picked and moved to the sorted subarray.

**Algorithm**:

1) FirstPtrOfUnsortedPart points to the 0 at the beginning, it marks the first element of the unsorted part of the array and it moves by one to the right at each iteration
2) Find the minimum element of the unsorted part 
3) Replace FirstPtrOfUnsortedPart with minimum element and move FirstPtrOfUnsortedPart by one (to the right)
4) Repeat until FirstPtrOfUnsortedPart is equal to array size-1 (no more elements in the unsorted part)

## Time Complexity

**Worst Case**

`O(n^2)`

**Average Case**

`O(n^2)`

**Best Case**
`O(n^2)`


## Space Complexity

`O(1)`

##  Advantages.

- Selection sort almost always outperforms bubble sort.
- Selection sort is preferable to insertion sort in terms of number of writes
- It can be seen as an advantage for some real-time applications that selection sort will perform identically regardless of the order of the array, while insertion sort's running time can vary considerably.





