package DFS

import (
	"personal/Algorithms/traversal"
	"fmt"
)

func returnNodeValues(graph *traversal.Graph, key int) []int {

	var stack []*traversal.Node
	var keys []int

	stack = append(stack, graph.GetNode(key))

	for len(stack) > 0 {

		//Pop from stack
		initialNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		//Mark node as visited and put it to the result
		initialNode.Visited = true
		keys = append(keys, initialNode.Key)
		fmt.Print(initialNode.Key, " ")


		//Push all neighbors of the node in stack (IN REVERSE), if they are not already visited
		neighbors, exist := graph.GetNeighbors(initialNode.Key)
		if !exist {
			continue
		}
		for  i:=len(neighbors) -1; i>=0 ; i--{
			if !neighbors[i].Visited{
				stack = append(stack, neighbors[i])
			}
		}
	}

	return keys
}
