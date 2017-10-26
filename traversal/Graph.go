package traversal

type Graph struct {
	Nodes map[int]*Node
}

type Node struct {
	Key       int
	Visited   bool
	neighbors []*Node
}

func CreateGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]*Node),
	}
}

func (graph *Graph) AddNode(key int) *Node {
	newNode := &Node{
		Key:     key,
		Visited: false,
	}

	graph.Nodes[key] = newNode
	return newNode
}

func (graph *Graph) GetNode(key int) *Node{
	return graph.Nodes[key]
}

func (graph *Graph) AddEdge(originKey, destinationKey int) bool {
	originNode, originNodeExists := graph.Nodes[originKey]
	if !originNodeExists {
		return false
	}

	destinationNode, destinationNodeExists := graph.Nodes[destinationKey]
	if !destinationNodeExists {
		return false
	}

	originNode.neighbors = append(originNode.neighbors, destinationNode)
	return true
}

func (graph *Graph) GetNeighbors(key int) ([]*Node, bool) {

	node, exists := graph.Nodes[key]
	if !exists {
		return nil, false
	}
	if len(node.neighbors) == 0 {
		return node.neighbors, false
	}
	return node.neighbors, true
}
