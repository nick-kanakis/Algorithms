# Merge Sort Theory

## Basic concepts

![alt text](http://www.geeksforgeeks.org/wp-content/uploads/Merge-Sort-Tutorial.png "MergeSort")

Merge Sort is a Divide and Conquer algorithm. 
It divides input array in two halves, calls itself for the two halves and then merges the two sorted halves.

**Algorithm**:

1. Find the middle point to divide the array into two halves:
2. Call mergeSort for first half
3. Call mergeSort for second half
4. Merge the two halves sorted in step 2 and 3

Merge Step (Given 2 sorted Arrays)

>	Have we reached the end of any of the arrays?
>
>	No:
>
>		Compare current elements of both arrays
>
>		Copy smaller element into sorted array
>
>		Move pointer of element containing smaller element
>
>	Yes:
>
>		Copy all remaining elements of non-empty array

## Time Complexity

**Worst Case**

`O(nlogn)`

**Average Case**

`O(nlogn)`

**Best Case**

`O(nlogn)`

The entire input must be iterated through, and this must occur `O(log(n))` times (the input can only be halved `O(log(n))` times). 
n items iterated log(n) times gives `O(n log(n))`.

## Space Complexity

`O(n)`

## MergeSort Advantages.

*Merge sort* is often the best choice for sorting a linked list: 

In case of linked lists the case is different mainly due to difference in memory allocation of arrays and linked lists. 
Unlike arrays, linked list nodes may not be adjacent in memory. 

In linked list, we can insert items in the middle in `O(1)` extra space and `O(1)` time. 
Therefore merge operation of *merge sort* can be implemented without extra space for linked lists.
Unlike arrays, we can not do random access in linked list. 

*Quick Sort* requires a lot of this kind of access. 
In linked list to access i'th index, we have to travel each and every node from the head to i'th node as 
we don’t have continuous block of memory. Therefore, the overhead increases for *quick sort*. 

*Merge sort* accesses data sequentially and the need of random access is low.
The slow random-access performance of a linked list makes some other algorithms (such as *quicksort*) perform poorly, 
and others (such as *heapsort*) completely impossible.

In Java, the `Arrays.sort()` methods use merge sort or a tuned *quicksort* depending on the datatypes







