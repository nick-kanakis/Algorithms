package traversal

import "testing"

func TestWeighedGraph(t *testing.T) {

	graph := CreateWeighedGraph()

	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(3)
	graph.AddNode(4)
	graph.AddNode(5)
	graph.AddNode(6)
	graph.AddNode(7)
	graph.AddNode(8)
	graph.AddNode(9)

	graph.AddEdge(1, 2, 3)
	graph.AddEdge(1, 3, 3)
	graph.AddEdge(3, 4, 3)
	graph.AddEdge(2, 4, 3)
	graph.AddEdge(2, 5, 3)
	graph.AddEdge(4, 5, 3)
	graph.AddEdge(4, 6, 3)
	graph.AddEdge(5, 7, 3)
	graph.AddEdge(6, 8, 3)
	graph.AddEdge(6, 9, 3)

	neighbors, _ := graph.GetNeighbors(1)
	if len(neighbors) != 2 {
		t.Error()
	}

	for neighbor := range neighbors {
		if neighbor.Key != 2 && neighbor.Key != 3 {
			t.Error()
		}
	}

	neighbors, _ = graph.GetNeighbors(2)
	if len(neighbors) != 2 {
		t.Error()
	}

	for neighbor := range neighbors {
		if neighbor.Key != 4 && neighbor.Key != 5 {
			t.Error()
		}
	}

	neighbors, _ = graph.GetNeighbors(3)
	if len(neighbors) != 1 {
		t.Error()
	}

	for neighbor := range neighbors {
		if neighbor.Key != 4 {
			t.Error()
		}
	}

	neighbors, _ = graph.GetNeighbors(4)
	if len(neighbors) != 2 {
		t.Error()
	}

	for neighbor := range neighbors {
		if neighbor.Key != 5 && neighbor.Key != 6 {
			t.Error()
		}
	}

	neighbors, _ = graph.GetNeighbors(5)
	if len(neighbors) != 1 {
		t.Error()
	}

	for neighbor := range neighbors {
		if neighbor.Key != 7 {
			t.Error()
		}
	}

	neighbors, _ = graph.GetNeighbors(6)
	if len(neighbors) != 2 {
		t.Error()
	}

	for neighbor := range neighbors {
		if neighbor.Key != 8 && neighbor.Key != 9 {
			t.Error()
		}
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
