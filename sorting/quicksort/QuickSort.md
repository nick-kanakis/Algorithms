# Quik Sort Theory

## Basic concepts

![alt text](http://www.geeksforgeeks.org/wp-content/uploads/gq/2014/01/QuickSort2.png "Quicksort")

Like Merge Sort, QuickSort is a Divide and Conquer algorithm. 
It picks an element as pivot and partitions the given array around the picked pivot.
In efficient implementations it is not a stable sort, meaning that the relative order of equal sort items is not preserved. 
Quicksort can operate in-place on an array, requiring small additional amounts of memory to perform the sorting.

Mathematical analysis of quicksort shows that, on average, the algorithm takes O(n log n) comparisons to sort n items. We need n-1 compares and we divide the array logn times (the pivot is the median number or
at least not an extreme min/max number)

In the worst case, it makes O(n^2) comparisons, though this behavior is rare.
To have O(n^2) you must always select as pivot the max or min number so you have to divide the list n times
and down n-1 comparisons

The key process in quickSort is partition(). Target of partitions is, given an array and an element x of array as pivot, 
put x at its correct position in sorted array and put all smaller elements (smaller than x) before x, 
and put all greater elements (greater than x) after x. All this should be done in linear time

**Algorithm**:

1) Choose the highest index value has pivot.
2) *Partitioning*: reorder the array so that all elements with values less than the pivot come before the pivot, 
while all elements with values greater than the pivot come after it (equal values can go either way). 
After this partitioning, the pivot is in its final position. This is called the partition operation.
    1) Create a ptrFirstElementGreaterThanPivot that points to the first element of the array,
       this is a pointer that will move from left to right.
    2) for each element in the array
        1) if current element < of the pivot element (pivot is the last element of the array)
        2) swap **current element with element pointed by the ptrFirstElementGreaterThanPivot**
        3) move ptrFirstElementGreaterThanPivot to the right
    3) **swap pivot with element in the ptrFirstElementGreaterThanPivot**
3) Recursively apply the above steps to the sub-array of elements with smaller values and separately
 to the sub-array of elements with greater values

**Explanation of Partitioning step**:

The target is to have all elements smaller than the pivot to the left and all elements greater than the pivot to the right of the pivot.
This need to be done in place and at O(n) time.

(1) It is important to remember that all elements left of ptrFirstElementGreaterThanPivot are less than the pivot, 
in other words ptrFirstElementGreaterThanPivot points to the first element that is right of the pivot.

With each iteration we examine one by one the elements. 
If the current examining element is less than the pivot it needs to be on the left of the ptrFirstElementGreaterThanPivot. 
So we swap the element pointed by ptrFirstElementGreaterThanPivot with the current examining element
and move the ptrFirstElementGreaterThanPivot by 1 to the right. In other words we swap an element 
that needs to be left of the ptrFirstElementGreaterThanPivot with the first element of the right side of pivot and move the pointer to the next element. 
This changes the order of the elements, but still all elements left of the ptrFirstElementGreaterThanPivot are less than the pivot and all elements right to the pivot
(including the element pointed by ptrFirstElementGreaterThanPivot) are greater than the pivot.

Lastly we need to put the pivot in the right position. To do that we swap ptrFirstElementGreaterThanPivot element with pivot. This again changes the 
order of the array but (1) is still valid.

## Time Complexity

**Worst Case**

`O(n^2)`

In case you always select the min/max as the pivot, so you will divide O(n) times the array.

**Average Case**

`O(nlogn)`

In most cases you will not peek the min/max every time, so you will divide the array in O(logn) times

## Space Complexity

`O(logn)` (due to recursive calls)

The in-place version of quicksort has a space complexity of `O(log n)`, 
even in the worst case, when it is carefully implemented using the following strategies:

- In-place partitioning is used. This unstable partition requires O(1) space.
- After partitioning, the partition with the fewest elements is (recursively) sorted first, 
requiring at most O(log n) space. Then the other partition is sorted using tail recursion or iteration, which doesn't add to the call stack.





