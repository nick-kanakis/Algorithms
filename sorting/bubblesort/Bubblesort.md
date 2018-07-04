# Bubble Sort  Theory

## Basic concepts

Bubble Sort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in wrong order.
Although the algorithm is simple, it is too slow and impractical for most problems.

**Algorithm**:

1) Until no swap is performed repeat:
    1) Iterate the array 
    2) If Array[i-1] > Array[i] swap them

## Time Complexity

**Worst Case**

`O(n^2)`

**Average Case**

`O(n^2)`

**Best Case**

`O(n)`

When the list is already sorted (best-case), the complexity of bubble sort is only O(n). 

## Space Complexity

`O(1)`

##  Advantages.

The only significant advantage that bubble sort has over most other implementations, even quicksort, but not insertion sort,
is that the ability to detect that the list is sorted efficiently is built into the algorithm.
When the list is already sorted (best-case), the complexity of bubble sort is only O(n). 
For example, if you have a big array and you now that only a small portion of elements are out of order
bubblesort is a good candidate.
By contrast, most other algorithms, even those with better average-case complexity, perform their entire sorting process 
on the set and thus are more complex. 
However, not only does insertion sort have this mechanism too, but it also performs better on a list that is substantially sorted.
 
Bubble sort should be avoided in the case of large collections. It will not be efficient in the case of a reverse-ordered collection

Bubble sort is not stable!



