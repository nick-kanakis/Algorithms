# Breadth First Search Theory

## Basic concepts

Order in which the nodes are expanded

![alt text](https://upload.wikimedia.org/wikipedia/commons/thumb/3/33/Breadth-first-tree.svg/320px-Breadth-first-tree.svg.png "")

starts at the tree root (or some arbitrary node of a graph, sometimes referred to as a 'search key') and explores the neighbor nodes first, 
before moving to the next level neighbours.

**Algorithm**:
1) Insert in the queue the starting node.
2) Until queue is empty repeat
    1) Remove next node in queue
    2) Visit and enqueue all neighbours of the node
    3) Go to i 


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
