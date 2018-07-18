# Dijkstra Theory

## Basic concepts

![alt text](http://www.geeksforgeeks.org/wp-content/uploads/Fig-11.jpg "Dijkstra")

Dijkstra's algorithm is an algorithm for finding the shortest paths between nodes in a graph.
The algorithm exists in many variants; Dijkstra's original variant found the shortest path between two nodes,
but a more common variant fixes a single node as the "source" node and finds shortest paths from the source to all other nodes in the graph, 
producing a shortest-path tree.

**Algorithm**:
In general alwayes select the "Cheapest" path, and expand the edge node, until there is no node left to go to.

1) Given a starting point create an array of (#NODES_OF_GRAPH) that will hold the shortest distance from source to i-node (dist[])
2) Create another array of (#NODES_OF_GRAPH) that will hold a boolean value for each node, if true node's i distance from SOURCE_NODE
is finalized (sptTree[] - shortest path tre)
3) Initialize the 2 arrays with infinity & false respectively.
4) The distance of source -> source is 0 (probably)
5) for i =0 to #NODES_OF_GRAPH -1
    1) Pick the minimum distance node e.g. U (from the minimum distance array) that is not finalized ( from the finalize array).
    2) Mark U true in sptTree
    3) Iterate neighbors of U.
        1) If neighbor (V) is not finalized (contained in sptTree) && the total distance from src to V (dist[U] + distance(U,V))
         is smaller than current value in dist[V]
        2) Update the dist[V] value with new dist[U] + distance(U,V) 

## Time Complexity

**Worst Case**

In general: `O(|E| + |V|log|V|)`

For connected graph: `O(|E|log|V|)`

To get the next minimum non-final node from the array you need O(log|V|) using a heap.
This needs to be done for all edges of the graph so |E| so the result total time complexity: `O(|E|log|V|)`

Where 
|V| = # vertices
|E| = # edges

## Space Complexity

`O(|V|^2)`

Where 
|V| = # vertices

At worst case I have to hold in memory all nodes and edges O(|V|+|E|) but in worst case |E| = |V^2| since
each node can connect to each other node so  O(|V|+|E|) =  O(|V|^2) 

##  Usages.

The shortest path algorithm is widely used in network routing protocols, most notably IS-IS (Intermediate System to Intermediate System) and Open Shortest Path First (OSPF). 
It is also employed as a subroutine in other algorithms such as Johnson's.



