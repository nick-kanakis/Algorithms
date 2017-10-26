package BFS

import (
	"personal/Algorithms/traversal"
	"fmt"
)

func returnNodeValues(graph *traversal.Graph, key int) []int {
	var queue []*traversal.Node
	var keys []int

	queue = append(queue, graph.GetNode(key))

	for len(queue) > 0 {
		initialNode := queue[0];
		queue = queue[1:]                   // delete first item of queue
		keys = append(keys, initialNode.Key) //insert key into results
		fmt.Print(initialNode.Key, " ")

		neighbors, exist := graph.GetNeighbors(initialNode.Key)

		if !exist {
			continue
		}

		for _, neighbor := range neighbors {
			if neighbor.Visited {
				//if already visited do not revisit
				continue
			}
			neighbor.Visited = true
			queue = append(queue, neighbor)
		}
	}
	return keys
}
