package BFS

import (
	"testing"
	"personal/Algorithms/traversal"
	"reflect"
)

func TestBFS( t *testing.T){

	graph := traversal.CreateGraph()

	correctResult := []int{1,2,3,4,5,6,7,8,9,0,0,0,0,0,0,0,0,0}

	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(3)
	graph.AddNode(4)
	graph.AddNode(5)
	graph.AddNode(6)
	graph.AddNode(7)
	graph.AddNode(8)
	graph.AddNode(9)

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(2, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(4, 5)
	graph.AddEdge(4, 6)
	graph.AddEdge(5, 7)
	graph.AddEdge(6, 8)
	graph.AddEdge(6, 9)

	keys := returnNodeValues(graph,1)
	if reflect.DeepEqual(correctResult, keys){
		t.Error()
	}
}