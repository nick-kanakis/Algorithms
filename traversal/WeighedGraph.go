package traversal

type WeighedGraph struct {
	Nodes map[int]*WeighedNode
}

type WeighedNode struct {
	Key       int
	Visited   bool
	neighbors map[*WeighedNode]int
}

func CreateWeighedGraph() *WeighedGraph {
	return &WeighedGraph{
		Nodes: make(map[int]*WeighedNode),
	}
}

func (graph *WeighedGraph) AddNode(key int) *WeighedNode {
	newNode := &WeighedNode{
		Key:     key,
		Visited: false,
		neighbors: make(map[*WeighedNode]int),
	}

	graph.Nodes[key] = newNode
	return newNode
}

func (graph *WeighedGraph) GetNode(key int) *WeighedNode {
	return graph.Nodes[key]
}

func (graph *WeighedGraph) AddEdge(originKey, destinationKey, weight int) bool {
	originNode, originNodeExists := graph.Nodes[originKey]
	if !originNodeExists {
		return false
	}

	destinationNode, destinationNodeExists := graph.Nodes[destinationKey]
	if !destinationNodeExists {
		return false
	}

	originNode.neighbors[destinationNode] = weight
	return true
}

func (graph *WeighedGraph) GetNeighbors(key int) (map[*WeighedNode]int, bool) {

	node, exists := graph.Nodes[key]
	if !exists {
		return nil, false
	}
	if len(node.neighbors) == 0 {
		return node.neighbors , false
	}
	return node.neighbors, true
}
