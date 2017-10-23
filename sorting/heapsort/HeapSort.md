# Heap Sort Theory

## Basic concepts
Heap sort is a comparison based sorting technique based on Binary Heap data structure.

The heapsort algorithm involves preparing the list by first turning it into a max heap. 
The algorithm then repeatedly swaps the first value of the list with the last value, 
decreasing the range of values considered in the heap operation by one,
and sifting the new first value into its position in the heap. 
This repeats until the range of considered values is one value in length

**Algorithm**:

1) Build a max heap from the input data. This needs `O(n)` operations.
2) At this point, the largest item is stored at the root of the heap. 
   Replace it with the last item of the heap followed by reducing the size of heap by 1. 
   Finally, heapify the root of tree. This need `O(longn)`
3) Repeat above steps while size of heap is greater than 1.
   
## Time Complexity

**Worst Case**

`O(nlogn)`

From the above algorithm we see that the step of creating the initial heap takes `O(n)`, will the heapify function
takes `O(longn)` and is called `n` times so:

`O(nlongn + n) = O(nlogn)`

## Space Complexity

`O(1)`

##  Advantages.

Heapsort primarily competes with quicksort, another very efficient general purpose nearly-in-place comparison-based sort algorithm.

Quicksort is typically somewhat faster due to some factors, but the worst-case running time for quicksort is O(n2), 
which is unacceptable for large data sets and can be deliberately triggered given enough knowledge of the implementation, creating a security risk.
Because of the O(n log n) upper bound on heapsort's running time and constant upper bound on its auxiliary storage, 
embedded systems with real-time constraints or systems concerned with security often use heapsort, such as the Linux kernel.
Heapsort also competes with merge sort, which has the same time bounds. Merge sort requires Ω(n) auxiliary space, 
but heapsort requires only a constant amount. Heapsort typically runs faster in practice on machines with small or slow data caches, 
and does not require as much external memory. On the other hand, merge sort has several advantages over heapsort:

1) Merge sort on arrays has considerably better data cache performance, 
   often outperforming heapsort on modern desktop computers because merge sort frequently accesses contiguous memory locations 
   (good locality of reference); heapsort references are spread throughout the heap.
2) Heapsort is not a stable sort; merge sort is stable.


