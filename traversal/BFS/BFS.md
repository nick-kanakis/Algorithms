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
    2) Visit and enqueue all neighbors of the node (Do not forget to mark the nodes as visited in order to not visit them again and go into a infinite loop)
    3) Go to i 


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
