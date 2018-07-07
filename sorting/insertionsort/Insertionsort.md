# Insertion Sort Theory

## Basic concepts

![alt text](http://www.geeksforgeeks.org/wp-content/uploads/gq/2013/03/insertion-sort.png "Insertion Sort")

Insertion sort is a simple sorting algorithm that builds the final sorted array (or list) one item at a time. 
It is much less efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort.


**Algorithm**:
1) If it is the first element, it is already sorted
2) Move to the next (left) element
3) Compare the selected element with all in the left side (sorted part) until you find the correct position (left element is the first smaller)
4) While element is smaller than the right on swap and compare with the new left
5) Repeat until there is no more elements left.

## Time Complexity

**Worst Case**

`O(n^2)`

The simplest worst case input is an array sorted in reverse order

**Average Case**

`O(n^2)`

**Best Case**

`O(n)`

The best case input is an array that is already sorted

## Space Complexity

`O(1)`

##  Advantages.

As we said in the beginning Insertion sort is not the most efficient algorithm. However, it provides several advantages

1) Simple implementation
2) Efficient for (quite) small data sets
3) Stable; i.e., does not change the relative order of elements with equal keys
4) In-place; i.e., only requires a constant amount O(1) of additional memory space
5) Online; i.e., can sort a list as it receives it

Insertion sort is better that bubblesort (another easy but slow sorting algo) because it stops iterating when the element is in the correct position.

Insertion sort will usually perform about half as many comparisons as selection sort, although it can perform just as many or far fewer depending on the order the array was in prior to sorting.
