package traversal

import "testing"

func TestGraph(t *testing.T) {

	graph := CreateGraph()

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

	neighbors, _ := graph.GetNeighbors(1)
	if len(neighbors) != 2 {
		t.Error()
	}

	if neighbors[0].key != 2 ||neighbors[1].key != 3{
		t.Error()
	}

	neighbors, _ = graph.GetNeighbors(2)
	if len(neighbors) != 2 {
		t.Error()
	}

	if neighbors[0].key != 4 ||neighbors[1].key != 5{
		t.Error()
	}

	neighbors, _ = graph.GetNeighbors(3)
	if len(neighbors) != 1 {
		t.Error()
	}

	if neighbors[0].key != 4 {
		t.Error()
	}

	neighbors, _ = graph.GetNeighbors(4)
	if len(neighbors) != 2 {
		t.Error()
	}

	if neighbors[0].key != 5 ||neighbors[1].key != 6{
		t.Error()
	}

	neighbors, _ = graph.GetNeighbors(5)
	if len(neighbors) != 1 {
		t.Error()
	}

	if neighbors[0].key != 7 {
		t.Error()
	}

	neighbors, _ = graph.GetNeighbors(6)
	if len(neighbors) != 2 {
		t.Error()
	}

	if neighbors[0].key != 8 ||neighbors[1].key != 9{
		t.Error()
	}

	_, exists := graph.GetNeighbors(7)
	if exists {
		t.Error()
	}

	_, exists = graph.GetNeighbors(8)
	if exists {
		t.Error()
	}

	_, exists = graph.GetNeighbors(9)
	if exists {
		t.Error()
	}

}
