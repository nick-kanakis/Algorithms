# Quik Sort Theory

## Basic concepts

![alt text](http://www.geeksforgeeks.org/wp-content/uploads/gq/2014/01/QuickSort2.png "Quicksort")

Like Merge Sort, QuickSort is a Divide and Conquer algorithm. 
It picks an element as pivot and partitions the given array around the picked pivot.
In efficient implementations it is not a stable sort, meaning that the relative order of equal sort items is not preserved. 
Quicksort can operate in-place on an array, requiring small additional amounts of memory to perform the sorting.

Mathematical analysis of quicksort shows that, on average, the algorithm takes O(n log n) comparisons to sort n items. 
In the worst case, it makes O(n^2) comparisons, though this behavior is rare.

The key process in quickSort is partition(). Target of partitions is, given an array and an element x of array as pivot, 
put x at its correct position in sorted array and put all smaller elements (smaller than x) before x, 
and put all greater elements (greater than x) after x. All this should be done in linear time

**Algorithm**:

1) Choose the highest index value has pivot.
2) *Partitioning*: reorder the array so that all elements with values less than the pivot come before the pivot, 
while all elements with values greater than the pivot come after it (equal values can go either way). 
After this partitioning, the pivot is in its final position. This is called the partition operation.
    1) Create a leftPointer, this is a pointer that will move from left to right.
    2) for each element in the array
        1) if current element < of the pivot element 
        2) swap **current element with element pointed by the leftPointer**
        3) Move leftPointer to the right
    3) swap pivot with element in the leftPointer
3) Recursively apply the above steps to the sub-array of elements with smaller values and separately
 to the sub-array of elements with greater values

## Time Complexity

**Worst Case**

`O(n^2)`

**Average Case**

`O(nlogn)`

To sort an array of n distinct elements, quicksort takes `O(n log n)` time in expectation, 
averaged over all n! permutations of n elements with equal probability

## Space Complexity

`O(logn)`

The in-place version of quicksort has a space complexity of `O(log n)`, 
even in the worst case, when it is carefully implemented using the following strategies:

- In-place partitioning is used. This unstable partition requires O(1) space.
- After partitioning, the partition with the fewest elements is (recursively) sorted first, 
requiring at most O(log n) space. Then the other partition is sorted using tail recursion or iteration, which doesn't add to the call stack.





