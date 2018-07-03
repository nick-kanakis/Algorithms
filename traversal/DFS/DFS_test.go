package DFS

import (
	"personal/Algorithms/traversal"
	"testing"
)

func TestBFS(t *testing.T) {

	graph := traversal.CreateGraph()

	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(3)
	graph.AddNode(4)
	graph.AddNode(5)
	graph.AddNode(6)

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 4)
	graph.AddEdge(1, 5)
	graph.AddEdge(2, 3)
	graph.AddEdge(4, 3)
	graph.AddEdge(5, 3)
	graph.AddEdge(5, 6)

	keys := returnNodeValues(graph, 1)

	for i := 0; i <= len(keys)-2; i++ {
		if keys[i] > keys[i+1] {
			t.Error()
		}
	}
}
