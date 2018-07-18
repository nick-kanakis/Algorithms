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
    3) Push all neighbors of the node in stack (IN REVERSE), if they are not already visited.
    4) Go to 1 
    
## Time Complexity

**Worst Case**

`O(|V|+|E|)`

All vertices and edges will be explored at the worst case so time complexity is O(|V|+|E|)
OR:

>>> Incident edge is the edge that connects 2 nodes

v1 + (incident edges) + v2 + (incident edges) + .... + vn + (incident edges) = 
(v1 + v2 + ... + vn) + [(incident_edges v1) + (incident_edges v2) + ... + (incident_edges vn)] =
O(|V|)+O(|E|) = O(|V|+|E|)

Where 
|V| = # vertices
|E| = # edges


## Space Complexity

`O(|V|)` 

All nodes will be enqueued and dequeued so the space complexity is O(|V|)

Where 
|V| = # vertices

##  Advantages over BFS

1) If tree is very wide (each node has a lot of children) BFS will take a lot of time so DFS 
might be a better option

2) If solutions are frequent but located deep in the tree DFS is preferred.





