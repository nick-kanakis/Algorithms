# Depth First Search Theory

## Basic concepts

![alt text](https://upload.wikimedia.org/wikipedia/commons/thumb/1/1f/Depth-first-tree.svg/320px-Depth-first-tree.svg.png "DFS")

One starts at the root (selecting some arbitrary node as the root in the case of a graph) 
and explores as far as possible along each branch before backtracking.

**Algorithm**:
1) Insert in the stack the starting node.
2) Until stack is empty repeat
    1) Pop next node from stack
    2) Mark node as visited 
    3) Push all neighbors of the node in stack (IN REVERSE), if they are not already visited
    4) Go to 1 
    
## Time Complexity

**Worst Case**

`O(|V|+|E|)`

Where 
|V| = # vertices
|E| = # edges


## Space Complexity

`O(|V|)`

Where 
|V| = # vertices

##  Advantages.





