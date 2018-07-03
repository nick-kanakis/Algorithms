package DFS

import (
	"fmt"
	"personal/Algorithms/traversal"
)

type Stack struct {
	nodes []*traversal.Node
}

func initializeStack() *Stack {
	return &Stack{
		nodes: make([]*traversal.Node, 0, 10),
	}
}

func (s *Stack) push(node *traversal.Node) {
	s.nodes = append(s.nodes, node)
}

func (s *Stack) pop() *traversal.Node {
	removedNode := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return removedNode
}

func (s *Stack) isEmpty() bool {
	if len(s.nodes) > 0 {
		return false
	}
	return true
}

func (s *Stack) peek() *traversal.Node {
	return s.nodes[len(s.nodes)-1]
}

func returnNodeValues(graph *traversal.Graph, key int) []int {

	stack := initializeStack()
	var keys []int

	stack.push(graph.GetNode(key))

	for !stack.isEmpty() {

		//Pop from stack
		initialNode := stack.pop()

		//Mark node as visited and put it to the result
		initialNode.Visited = true
		keys = append(keys, initialNode.Key)
		fmt.Print(initialNode.Key, " ")

		//Push all neighbors of the node in stack (IN REVERSE), if they are not already visited
		neighbors, exist := graph.GetNeighbors(initialNode.Key)
		if !exist {
			continue
		}
		for i := len(neighbors) - 1; i >= 0; i-- {
			if !neighbors[i].Visited {
				stack.push(neighbors[i])
			}
		}
	}
	return keys
}

/*
	Alternative implementation.
*/
func returnNodeValues2(graph *traversal.Graph, key int) []int {
	stack := initializeStack()
	var keys []int

	current := graph.GetNode(key)
	current.Visited = true
	stack.push(current)

	for !stack.isEmpty() {
		current, exist := findNextChild(current, graph)

		if !exist {
			stack.pop()
			if !stack.isEmpty() {
				current = stack.peek()
			} else {
				break
			}
		} else {
			current.Visited = true
			keys = append(keys, current.Key)
			stack.push(current)
		}
	}

	return keys
}

func findNextChild(node *traversal.Node, graph *traversal.Graph) (*traversal.Node, bool) {

	if node == nil {
		return nil, false
	}
	neighbors, _ := graph.GetNeighbors(node.Key)

	for _, neighbor := range neighbors {
		if !neighbor.Visited {
			return neighbor, true
		}
	}
	return nil, false
}
