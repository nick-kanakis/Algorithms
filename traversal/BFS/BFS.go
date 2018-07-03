package BFS

import (
	"fmt"
	"personal/Algorithms/traversal"
)

type Queue struct {
	nodes []*traversal.Node
}

func (q *Queue) enqueue(node *traversal.Node) {
	q.nodes = append(q.nodes, node)
}

func (q *Queue) dequeue() *traversal.Node {
	removedNode := q.nodes[0]
	q.nodes = q.nodes[1:]
	return removedNode
}

func (q *Queue) isEmpty() bool {
	if len(q.nodes) > 0 {
		return false
	}
	return true
}

func initializeQueue() *Queue {
	return &Queue{
		nodes: make([]*traversal.Node, 0, 10),
	}
}

func returnNodeValues(graph *traversal.Graph, key int) []int {
	queue := initializeQueue()
	var keys []int

	queue.enqueue(graph.GetNode(key))

	for !queue.isEmpty() {
		initialNode := queue.dequeue()       // delete first item of queue
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
			queue.enqueue(neighbor)
		}
	}
	return keys
}
