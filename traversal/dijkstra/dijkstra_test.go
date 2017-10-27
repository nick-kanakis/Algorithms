package dijkstra

import (
	"testing"
	"personal/Algorithms/traversal"
	"fmt"
)

func TestDijkstra(t *testing.T){

	graph := traversal.CreateWeighedGraph()
	graph.AddNode(0)
	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(3)
	graph.AddNode(4)
	graph.AddNode(5)
	graph.AddNode(6)
	graph.AddNode(7)
	graph.AddNode(8)

	graph.AddEdge(0,1,4)
	graph.AddEdge(0,7,8)
	graph.AddEdge(1,7,11)
	graph.AddEdge(1,2,8)
	graph.AddEdge(7,8,7)
	graph.AddEdge(7,6,1)
	graph.AddEdge(2,8,2)
	graph.AddEdge(2,3,7)
	graph.AddEdge(2,5,4)
	graph.AddEdge(6,8,6)
	graph.AddEdge(6,5,2)
	graph.AddEdge(3,5,14)
	graph.AddEdge(3,4,9)
	graph.AddEdge(5,4,10)

	sortestPaths := GetShortestPaths(graph, 0)

	for key, value := range sortestPaths {
		fmt.Println("From node 0 it takes",value, "to get to", key)
	}

	if sortestPaths[0] != 0 {
		t.Error()
	}
	if sortestPaths[1] != 4 {
		t.Error()
	}
	if sortestPaths[2] !=12 {
		t.Error()
	}
	if sortestPaths[3] != 19 {
		t.Error()
	}
	if sortestPaths[4] != 21 {
		t.Error()
	}
	if sortestPaths[5] != 11 {
		t.Error()
	}
	if sortestPaths[6] != 9 {
		t.Error()
	}
	if sortestPaths[7] != 8 {
		t.Error()
	}
	if sortestPaths[8] != 14 {
		t.Error()
	}
}